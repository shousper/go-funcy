package model

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
)

// Model type information for generation output
// Passed to fragments to create functions
type Model struct {
	Type        string
	IsInterface bool
	Map         MapModel
	Slice       FieldModel
	GroupBys    []FieldModel
}

// FieldModel represents struct or interface field/method information used for generation
type FieldModel struct {
	Name     string
	Type     string
	Qual     string
	Pointer  bool
	Accessor string
}

// MapModel represents information for map fragment generation
type MapModel struct {
	Name string
	Key  *FieldModel
}

// Create a generation model
func Create(cfg *Config) Model {
	if cfg.MapPrefix == "" {
		cfg.MapPrefix = "MapOf"
	}
	if cfg.SlicePrefix == "" {
		cfg.SlicePrefix = "SliceOf"
	}

	typeName := cfg.Type.Name.Name

	result := Model{
		Type: typeName,
		Slice: FieldModel{
			Name: cfg.SlicePrefix + typeName,
			Type: typeName,
		},
		Map: MapModel{
			Name: cfg.MapPrefix + typeName,
		},
	}

	var fieldList []*ast.Field

	switch v := cfg.Type.Type.(type) {
	case *ast.StructType:
		// Struct types are always pointers in slices or maps.
		result.Type = typeName
		// Slice type will be the pointer type
		result.Slice.Type = result.Type

		fieldList = v.Fields.List
	case *ast.InterfaceType:
		fieldList = v.Methods.List
		result.IsInterface = true
	}

	for _, f := range fieldList {
		fieldName := f.Names[0].Name
		field := resolveField(cfg, fieldName, f.Type)
		if field.Type == "" {
			continue
		}
		if result.IsInterface {
			field.Accessor += "()"
		}
		if result.Map.Key == nil && fieldName == cfg.KeyField {
			result.Map.Key = field
		}
		if shouldGroupBy(cfg, fieldName) {
			result.GroupBys = append(result.GroupBys, *field)
		}
	}

	return result
}

func shouldGroupBy(cfg *Config, field string) bool {
	if len(cfg.GroupFields) == 0 {
		return true
	}
	for _, g := range cfg.GroupFields {
		if g == field {
			return true
		}
	}
	return false
}

func resolveField(cfg *Config, fieldName string, expr ast.Expr) *FieldModel {
	field := FieldModel{
		Name:     strcase.ToCamel(fieldName),
		Accessor: fieldName,
	}

	switch ft := expr.(type) {
	case *ast.StarExpr:
		field.Pointer = true
		switch i := ft.X.(type) {
		case *ast.Ident:
			field.Type = i.Name
			return &field
		case *ast.SelectorExpr:
			field.Type = i.Sel.Name
			selectedField := resolveField(cfg, field.Type, i.X)

			for _, im := range cfg.Imports {
				importPath := im.Path.Value[1 : len(im.Path.Value)-1]
				importName := ""
				if im.Name != nil {
					importName = im.Name.Name
				}
				if importName == selectedField.Type || strings.HasSuffix(importPath, "/"+selectedField.Type) {
					field.Qual = importPath
				}
			}
			return &field
		}
	case *ast.Ident:
		field.Type = ft.Name
		return &field
	case *ast.FuncType:
		if ft.Results.NumFields() == 1 {
			return resolveField(cfg, fieldName, ft.Results.List[0].Type)
		}
	case *ast.ArrayType:
		switch l := ft.Len.(type) {
		case *ast.BasicLit:
			field.Type = fmt.Sprintf("[%s]%s", l.Value, ft.Elt)
			return &field
		}
	}

	logrus.Debugf("No generation for field of type: %#v", field)
	return nil
}
