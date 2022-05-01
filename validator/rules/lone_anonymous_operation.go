package validator

import (
	"github.com/mertyildiran/gqlparser/ast"
	. "github.com/mertyildiran/gqlparser/validator"
)

func init() {
	AddRule("LoneAnonymousOperation", func(observers *Events, addError AddErrFunc) {
		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			if operation.Name == "" && len(walker.Document.Operations) > 1 {
				addError(
					Message(`This anonymous operation must be the only defined operation.`),
					At(operation.Position),
				)
			}
		})
	})
}
