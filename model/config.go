package model

import "go/ast"

// Config for the generation process
type Config struct {
	Type    *ast.TypeSpec
	Imports []*ast.ImportSpec

	Generators  []string
	KeyField    string
	GroupFields []string
	MapPrefix   string
	SlicePrefix string
}
