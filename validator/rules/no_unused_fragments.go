package validator

import (
	"github.com/mertyildiran/gqlparser/ast"
	. "github.com/mertyildiran/gqlparser/validator"
)

func init() {
	AddRule("NoUnusedFragments", func(observers *Events, addError AddErrFunc) {

		inFragmentDefinition := false
		fragmentNameUsed := make(map[string]bool)

		observers.OnFragmentSpread(func(walker *Walker, fragmentSpread *ast.FragmentSpread) {
			if !inFragmentDefinition {
				fragmentNameUsed[fragmentSpread.Name] = true
			}
		})

		observers.OnFragment(func(walker *Walker, fragment *ast.FragmentDefinition) {
			inFragmentDefinition = true
			if !fragmentNameUsed[fragment.Name] {
				addError(
					Message(`Fragment "%s" is never used.`, fragment.Name),
					At(fragment.Position),
				)
			}
		})
	})
}
