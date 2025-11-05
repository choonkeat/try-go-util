package logrusutil

import (
	"github.com/choonkeat/try-go-util/pkg/logutil"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
}

func (l *LogrusLogger) Log(s string) {
	logrus.Info(s)
}


var _ logutil.Logger = &LogrusLogger{}
