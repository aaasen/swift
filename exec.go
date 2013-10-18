package swift

import ()

func Exec(ast *AST) {
	if ast.Children != nil {
		for _, child := range ast.Children {
			if child.Children != nil {
				Exec(child)
			}
		}
	}

	if value, ok := ast.Value.(string); ok {
		switch value {
		case "add":
			ast.Value = Add(ast.Children[0].Value.(int), ast.Children[1].Value.(int))
			ast.Children = nil
		case "subtract":
			ast.Value = Subtract(ast.Children[0].Value.(int), ast.Children[1].Value.(int))
			ast.Children = nil
		case "divide":
			ast.Value = Divide(ast.Children[0].Value.(int), ast.Children[1].Value.(int))
			ast.Children = nil
		case "multiply":
			ast.Value = Multiply(ast.Children[0].Value.(int), ast.Children[1].Value.(int))
			ast.Children = nil
		}
	}
}
