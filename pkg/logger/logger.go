package logger

import (
	"os"
	"time"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New();

func Init(level string) {
	Log.Out = os.Stdout;
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
	});

	lvl, err := logrus.ParseLevel(level);
	if err != nil {
		lvl = logrus.InfoLevel;
	}
	Log.SetLevel(lvl);
}