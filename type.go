package goutils

import (
	"reflect"
)

// Returns a string of variable type
func Type(arg interface{}) string {
	return reflect.TypeOf(arg).String()
}

// Int
func IsInt(arg interface{}) bool {
	switch arg.(type) {
	case int:
		return true
	default:
		return false
	}
}

// Int64
func IsInt64(arg interface{}) bool {
	switch arg.(type) {
	case int64:
		return true
	default:
		return false
	}
}

// Float64
func IsFloat(arg interface{}) bool {
	return Type(arg) == "float64"
}

// String
func IsString(arg interface{}) bool {
	return Type(arg) == "string"
}

func IsTime(arg interface{}) bool {
	return Type(arg) == "time.Time"
}

func IsBool(arg interface{}) bool {
	return Type(arg) == "bool"
}

func IsSlice(arg interface{}) bool {
	return reflect.TypeOf(arg).Kind().String() == "slice"
}
