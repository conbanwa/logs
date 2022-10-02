package logs

import (
	"os"
)

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
		os.Exit(1)
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
		os.Exit(1)
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
		os.Exit(1)
	}
}
