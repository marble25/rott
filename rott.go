package rott

import (
	"os"
	"sync"
)

type Logger struct {
	Filename string `json:"filename" yaml:"filename"`

	size int64
	file *os.File
	mu   sync.Mutex
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		l.file, err = os.OpenFile(l.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return 0, err
		}
	}

	n, err = l.file.Write(p)
	l.size += int64(n)

	return n, err
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.file.Close()
}
