package gqlparser

import (
	"github.com/mertyildiran/gqlparser/ast"
	"github.com/mertyildiran/gqlparser/gqlerror"
	"github.com/mertyildiran/gqlparser/parser"
	"github.com/mertyildiran/gqlparser/validator"
	_ "github.com/mertyildiran/gqlparser/validator/rules"
)

func LoadSchema(str ...*ast.Source) (*ast.Schema, *gqlerror.Error) {
	return validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
}

func MustLoadSchema(str ...*ast.Source) *ast.Schema {
	s, err := validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
	if err != nil {
		panic(err)
	}
	return s
}

func LoadQuery(schema *ast.Schema, str string) (*ast.QueryDocument, gqlerror.List) {
	query, err := parser.ParseQuery(&ast.Source{Input: str})
	if err != nil {
		return nil, gqlerror.List{err}
	}
	errs := validator.Validate(schema, query)
	if errs != nil {
		return nil, errs
	}

	return query, nil
}

func MustLoadQuery(schema *ast.Schema, str string) *ast.QueryDocument {
	q, err := LoadQuery(schema, str)
	if err != nil {
		panic(err)
	}
	return q
}
