package logger

import (
	"os"
)

type LogWriter struct {
	stdOutput  *os.File
	fileOutput *os.File
}

// Write implements io.Writer.
func (mw LogWriter) Write(p []byte) (n int, err error) {
	n, err = mw.stdOutput.Write(p)
	if err != nil {
		return n, err
	}

	n, err = mw.fileOutput.Write(p)
	if err != nil {
		return n, err
	}

	return len(p), nil
}

func NewLoggerWriter(file *os.File) LogWriter {
	return LogWriter{
		stdOutput:  os.Stdout,
		fileOutput: file,
	}
}
