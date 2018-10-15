package model

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
)

type Model struct {
	Type        string
	IsInterface bool
	Map         MapModel
	Slice       FieldModel
	GroupBys    []*FieldModel
	Imports     []*ast.ImportSpec
}

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
		fieldType, imports := resolveFieldType(cfg, f.Type)
		field := FieldModel{
			Name:     strcase.ToCamel(fieldName),
			Type:     fieldType,
			Accessor: fieldName,
		}
		if field.Type == "" {
			continue
		}
		if result.IsInterface {
			field.Accessor += "()"
		}
		if result.Map.Key == nil && fieldName == cfg.KeyField {
			result.Map.Key = &field
		}
		result.Imports = append(result.Imports, imports...)
		result.GroupBys = append(result.GroupBys, &field)
	}

	return result
}

func resolveFieldType(cfg *Config, field ast.Expr) (string, []*ast.ImportSpec) {
	switch ft := field.(type) {
	case *ast.StarExpr:
		switch i := ft.X.(type) {
		case *ast.Ident:
			return "*" + i.Name, nil
		case *ast.SelectorExpr:
			selectedType, imports := resolveFieldType(cfg, i.X)
			for _, im := range cfg.Imports {
				if im.Name != nil && im.Name.Name == selectedType || strings.HasSuffix(im.Path.Value, "/"+selectedType+`"`) {
					imports = append(imports, im)
				}
			}
			return "*" + selectedType + "." + i.Sel.Name, imports
		}
	case *ast.Ident:
		return ft.Name, nil
	case *ast.FuncType:
		if ft.Results.NumFields() == 1 {
			return resolveFieldType(cfg, ft.Results.List[0].Type)
		}
	case *ast.ArrayType:
		switch l := ft.Len.(type) {
		case *ast.BasicLit:
			return fmt.Sprintf("[%s]%s", l.Value, ft.Elt), nil
		}
	}

	logrus.Debugf("No generation for field of type: %#v", field)
	return "", nil
}
