package achtung

import (
	"io"
	"log"
	"os"
	"strings"
)

type gameLogger struct {
	file *os.File
}

func newLogger(path string) (*gameLogger, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	return &gameLogger{
		file: f,
	}, nil
}

func (l gameLogger) close() {
	l.file.Close()
}

func (l gameLogger) write(format string, v ...any) {
	if strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	log.Printf(format, v)
}
