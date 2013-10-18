package swift

func Add(a, b int) int {
	return a + b
}

var funcs = map[string]interface{}{
	"add": Add,
}
