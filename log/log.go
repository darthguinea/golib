package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

var logLevel int

const (
	OFF   = 0
	INFO  = 1
	WARN  = 2
	FATAL = 3
	DEBUG = 4
)

func SetLevel(level int) {
	logLevel = level
}

func Fprintf(w io.Writer, format string, a ...interface{}) {
	fmt.Fprintf(w, format, a)
}

func Print(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format+"\n", a...)
}

func Info(format string, a ...interface{}) (n int, err error) {
	if logLevel > OFF {
		t := time.Now()
		value := "[" + t.Format("2006-01-02 15:04:05") + "] [INFO ] " + format + "\n"
		return fmt.Fprintf(os.Stdout, value, a...)
	}
	return 0, nil
}

func Error(format string, a ...interface{}) (n int, err error) {
	t := time.Now()
	value := "[" + t.Format("2006-01-02 15:04:05") + "] [ERROR] " + format + "\n"
	return fmt.Fprintf(os.Stdout, value, a...)
}

func Warn(format string, a ...interface{}) (n int, err error) {
	if logLevel >= WARN {
		t := time.Now()
		value := "[" + t.Format("2006-01-02 15:04:05") + "] [WARN ] " + format + "\n"
		return fmt.Fprintf(os.Stderr, value, a...)
	}
	return 0, nil
}

func Fatal(v ...interface{}) {
	t := time.Now()
	value := "[" + t.Format("2006-01-02 15:04:05") + "] [FATAL] " + fmt.Sprintln(v...)
	fmt.Printf(value)
}

func Debug(format string, a ...interface{}) (n int, err error) {
	if logLevel >= DEBUG {
		t := time.Now()
		value := "[" + t.Format("2006-01-02 15:04:05") + "] [DEBUG] " + format + "\n"
		return fmt.Fprintf(os.Stderr, value, a...)
	}
	return 0, nil
}
