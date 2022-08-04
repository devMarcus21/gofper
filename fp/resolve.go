package fp

func Resolve(arg any) Operator {
	switch arg.(type) {
	case Operator:
		return arg.(Operator)
	default:
		return func(param any) any {
			return arg
		}
	}
}
