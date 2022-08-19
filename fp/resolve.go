package fp

func Resolve(arg Obj) func(Obj) Obj {
	return func(state Obj) Obj {
		return deepResolve(state)
	}
}

func deepResolve(state Obj) Obj {
	newState := Obj{}

	for propName, prop := range state {
		switch prop.(type) {
		case nil:
			continue // evict nil values
		case Obj:
			newState[propName] = deepResolve(prop.(Obj))
		case map[string]any:
			newState[propName] = deepResolve(Obj(prop.(map[string]any)))
		default:
			newState[propName] = prop
		}
	}

	return newState
}

// Useful for appending new properties onto state while in pipe
func ResolveInto(append Obj) func(Obj) Obj {
	return func(state Obj) Obj {
		newState := Obj{}

		for propName, prop := range state {
			newState[propName] = prop
		}

		for propName, prop := range append {
			newState[propName] = prop
		}

		return newState
	}
}
