package settings

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func SetupLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableLevelTruncation: false})
	file, err := os.OpenFile("/var/log/nightwatch.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err == nil {
		Log.Out = file
	} else {
		fmt.Println(err)
		Log.Info("Failed to log to file, using default stderr")
	}
}
