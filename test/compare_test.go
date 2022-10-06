package logs_test

import (
	"testing"

	"github.com/conbanwa/logs"
)

var array = []int{1, 2, 3}
var array2 = []float32{1, 2, 3}
var maps = map[string]int{"a": 1, "b": 2, "c": 3}
var maps3 = map[string]int{"1": 1, "2": 2, "3": 3}
var maps2 = map[string]int64{"1": 1, "2": 2, "3": 3}

func TestSameNil(t *testing.T) {
	t.Log(logs.Same(nil, nil))
	t.Log(logs.Same(nil, 1))
	t.Log(logs.Same(1, nil))
}

func TestWrite(t *testing.T) {
	logs.Inline("test")
}

func TestEqual(t *testing.T) {
	logs.NotSame(2+3, '5', "2+3 is not 5")
	t.Log(logs.Same(array, array2))
	t.Log(logs.Same(maps2, maps3))
	t.Log(logs.Same(maps, maps2))
	t.Log(logs.Same(maps, maps3))
	t.Log(logs.Same(array, maps))
	t.Log(logs.Same(array, maps2))
	t.Log(logs.Same(array, maps3))
	t.Log(logs.Same(array2, maps))
	t.Log(logs.Same(array2, maps2))
	t.Log(logs.Same(array2, maps3))
}

func TestEqualNumber(t *testing.T) {
	t.Log(logs.Same(1, float32(1)))
	t.Log(logs.Same("23", 23.0))
	t.Log(logs.Same("23.0", 23.0))
	t.Log(logs.Same("23.0", 23))
	t.Log(logs.Same("23.3", 23.3))
	t.Log(logs.Same("46.3", 23.1+23.2))
	t.Log(logs.Same("46.003", 23.002+23.001))
	t.Log(logs.Same("46.000003", 23.000002+23.000001))
	t.Log(logs.Same("46.0000000000000000000003", 23.0000000000000000000002+23.0000000000000000000001))
	t.Log(logs.Same(46.3, 23.1+23.2))
	t.Log(logs.Same(46.0000000000003, 23.0000000000002+23.0000000000001))
	t.Log(logs.Same(46.0000000000000000000003, 23.0000000000000000000002+23.0000000000000000000001))

}
