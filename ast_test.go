package swift

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	arg10 := &AST{nil, 1}
	arg11 := &AST{nil, 2}
	arg0 := &AST{[]*AST{arg10, arg11}, "subtract"}

	arg1 := &AST{nil, 3}

	ast := &AST{[]*AST{arg0, arg1}, "add"}

	fmt.Println(ast)

	Exec(ast)

	fmt.Println(ast)

}
