package model

import "go/ast"

type Config struct {
	Type    *ast.TypeSpec
	Imports []*ast.ImportSpec

	KeyField    string
	MapPrefix   string
	SlicePrefix string
}
