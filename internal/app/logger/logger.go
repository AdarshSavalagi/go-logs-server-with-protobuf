package logger

import "github.com/sirupsen/logrus"

func InitLogger()*logrus.Logger{
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}