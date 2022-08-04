package fp

func Pipe(args ...any) func(any) any {
	return func(arg any) any {
		if len(args) <= 0 {
			return arg
		}

		return Pipe(args[1:]...)(Resolve(args[0])(arg))
	}
}
