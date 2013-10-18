package swift

// Abstract Syntax Tree

type AST struct {
	Children []*AST
	Value    interface{}
}
