package rott

import (
	"os"
	"testing"
)

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

func TestRotate(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		t.Errorf("expected nil, got:%v\n", err)
		t.FailNow()
	}
	file.Write([]byte("test"))

	l := Logger{
		Filename: "test.log",
	}
	defer l.Close()

	err = l.Rotate()
	if err != nil {
		t.Errorf("expected nil, got:%v\n", err)
		t.FailNow()
	}

	_, err = os.Stat("test.log.1")
	if err != nil {
		t.Errorf("expected nil, got:%v\n", err)
		t.FailNow()
	}

	_, err = os.Stat("test.log")
	if err != nil {
		t.Errorf("expected nil, got:%v\n", err)
		t.FailNow()
	}
}
