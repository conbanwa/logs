package logs

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Equal(a, b interface{}) bool {
	if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() {
		if a == b {
			return true
		}
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

func IfDiffer(a, b interface{}, args ...interface{}) bool {
	if !Equal(a, b) {
		Log.output(ERROR, E, Concat(append(args, a, "not equal to", b)...))
		return true
	}
	return false
}
func PanicIfDiffer(a, b interface{}, args ...interface{}) {
	if !Equal(a, b) {
		Log.output(PANIC, P, Concat(append(args, a, "not equal to", b)...))
		Panic(a, b)
	}
}
func ExitIfDiffer(a, b interface{}, args ...interface{}) {
	if !Equal(a, b) {
		Log.output(FATAL, F, Concat(append(args, a, "not equal to", b)...))
		os.Exit(1)
	}
}
func AllEqual(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if !Equal(args[i], args[j]) {
				Log.output(ERROR, E, Concat(args[i], "not equal to", args[j]))
				return false
			}
		}
	}
	return true
}
func AllDistinct(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if Equal(args[i], args[j]) {
				Log.output(ERROR, E, Concat(args[i], "is equal to", args[j]))
				return false
			}
		}
	}
	return true
}
func IfNuance(a, b interface{}, args ...interface{}) bool {
	if a != b {
		Log.output(ERROR, E, Concat(append(args, a, "not exactly equal to", b)...))
		return true
	}
	return false
}
func PanicIfNuance(a, b interface{}, args ...interface{}) {
	if a != b {
		Log.output(PANIC, P, Concat(append(args, a, "not exactly equal to", b)...))
		Panic(a, b)
	}
}
func ExitIfNuance(a, b interface{}, args ...interface{}) {
	if a != b {
		Log.output(FATAL, F, Concat(append(args, a, "not exactly equal to", b)...))
		os.Exit(1)
	}
}
func ToFloat64(a interface{}) (f float64, err error) {
	switch a := a.(type) {
	case int:
		f = float64(a)
	case int8:
		f = float64(a)
	case int16:
		f = float64(a)
	case int32:
		f = float64(a)
	case int64:
		f = float64(a)
	case uint:
		f = float64(a)
	case uint8:
		f = float64(a)
	case uint16:
		f = float64(a)
	case uint32:
		f = float64(a)
	case uint64:
		f = float64(a)
	case float32:
		f = float64(a)
	case float64:
		f = a
	case string:
		f, err = strconv.ParseFloat(a, 64)
	default:
		err = errors.New("not a number")
	}
	return
}
