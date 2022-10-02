package logs

import (
	"fmt"
	"os"
	"strconv"
)

func Sprint(args ...interface{}) string {
	Log.output(INFO, I, Concat(args...))
	return Concat(args...)
}

func Same(a, b interface{}) bool {
	if a == b {
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
		err = fmt.Errorf("not a number")
	}
	return
}

func Inline(args ...interface{}) {
	str := Concat(args...)
	//output to stdout
	b := []byte(str)
	os.Stderr.Write(b)
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
		os.Exit(1)
	}
}
