package validator

import (
	_ "embed"

	"github.com/mertyildiran/gqlparser/v2/ast"
)

//go:embed prelude.graphql
var preludeGraphql string

var Prelude = &ast.Source{
	Name:    "prelude.graphql",
	Input:   preludeGraphql,
	BuiltIn: true,
}
