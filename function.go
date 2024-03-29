package gdk

import (
	"errors"
	"reflect"
)

var (
	// ErrParamsNotAdapted  params length invalid
	ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

// Functions bundle of functions
type Functions map[string]reflect.Value

// NewFunctions function maps
func NewFunctions() Functions {
	return make(Functions, 0)
}

// Bind the function with the given function name
func (f Functions) Bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()
	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	f[name] = v
	return
}

// Call the function with the given name and params
func (f Functions) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := f[name]; !ok {
		err = errors.New(name + " does not exist.")
		return
	}

	if len(params) != f[name].Type().NumIn() {
		err = ErrParamsNotAdapted
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in)
	return
}
