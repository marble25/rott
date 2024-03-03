package rott

import "testing"

func TestNewFile(t *testing.T) {
	l := Logger{
		Filename: "test.log",
	}
	defer l.Close()

	b := []byte("test")
	n, err := l.Write(b)
	if err != nil {
		t.Errorf("expected nil, got:%v\n", err)
		t.FailNow()
	}

	if n != len(b) {
		t.Errorf("expected %d, got:%d\n", len(b), n)
		t.FailNow()
	}
}
