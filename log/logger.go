package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

//func init() {
//	initConfig()
//}
func GetLogger() *logrus.Logger {
	logger := logrus.New()
	env := os.Getenv("ENV")
	if env == "pro" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		//logName := viper.GetString("app") + ".log"
		//file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY, 0666)
		//if err != nil {
		//	fmt.Println(err)
		//} else {
		//	logger.SetOutput(file)
		//logger.Out = file
		//}

		logger.SetFormatter(&logrus.TextFormatter{DisableColors: false,
			FullTimestamp: false})
	}
	//logger.SetOutput(os.Stdout,file)
	//logger.SetOutput(logger.Writer())
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
