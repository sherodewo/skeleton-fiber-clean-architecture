package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"runtime"
	"skeleton-fiber-clean-architecture/config"
	"strings"
)

var log = logrus.New()

func InitLogger(cfg *config.LoggerConfig) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    cfg.MaxSize,    // megabytes
		MaxAge:     cfg.MaxAge,     // days
		MaxBackups: cfg.MaxBackups, // number of backups
		Compress:   cfg.Compress,   // compress logs
	})

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.SetReportCaller(true)

	// Set the log level to info by default, can be configured as needed
	log.SetLevel(logrus.InfoLevel)
}

func getCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???", 0, "???"
	}

	funcName := runtime.FuncForPC(pc).Name()
	shortFile := file[strings.LastIndex(file, "/")+1:]

	return shortFile, line, funcName
}

func LogInfo(message string) {
	file, line, function := getCallerInfo()
	log.WithFields(logrus.Fields{
		"file":     file,
		"line":     line,
		"function": function,
	}).Info(message)
}

func LogError(err error) {
	file, line, function := getCallerInfo()
	log.WithFields(logrus.Fields{
		"file":     file,
		"line":     line,
		"function": function,
	}).Error(err)
}
