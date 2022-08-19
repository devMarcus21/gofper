package fp

func Pipe(args ...func(Obj) Obj) func(Obj) Obj {
	return func(arg Obj) Obj {
		if len(args) <= 0 {
			return arg
		}

		if args[0] == nil {
			return Pipe(args[1:]...)(Obj{}) // must pass empty state because passing nil will cause different behavior
		}

		return Pipe(args[1:]...)(args[0](arg))
	}
}
