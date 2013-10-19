package swift

import (
	"fmt"
)

func Exec(ast *AST, env *Environment) {
	if ast.Children != nil {
		for _, child := range ast.Children {
			if child.Children != nil {
				Exec(child, env)
			}
		}
	}

	if ast.Type == FUNC {
		switch ast.Value {
		case "add":
			ast.Value = Add(env.Get(ast.Children[0].Value).(int),
				env.Get(ast.Children[1].Value).(int))
			ast.Children = nil
		case "subtract":
			ast.Value = Subtract(env.Get(ast.Children[0].Value).(int),
				env.Get(ast.Children[1].Value).(int))
			ast.Children = nil
		case "divide":
			ast.Value = Divide(env.Get(ast.Children[0].Value).(int),
				env.Get(ast.Children[1].Value).(int))
			ast.Children = nil
		case "multiply":
			ast.Value = Multiply(env.Get(ast.Children[0].Value).(int),
				env.Get(ast.Children[1].Value).(int))
			ast.Children = nil
		case "print":
			fmt.Println(env.Get(ast.Children[0].Value))
			ast.Remove(ast)
		case "set":
			env.Vars[ast.Children[0].Value.(string)] = ast.Children[1].Value
			ast.Remove(ast)
		}
	}
}
