package tests

import (
	"testing"

	"github.com/devMarcus21/gofper/fp"
)

func TestGet_GetProperty_ReturnsObj(t *testing.T) {
	expected := fp.Obj{"result": 123}

	result := fp.Pipe(fp.Get("bar"))(fp.Obj{"bar": expected["result"]})

	CheckEquality(t, "TestGet_GetProperty_ReturnsObj", expected, result)
}

func TestGet_GetPropertyPassNil_ReturnsObj(t *testing.T) {
	expected := fp.Obj{"result": nil}

	result := fp.Pipe(fp.Get("bar"))(nil)

	CheckEquality(t, "TestGet_GetPropertyPassNil_ReturnsObj", expected, result)
}

func TestGet_GetNestedProperties_ReturnsObj(t *testing.T) {
	expected := fp.Obj{"result": "foo"}
	fun := fp.Pipe(fp.Get("bar", "key"))

	result := fun(
		fp.Obj{
			"bar": fp.Obj{
				"key": expected["result"],
			},
		},
	)

	CheckEquality(t, "TestGet_GetNestedProperties_ReturnsObj", expected, result)
}

func TestGet_GetNestedPropertiesEmptyObj_ReturnsNilObj(t *testing.T) {
	expected := fp.Obj{"result": nil}
	fun := fp.Pipe(fp.Get("bar", "key"))

	result := fun(fp.Obj{})

	CheckEquality(t, "TestGet_GetNestedPropertiesEmptyObj_ReturnsEmptyObj", expected, result)
}

func TestGet_GetPassedNoFunctions_ReturnsNilObj(t *testing.T) {
	expected := fp.Obj{"result": nil}
	fun := fp.Pipe(fp.Get())

	result := fun(fp.Obj{})

	CheckEquality(t, "TestGet_GetPassedNoFunctions_ReturnsNilObj", expected, result)
}

func TestGet_GetNestedPropertiesMap_ReturnsObj(t *testing.T) {
	expected := fp.Obj{"result": "foo"}
	fun := fp.Pipe(fp.Get("bar", "key"))

	result := fun(
		fp.Obj{
			"bar": map[string]any{
				"key": expected["result"],
			},
		},
	)

	CheckEquality(t, "TestGet_GetNestedPropertiesMap_ReturnsObj", expected, result)
}

func TestGet_GetNestedPropertiesMapOfNestedMap_ReturnsObj(t *testing.T) {
	expected := fp.Obj{"result": "foo"}
	fun := fp.Pipe(fp.Get("bar", "key"))

	result := fun(
		map[string]any{
			"bar": map[string]any{
				"key": expected["result"],
			},
		},
	)

	CheckEquality(t, "TestGet_GetNestedPropertiesMapOfNestedMap_ReturnsObj", expected, result)
}

func TestGet_GetNestedPropertiesOfObjInMap_ReturnsObj(t *testing.T) {
	expected := fp.Obj{"result": "foo"}
	fun := fp.Pipe(fp.Get("bar", "key"))

	result := fun(
		map[string]any{
			"bar": fp.Obj{
				"key": expected["result"],
			},
		},
	)

	CheckEquality(t, "TestGet_GetNestedPropertiesMapOfNestedMap_ReturnsObj", expected, result)
}
