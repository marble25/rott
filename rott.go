package rott

import (
	"os"
	"path"
	"path/filepath"
	"sync"
)

type Logger struct {
	Filename          string       `json:"filename" yaml:"filename"`
	BackupOption      BackupOption `json:"backup_condition" yaml:"backup_condition"`
	BackupFilePattern string       `json:"backup_file_pattern" yaml:"backup_file_pattern"`

	size int64
	file *os.File
	mu   sync.Mutex
}

type BackupOption int

const (
	NoBackup BackupOption = iota
	ByTime
	BySize
)

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

func (l *Logger) Rotate() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.rotate()
}

func (l *Logger) filename() string {
	if l.Filename != "" {
		return l.Filename
	}
	tmpDir := os.TempDir()
	return path.Join(tmpDir, "rott.log")
}

func (l *Logger) createFile() error {
	// Create the directory if it doesn't exist
	abs, err := filepath.Abs(l.filename())
	if err != nil {
		return err
	}

	dir := filepath.Dir(abs)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.OpenFile(l.filename(), os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	info, err := os.Stat(l.filename())
	if err != nil {
		return err
	}

	l.size = info.Size()
	l.file = file

	return nil
}

func (l *Logger) openFile() error {
	// Create file if it doesn't exist
	info, err := os.Stat(l.filename())
	if err != nil {
		return l.createFile()
	}

	// Open file
	file, err := os.OpenFile(l.filename(), os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		// if we can't open the file, try to create it
		return l.createFile()
	}

	l.size = info.Size()
	l.file = file

	return nil
}

func (l *Logger) backupFile() error {
	// If backup is not enabled, return
	if l.BackupOption == NoBackup {
		return nil
	}

	err := os.Rename(l.filename(), l.filename()+".1")
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) rotate() error {
	// Close the current file
	l.file.Close()

	// Backup current file
	err := l.backupFile()
	if err != nil {
		return err
	}

	// Open a new file
	err = l.openFile()
	if err != nil {
		return err
	}

	return nil
}
