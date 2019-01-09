package model

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
)

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
			Type: TypeModel{Name: typeName},
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
		result.Slice.Type = TypeModel{Name: result.Type}

		fieldList = v.Fields.List
	case *ast.InterfaceType:
		fieldList = v.Methods.List
		result.IsInterface = true
	}

	for _, f := range fieldList {
		fieldName := f.Names[0].Name
		field := resolveField(cfg, fieldName, f.Type)
		if field == nil {
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
	fieldType := walkFieldType(cfg, &TypeModel{}, expr)
	if fieldType == nil {
		return nil
	}
	return &FieldModel{
		Name:     strcase.ToCamel(fieldName),
		Accessor: fieldName,
		Type:     *fieldType,
	}
}

func walkFieldType(cfg *Config, typ *TypeModel, expr ast.Expr) *TypeModel {
	switch ft := expr.(type) {
	case *ast.StarExpr:
		typ.Pointer = true
		return walkFieldType(cfg, typ, ft.X)
	case *ast.SelectorExpr:
		typ.Name = ft.Sel.Name
		typ.Qualifier = resolveQualifier(cfg, ft.X)
		return typ
	case *ast.Ident:
		typ.Name = ft.Name
		return typ
	case *ast.FuncType:
		if ft.Results.NumFields() == 1 {
			return walkFieldType(cfg, typ, ft.Results.List[0].Type)
		}
	case *ast.ArrayType:
		switch l := ft.Len.(type) {
		case *ast.BasicLit:
			typ.Name = fmt.Sprintf("[%s]%s", l.Value, ft.Elt)
			return typ
		}
	}

	// TODO("This basically means 'unsupported', handle it better.")
	logrus.Debugf("No generation for type: %s", reflect.TypeOf(expr))
	return nil
}

func resolveQualifier(cfg *Config, expr ast.Expr) string {
	importedType := walkFieldType(cfg, &TypeModel{}, expr)
	for _, im := range cfg.Imports {
		importPath := im.Path.Value[1 : len(im.Path.Value)-1]
		importName := ""
		if im.Name != nil {
			importName = im.Name.Name
		}
		if importName == importedType.Name || strings.HasSuffix(importPath, "/"+importedType.Name) {
			return importPath
		}
	}
	return ""
}
