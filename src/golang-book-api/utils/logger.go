package utils

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

type loggerTextFormatter struct {
	logrus.TextFormatter
}

// Format logger output
func (f *loggerTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf(" %s \x1b[%dm%s\x1b[0m %s\n", entry.Time.Format(f.TimestampFormat), setLogLevelColor(entry), "["+strings.ToUpper(entry.Level.String())+"]", entry.Message)), nil
}

func setLogLevelColor(entry *logrus.Entry) int {
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return 36
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return 31
	default:
		return 37
	}
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
