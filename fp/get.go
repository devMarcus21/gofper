package fp

func Get(keys ...string) func(Obj) Obj {
	return func(arg Obj) Obj {
		if len(keys) <= 0 {
			return Obj{"result": nil} // TODO think about these nil cases
		}

		if len(keys) == 1 {
			return Obj{"result": arg[keys[0]]}
		}

		switch arg[keys[0]].(type) {
		case Obj:
			return Get(keys[1:]...)(arg[keys[0]].(Obj))
		case map[string]any:
			return Get(keys[1:]...)(Obj(arg[keys[0]].(map[string]any)))
		default:
			return Obj{"result": nil} // TODO think about these nil cases
		}
	}
}
