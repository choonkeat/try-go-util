package logutil

import "fmt"

type Logger interface {
	Log(s string)
}

type FmtLogger struct {
}

func (l *FmtLogger) Log(s string) {
	fmt.Printf("[fmt] %s\n", s)
}
