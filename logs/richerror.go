package logs

import "time"

type RichError struct {
	Message   string
	MetaData  map[string]string
	Operation string
	Time      time.Time
}

func (err *RichError) Error() string {
	return err.Message
}
