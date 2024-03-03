package rott

import (
	"os"
	"path/filepath"
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
		err := l.openFile()
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

func (l *Logger) createFile() error {
	// Create the directory if it doesn't exist
	abs, err := filepath.Abs(l.Filename)
	if err != nil {
		return err
	}

	dir := filepath.Dir(abs)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) openFile() error {
	// Create file if it doesn't exist
	info, err := os.Stat(l.Filename)
	if err != nil {
		return l.createFile()
	}

	// Open file
	file, err := os.OpenFile(l.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		// if we can't open the file, try to create it
		return l.createFile()
	}

	l.size = info.Size()
	l.file = file

	return nil
}
