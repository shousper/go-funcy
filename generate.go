package funcy

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/shousper/go-funcy/fragments"
	"github.com/shousper/go-funcy/model"
)

type specInfo struct {
	typ     *ast.TypeSpec
	imports []*ast.ImportSpec
}

// Generate populates the targetPath with the funcy funcs for the targetType
func Generate(targetPath, targetType string, cfg *model.Config) error {
	basePath := targetPath
	goPathSrc := path.Join(build.Default.GOPATH, "src")
	if !strings.HasPrefix(targetPath, goPathSrc) {
		basePath = path.Join(build.Default.GOPATH, "src", targetPath)
	}
	pkg, err := build.ImportDir(basePath, build.ImportComment)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	t, err := loadPackage(basePath, fset, pkg, targetType)
	if err != nil {
		return err
	}

	cfg.Type = t.typ
	cfg.Imports = t.imports
	return writeType(basePath, pkg.Name, cfg)
}

func writeType(basePath, sourcePackage string, cfg *model.Config) error {
	f := jen.NewFile(sourcePackage)

	m := model.Create(cfg)

	fragments.SliceOf(f, m)
	fragments.SliceOfGroupBys(f, m)

	if m.Map.Key != nil {
		fragments.SliceOfAsMap(f, m)

		fragments.MapOf(f, m)
		fragments.MapOfKeys(f, m)
		fragments.MapOfValues(f, m)
		fragments.MapOfSelect(f, m)
		fragments.MapOfGroupBys(f, m)
	}

	name := strcase.ToSnake(cfg.Type.Name.Name)
	return f.Save(path.Join(basePath, name+".funcy.go"))
}

func loadPackage(basePath string, fset *token.FileSet, pkg *build.Package, typeName string) (*specInfo, error) {
	result := new(specInfo)

	for _, goFile := range pkg.GoFiles {
		node, err := parser.ParseFile(fset, path.Join(basePath, goFile), nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse '%s': %v", goFile, err)
		}

		ast.Inspect(node, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.ImportSpec:
				result.imports = append(result.imports, v)
			case *ast.TypeSpec:
				if v.Name.Name == typeName {
					result.typ = v
				}
			}
			return true
		})
	}

	if result.typ == nil {
		return nil, fmt.Errorf("type '%s' not found", typeName)
	}

	return result, nil
}
