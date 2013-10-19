package swift

import (
	"fmt"
	"testing"
)

func TestVar(t *testing.T) {
	setName := &AST{nil, "a", VAL}
	setValue := &AST{nil, 3, VAL}
	set := &AST{[]*AST{setName, setValue}, "set", FUNC}

	addVal := &AST{nil, 1, VAL}
	addVar := &AST{nil, "a", VAR}
	add := &AST{[]*AST{addVal, addVar}, "add", FUNC}

	main := &AST{[]*AST{set, add}, "main", FUNC}

	fmt.Println(main)

	env := NewEnvironment()

	Exec(main, env)

	fmt.Println(main)
}

func TestAST(t *testing.T) {
	arg10 := &AST{nil, 1, VAR}
	arg11 := &AST{nil, 2, VAR}
	arg0 := &AST{[]*AST{arg10, arg11}, "subtract", FUNC}

	arg1 := &AST{nil, 3, VAR}

	ast := &AST{[]*AST{arg0, arg1}, "add", FUNC}

	fmt.Println(ast)

	env := NewEnvironment()

	Exec(ast, env)

	fmt.Println(ast)

}
