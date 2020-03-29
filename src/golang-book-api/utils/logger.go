package utils

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

// ASCII colors numbers
const (
	red    = 31
	grey   = 37
	blue   = 36
	yellow = 33
)

// loggerTextFormatter object for formatting text in chaincode logs
type loggerTextFormatter struct {
	logrus.TextFormatter
}

// NewLogger function returns Logger object
func NewLogger(logLevel logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &loggerTextFormatter{logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	},
	}
	logger.Level = logLevel
	return logger
}

// Format logger output
func (f *loggerTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logLevel := "[" + strings.ToUpper(entry.Level.String()) + "]"
	return []byte(
		fmt.Sprintf(" %s \x1b[%dm%s\x1b[0m %s\n",
			entry.Time.Format(f.TimestampFormat),
			setLogLevelColor(entry),
			logLevel,
			entry.Message,
		),
	), nil
}

func setLogLevelColor(entry *logrus.Entry) int {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = blue
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = grey
	}
	return levelColor
}
