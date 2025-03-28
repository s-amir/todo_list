package logs

import (
	"encoding/json"
	"os"
	"time"
)

type Log struct {
	logs []*RichError
}

func (log *Log) Append(err error) {
	//type assertion
	var finalError *RichError
	rErr, ok := err.(*RichError)
	if ok {
		finalError = rErr
	} else if rErr, ok := err.(*SimpleError); ok {
		finalError = &RichError{
			Message:   rErr.Output,
			MetaData:  nil,
			Operation: "UNKNOWN",
			Time:      time.Now(),
		}
	} else {
		finalError = &RichError{Message: err.Error(),
			MetaData:  map[string]string{"type": "unknown"},
			Operation: "Unknown",
			Time:      time.Now(),
		}
	}

	log.logs = append(log.logs, finalError)
}
func (l *Log) Save() {
	//for i, e := range l.Errors {
	//	fmt.Printf("i: %d, operation: %s, message: %s, meta-data: %+v\n",
	//		i, e.Operation, e.Message, e.MetaData)
	//}

	f, _ := os.OpenFile("errors.Log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()

	data, _ := json.Marshal(l.logs)
	f.Write(data)
}
