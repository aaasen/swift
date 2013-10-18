package swift

func Exec(ast *AST) {
	if ast.Children != nil {
		for _, child := range ast.Children {
			if child.Children != nil {
				Exec(child)
			}
		}
	}

	if aFunc, ok := funcs[ast.Value.(string)].(func(int, int) int); ok {
		ast.Value = aFunc(ast.Children[0].Value.(int), ast.Children[1].Value.(int))
	}
}
