package tests

import (
	"testing"

	"github.com/devMarcus21/gofper/fp"
)

var functions map[string]func(arg fp.Obj) fp.Obj = map[string]func(arg fp.Obj) fp.Obj{
	"fun1": func(arg fp.Obj) fp.Obj {
		return fp.Obj{"barbar": arg["bar"]}
	},
	"fun2": func(arg fp.Obj) fp.Obj {
		return fp.Obj{"barbar": arg["barbar"], "aan": "cat"}
	},
}

func TestPipe_OneFunction_ReturnObj(t *testing.T) {
	fun1 := functions["fun1"]
	fun := fp.Pipe(fun1)
	expected := fp.Obj{"barbar": 123}

	result := fun(fp.Obj{"bar": 123})

	CheckEquality(t, "TestPipe_OneFunction_ReturnObj", expected, result)
}

func TestPipe_TwoFunctions_ReturnObj(t *testing.T) {
	fun1 := functions["fun1"]
	fun2 := functions["fun2"]
	fun := fp.Pipe(fun1, fun2)
	expected := fp.Obj{"aan": "cat", "barbar": 5}

	result := fun(fp.Obj{"bar": 5})

	CheckEquality(t, "TestPipe_TwoFunctions_ReturnObj", expected, result)
}

func TestPipe_OneFunctionGivenNil_ReturnNilResult(t *testing.T) {
	fun1 := functions["fun1"]
	fun := fp.Pipe(fun1)
	expected := fp.Obj{"barbar": nil}

	result := fun(nil)

	CheckEquality(t, "TestPipe_OneFunctionGivenNil_ReturnNilResult", expected, result)
}

func TestPipe_NoFunction_ReturnObj(t *testing.T) {
	fun := fp.Pipe()
	expected := fp.Obj{"barbar": nil}

	result := fun(expected)

	CheckEquality(t, "TestPipe_NoFunction_ReturnObj", expected, result)
}

func TestPipe_PassedNil_ReturnObj(t *testing.T) {
	fun := fp.Pipe(nil)
	expected := fp.Obj{}

	result := fun(fp.Obj{"barbar": "what"})

	CheckEquality(t, "TestPipe_PassedNil_ReturnObj", expected, result)
}
