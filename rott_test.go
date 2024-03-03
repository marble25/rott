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
	shouldBeNil(err, t)
	shouldBeEqual(len(b), n, t)
}

func TestRotate(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	shouldBeNil(err, t)

	file.Write([]byte("test"))

	l := Logger{
		Filename: "test.log",
	}
	defer l.Close()

	err = l.Rotate()
	shouldBeNil(err, t)

	_, err = os.Stat("test.log.1")
	shouldBeNil(err, t)

	_, err = os.Stat("test.log")
	shouldBeNil(err, t)
}
