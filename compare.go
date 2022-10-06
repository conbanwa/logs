package logs

import (
	"fmt"
	"reflect"
)

func Same(a, b interface{}) bool {
	if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() && a == b {
		return true
	}
	fa, era := ToFloat64(a)
	fb, erb := ToFloat64(b)
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

func IfSame(a, b interface{}, args ...interface{}) bool {
	if Same(a, b) {
		Log.output(INFO, A, Concat(append(args, ":", a)...))
		return true
	}
	return false
}

func NotSame(a, b interface{}, args ...interface{}) bool {
	if !Same(a, b) {
		Log.output(ERROR, E, Concat(append(args, ": ", a, "not equal to", b)...))
		return true
	}
	return false
}

func PanicNotSame(a, b interface{}, args ...interface{}) {
	if !Same(a, b) {
		Log.output(PANIC, P, Concat(append(args, ": ", a, "not equal to", b)...))
		Panic(a, b)
	}
}

func ExitNotSame(a, b interface{}, args ...interface{}) {
	if !Same(a, b) {
		Log.output(FATAL, F, Concat(append(args, ": ", a, "not equal to", b)...))
		exit(1)
	}
}

func NotNil(err error, args ...interface{}) bool {
	if err != nil {
		Log.output(ERROR, E, err.Error()+":"+Concat(args...))
		return true
	}
	return false
}

func AppendNotNil(err *error, args ...interface{}) bool {
	if *err != nil {
		*err = fmt.Errorf("%v%v", *err, Concat(args...))
		Log.output(ERROR, E, (*err).Error()+":"+Concat(args...))
		return true
	}
	return false
}

func PanicNotNil(err error, args ...interface{}) {
	if err != nil {
		Log.output(PANIC, P, err.Error()+":"+Concat(args...))
		Panic(err.Error() + Concat(args...))
	}
}

func ExitNotNil(err error, args ...interface{}) {
	if err != nil {
		Log.output(FATAL, F, err.Error()+":"+Concat(args...))
		exit(1)
	}
}

func Uniform(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if !Same(args[i], args[j]) {
				Log.output(ERROR, C, Concat(args[i], "not equal to", args[j]))
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
				Log.output(ERROR, C, Concat(args[i], "is equal to", args[j]))
				return false
			}
		}
	}
	return true
}

func NotValueAndTypeEqual(a, b interface{}, args ...interface{}) bool {
	if a != b {
		Log.output(ERROR, E, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		return true
	}
	return false
}

func PanicNotValueAndTypeEqual(a, b interface{}, args ...interface{}) {
	if a != b {
		Log.output(PANIC, P, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		Panic(a, b)
	}
}

func ExitNotValueAndTypeEqual(a, b interface{}, args ...interface{}) {
	if a != b {
		Log.output(FATAL, F, Concat(append(args, ": ", a, "not exactly equal to", b)...))
		exit(1)
	}
}

func IfFalse(b bool, args ...interface{}) bool {
	if !b {
		Log.output(INFO, A, Concat(args...))
		return true
	}
	return false
}

func PanicIfFalse(b bool, args ...interface{}) {
	if !b {
		Log.output(PANIC, A, Concat(args...))
		Panic(args...)
	}
}

func Assert(b bool, args ...interface{}) {
	if !b {
		Log.output(FATAL, A, Concat(args...))
		exit(1)
	}
}
