package validator

import (
	"github.com/mertyildiran/gqlparser/v2/ast"
	. "github.com/mertyildiran/gqlparser/v2/validator"
)

func init() {
	AddRule("UniqueVariableNames", func(observers *Events, addError AddErrFunc) {
		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			seen := map[string]int{}
			for _, def := range operation.VariableDefinitions {
				// add the same error only once per a variable.
				if seen[def.Variable] == 1 {
					addError(
						Message(`There can be only one variable named "$%s".`, def.Variable),
						At(def.Position),
					)
				}
				seen[def.Variable]++
			}
		})
	})
}
