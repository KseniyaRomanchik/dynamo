package printer

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	log "github.com/sirupsen/logrus"
)

var Printer printer

type printFormatter struct {}

func (f *printFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

type printer interface {
	PrintJSON(interface{}) error
	PrintText(fmt.Stringer) error
	PrintYAML(interface{}) error
	PrintAWSErr(error) error
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
	res, err := json.Marshal(j)
	if err != nil {
		return fmt.Errorf("Cannot print JSON: %+v", err)
	}

	l.Printf("[JSON] %s\n", res)

	return nil
}

func (l *printLogger) PrintText(t fmt.Stringer) error{
	l.Printf("[TEXT] %s\n", t.String())
	return nil
}

func (l *printLogger) PrintYAML(y interface{}) error{
	return nil
}

func (l *printLogger) PrintAWSErr(e error) error{
	aerr, ok := e.(awserr.Error)
	if !ok {
		return fmt.Errorf("cannot convert to aws error: %+v", e)
	}

	l.Errorf("[ERROR] aws: %s - %s\n", aerr.Code(), aerr.Message())
	return nil
}