package tests

import (
	"testing"

	"github.com/devMarcus21/gofper/fp"
)

func TestGet_TapPassedFunction_ReturnsOriginalArg(t *testing.T) {
	arg := fp.Obj{"foo": "bar"}
	passedFun := func(fp.Obj) fp.Obj {
		return fp.Obj{"foo": "NotBar"}
	}
	fun := fp.Pipe(fp.Tap(passedFun))

	result := fun(arg)

	CheckEquality(t, "TestGet_TapPassedFunction_ReturnsOriginalArg", arg, result)
}
