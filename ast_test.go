package swift

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	arg1 := &AST{nil, 1}
	arg2 := &AST{nil, 2}
	ast := &AST{[]*AST{arg1, arg2}, "add"}

	aFunc := funcs[ast.Value.(string)].(func(int, int) int)

	fmt.Println(aFunc(ast.Children[0].Value.(int), ast.Children[1].Value.(int)))

}
