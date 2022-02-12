package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	DebugLevel = "Debug"
)

var Logger = logrus.New()

type MyTextFormatter struct{}

func (t *MyTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timeStr := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", entry.Time.Year(), entry.Time.Month(), entry.Time.Day(), entry.Time.Hour(), entry.Time.Minute(), entry.Time.Second())
	filePath := strings.Split(entry.Caller.File, "/")
	direct := filePath[len(filePath)-2]
	name := filePath[len(filePath)-1]
	result := fmt.Sprintf("%s %s/%s:%d [%s] %s", timeStr, direct, name, entry.Caller.Line, entry.Level, entry.Message)
	return append([]byte(result), '\n'), nil
}

func init() {
	Logger.SetFormatter(&MyTextFormatter{})
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.InfoLevel)
	level := os.Getenv("LOG_LEVEL")
	if level == DebugLevel {
		Logger.SetLevel(logrus.DebugLevel)
	}
	Logger.SetReportCaller(true)
	Logger.Infof("log level: %v", level)
	// write to log file
	//fileName := fmt.Sprintf("logrus-%d.log", time.Now().Unix())
	//file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	Logger.Out = file
	//} else {
	//	Logger.Info("Failed to log to file, using default stderr")
	//}
}
