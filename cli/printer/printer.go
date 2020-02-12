package printer

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var Printer printer

type printFormatter struct {}

func (f *printFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

type printer interface {
	PrintJSON(interface{}) error
	PrintText(interface{}) error
	PrintYAML(interface{}) error
}

type printLogger struct {
	*log.Logger
}

func Init() {
	logger := &printLogger{log.New()}
	logger.SetFormatter(&printFormatter{})

	Printer = logger
}

func (l *printLogger) PrintJSON(j interface{}) error {
	log.Debugf("%s\n", j)
	res, err := json.Marshal(j)
	if err != nil {
		return fmt.Errorf("Cannot print JSON: %+v", err)
	}

	l.Printf("[JSON] %s\n", res)

	return nil
}

func (l *printLogger) PrintText(t interface{}) error{
	return nil
}

func (l *printLogger) PrintYAML(y interface{}) error{
	return nil
}