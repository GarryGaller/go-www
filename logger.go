package www

import (
	"fmt"
	"io"
    "log"
    "os"
)

var (
	defaultLogger = log.New(os.Stderr, "", log.LstdFlags)
)


type Logger interface {
	Printf(string, ...interface{})
}

type LeveledLogger interface {
	Error(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
}

type StandardLogger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Flags() int
	Output(calldepth int, s string) error
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Prefix() string
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	SetFlags(flag int)
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	Writer() io.Writer
}

func (cl *StandardClient) Log() Logger {
	if cl.Logger != nil {
		switch v := cl.Logger.(type) {
		case Logger:
			return v
		default:
			panic(fmt.Sprintf(
				"invalid logger type passed, must be Logger, was %T",
				cl.Logger),
			)
		}
	}
	return nil
}

func (cl *StandardClient) LLog() LeveledLogger {
	if cl.Logger != nil {
		switch v := cl.Logger.(type) {
		case LeveledLogger:
			return v
		default:
			panic(fmt.Sprintf(
				"invalid logger type passed, must be LeveledLogger, was %T",
				cl.Logger),
			)
		}
	}
	return nil
}

func (cl *StandardClient) SLog() StandardLogger {
	if cl.Logger != nil {
		switch v := cl.Logger.(type) {
		case StandardLogger:
			return v
		default:
			panic(fmt.Sprintf(
				"invalid logger type passed, must be StandardLogger, was %T",
				cl.Logger),
			)
		}
	}
	return nil
}
