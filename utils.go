package rott

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func shouldBeNil(obj interface{}, t *testing.T) {
	if obj != nil {
		_, file, line, _ := runtime.Caller(2)
		t.Errorf("%s:%d: expected nil, got:%v\n", filepath.Base(file), line, obj)
		t.FailNow()
	}
}

func shouldNotBeNil(obj interface{}, t *testing.T) {
	if obj == nil {
		_, file, line, _ := runtime.Caller(2)
		t.Errorf("%s:%d: expected not nil, got:%v\n", filepath.Base(file), line, obj)
		t.FailNow()
	}
}

func shouldBeEqual(expected, actual interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(2)
		t.Errorf("%s:%d: expected %v (%T), got:%v (%T)\n", filepath.Base(file), line, expected, expected, actual, actual)
		t.FailNow()
	}
}
