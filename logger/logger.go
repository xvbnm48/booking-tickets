package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)
}

func getCallerInfo() (string, int) {
	_, file, line, ok := runtime.Caller(2) // 2 levels up to get the caller of the log function
	if !ok {
		return "unknown", 0
	}
	// Use filepath.Base to get only the file name (not the full path)
	return filepath.Base(file), line
}

// Info logs an informational message with file and line info
func Info(msg string) {
	file, line := getCallerInfo()
	Log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Info(msg)
}

// Error logs an error message with file and line info
func Error(msg string, err error) {
	file, line := getCallerInfo()
	Log.WithFields(logrus.Fields{
		"file":  file,
		"line":  line,
		"error": err.Error(),
	}).Error(msg)
}

// Debug logs a debug message with file and line info
func Debug(msg string) {
	file, line := getCallerInfo()
	Log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Debug(msg)
}
