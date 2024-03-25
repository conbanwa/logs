package logs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Same(a, b interface{}) bool {
	return assert.ObjectsAreEqualValues(a, b)
}

func IfSame(a, b interface{}, args ...interface{}) bool {
	if assert.ObjectsAreEqualValues(a, b) {
		Log.output(L_INFO, A, Concat(append(args, ":", A)...))
		return true
	}
	return false
}

func IfNotSame(a, b interface{}, args ...interface{}) bool {
	if !assert.ObjectsAreEqualValues(a, b) {
		Log.output(L_INFO, i, Concat(append(args, ": ", a, "not equal to", b)...))
		return true
	}
	return false
}

func ErrorIfNotSame(a, b interface{}, args ...interface{}) bool {
	if !assert.ObjectsAreEqualValues(a, b) {
		Log.output(L_ERROR, e, Concat(append(args, ": ", a, "not equal to", b)...))
		return true
	}
	return false
}

func PanicIfNotSame(a, b interface{}, args ...interface{}) bool {
	if !assert.ObjectsAreEqualValues(a, b) {
		Log.output(L_PANIC, p, Concat(append(args, ": ", a, "not equal to", b)...))
		P(a, b)
		return true
	}
	return false
}

func FatalIfNotSame(a, b interface{}, args ...interface{}) bool {
	if !assert.ObjectsAreEqualValues(a, b) {
		Log.output(L_FATAL, f, Concat(append(args, ": ", a, "not equal to", b)...))
		exit(1)
		return true
	}
	return false
}

func NotNil(err interface{}, args ...interface{}) bool {
	if err != nil {
		Log.output(L_INFO, i, fmt.Sprint(err)+Concat(args...))
		return true
	}
	return false
}

func ErrorIfNotNil(err interface{}, args ...interface{}) bool {
	if err != nil {
		Log.output(L_ERROR, e, fmt.Sprint(err)+Concat(args...))
		return true
	}
	return false
}

func PanicIfNotNil(err interface{}, args ...interface{}) bool {
	if err != nil {
		Log.output(L_PANIC, p, fmt.Sprint(err)+Concat(args...))
		P(err)
		return true
	}
	return false
}

func FatalIfNotNil(err interface{}, args ...interface{}) bool {
	if err != nil {
		Log.output(L_FATAL, f, fmt.Sprint(err)+Concat(args...))
		exit(1)
		return true
	}
	return false
}

func IfFalse(b bool, args ...interface{}) bool {
	if !b {
		Log.output(L_INFO, A, Concat(args...))
		return true
	}
	return false
}

func PanicIfFalse(b bool, args ...interface{}) bool {
	if !b {
		Log.output(L_PANIC, A, Concat(args...))
		P(args...)
		return true
	}
	return false
}

// Assert is a shortcut for FatalIfFalse
func Assert(b bool, args ...interface{}) bool {
	if !b {
		Log.output(L_FATAL, A, Concat(args...))
		exit(1)
		return true
	}
	return false
}

func If(b bool, args ...interface{}) bool {
	if b {
		Log.output(L_INFO, A, Concat(args...))
		return true
	}
	return false
}

func PanicIf(b bool, args ...interface{}) bool {
	if b {
		Log.output(L_PANIC, A, Concat(args...))
		P(args...)
		return true
	}
	return false
}

func FatalIf(b bool, args ...interface{}) bool {
	if b {
		Log.output(L_FATAL, A, Concat(args...))
		exit(1)
		return true
	}
	return false
}
