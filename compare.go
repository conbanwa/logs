package logs

import (
	"fmt"
	"reflect"
)

func Same(a, b interface{}) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() && a == b {
		return true
	}
	fa, era := ParseFloat64(a)
	fb, erb := ParseFloat64(b)
	if era == nil && erb == nil {
		return fa == fb
	}
	if fmt.Sprint(a) == fmt.Sprint(b) {
		return true
	}
	if fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b) {
		return true
	}
	return false
}

func IfSame(d, b interface{}, args ...interface{}) bool {
	if Same(d, b) {
		Log.output(L_INFO, a, Concat(append(args, ":", a)...))
		return true
	}
	return false
}

func NotSame(a, b interface{}, args ...interface{}) bool {
	if !Same(a, b) {
		Log.output(L_INFO, i, Concat(append(args, ": ", a, "not equal to", b)...))
		return true
	}
	return false
}

func ErrorIfNotSame(a, b interface{}, args ...interface{}) bool {
	if !Same(a, b) {
		Log.output(L_ERROR, e, Concat(append(args, ": ", a, "not equal to", b)...))
		return true
	}
	return false
}

func PanicIfNotSame(a, b interface{}, args ...interface{}) bool {
	if !Same(a, b) {
		Log.output(L_PANIC, p, Concat(append(args, ": ", a, "not equal to", b)...))
		Panic(a, b)
		return true
	}
	return false
}

func FatalIfNotSame(a, b interface{}, args ...interface{}) bool {
	if !Same(a, b) {
		Log.output(L_FATAL, f, Concat(append(args, ": ", a, "not equal to", b)...))
		exit(1)
		return true
	}
	return false
}

func AppendIfNotNil(err *error, args ...interface{}) bool {
	if *err != nil {
		*err = fmt.Errorf("%v%v", *err, Concat(args...))
		Log.output(L_ERROR, e, (*err).Error()+":"+Concat(args...))
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
		Panic(err)
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

func Uniform(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if !Same(args[i], args[j]) {
				Log.output(L_ERROR, c, Concat(args...))
				return false
			}
		}
	}
	return true
}

func Distinct(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if Same(args[i], args[j]) {
				Log.output(L_ERROR, c, Concat(args...))
				return false
			}
		}
	}
	return true
}

func ValueOrTypeNotEqual(a, b interface{}, args ...interface{}) bool {
	if a != b {
		Log.output(L_INFO, i, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		return true
	}
	return false
}

func ErrorIfValueOrTypeNotEqual(a, b interface{}, args ...interface{}) bool {
	if a != b {
		Log.output(L_ERROR, e, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		return true
	}
	return false
}

func PanicIfValueOrTypeNotEqual(a, b interface{}, args ...interface{}) bool {
	if a != b {
		Log.output(L_PANIC, p, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		Panic(a, b)
		return true
	}
	return false
}

func FatalIfValueOrTypeNotEqual(a, b interface{}, args ...interface{}) bool {
	if a != b {
		Log.output(L_FATAL, f, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		exit(1)
		return true
	}
	return false
}

func IfFalse(b bool, args ...interface{}) bool {
	if !b {
		Log.output(L_INFO, a, Concat(args...))
		return true
	}
	return false
}

func PanicIfFalse(b bool, args ...interface{}) bool {
	if !b {
		Log.output(L_PANIC, a, Concat(args...))
		Panic(args...)
		return true
	}
	return false
}

// Assert is a shortcut for FatalIfFalse
func Assert(b bool, args ...interface{}) bool {
	if !b {
		Log.output(L_FATAL, a, Concat(args...))
		exit(1)
		return true
	}
	return false
}

func If(b bool, args ...interface{}) bool {
	if b {
		Log.output(L_INFO, a, Concat(args...))
		return true
	}
	return false
}

func PanicIf(b bool, args ...interface{}) bool {
	if b {
		Log.output(L_PANIC, a, Concat(args...))
		Panic(args...)
		return true
	}
	return false
}

func FatalIf(b bool, args ...interface{}) bool {
	if b {
		Log.output(L_FATAL, a, Concat(args...))
		exit(1)
		return true
	}
	return false
}
