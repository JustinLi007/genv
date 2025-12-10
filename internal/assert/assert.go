package assert

import (
	"fmt"
	"os"
	"reflect"
)

func NoErr(e error, msg string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "genv: %s: %s\n", msg, e)
		os.Exit(1)
	}
}

func Nil(v any, msg string) {
	if !IsNil(v) {
		fmt.Fprintf(os.Stderr, "genv: %s\n", msg)
		os.Exit(1)
	}
}

func NotNil(v any, msg string) {
	if IsNil(v) {
		fmt.Fprintf(os.Stderr, "genv: %s\n", msg)
		os.Exit(1)
	}
}

func IsNil(o any) bool {
	if o == nil {
		return true
	}

	v := reflect.ValueOf(o)
	switch v.Kind() {
	case reflect.Interface, reflect.Func,
		reflect.Chan, reflect.Slice,
		reflect.Map, reflect.Pointer,
		reflect.UnsafePointer:
		return v.IsNil()
	}

	return false
}

func True(cond bool, msg string) {
	if cond == false {
		fmt.Fprintf(os.Stderr, "genv: %s\n", msg)
		os.Exit(1)
	}
}

func False(cond bool, msg string) {
	if cond == true {
		fmt.Fprintf(os.Stderr, "genv: %s\n", msg)
		os.Exit(1)
	}
}
