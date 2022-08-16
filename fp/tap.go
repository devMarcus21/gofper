package fp

func Tap(fun func(Obj) Obj) func(Obj) Obj {
	return func(arg Obj) Obj {
		fun(arg)
		return arg
	}
}
