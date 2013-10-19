package swift

type Environment struct {
	Vars  map[string]interface{}
	Funcs map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		make(map[string]interface{}),
		make(map[string]interface{}),
	}
}

func (self *Environment) Get(e interface{}) interface{} {
	if stringVal, ok := e.(string); ok {
		if val, ok := self.Vars[stringVal]; ok {
			return val
		} else {
			// variable undefined
		}
	}

	return e
}
