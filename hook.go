package hookrus

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Hook used to send log to Logentries
type Hook struct {
	writer    io.Writer
	formatter logrus.Formatter
	token     string
}

// NewHook returns a Logrus hook for Logentries Token-based logging
func NewHook(writer io.Writer, formatter logrus.Formatter, token string) Hook {
	return Hook{
		writer:    writer,
		formatter: formatter,
		token:     token,
	}
}

// Fire formats and sends JSON entry to Logentries service
func (h Hook) Fire(e *logrus.Entry) error {
	line, err := h.format(e)
	if err != nil {
		return err
	}
	_, err = h.writer.Write([]byte(h.token + line))
	return err
}

// Levels returns all logrus levels.
func (h Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook Hook) format(entry *logrus.Entry) (string, error) {
	serialized, err := hook.formatter.Format(entry)
	if err != nil {
		return "", err
	}
	str := string(serialized)
	return str, nil
}
