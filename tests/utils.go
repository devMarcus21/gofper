package tests

import (
	"reflect"
	"testing"
)

func CheckEquality(t *testing.T, functionName string, expected interface{}, got interface{}) {

	message := "%s %s. Expected %v of type %T, got %v of type %T"

	if reflect.DeepEqual(expected, got) {
		t.Logf(message, functionName, "PASSED", expected, expected, got, got)
	} else {
		t.Errorf(message, functionName, "FAILED", expected, expected, got, got)
	}
}
