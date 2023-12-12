package common

import (
	"os"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
	"github.com/sirupsen/logrus"
)

type MyLogger interface {
	InitializeLogger() error // membuat sebuah file logger
	LogInfo(requestLog utilsmodel.RequestLog)
	LogWarn(requestLog utilsmodel.RequestLog)
	LogFatal(requestLog utilsmodel.RequestLog)
}

type myLogger struct {
	cfg config.LogConfig
	log *logrus.Logger
}

func (m *myLogger) InitializeLogger() error {
	file, err := os.OpenFile(m.cfg.LogerFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	m.log = logrus.New()
	m.log.SetOutput(file)
	return nil
}

func (m *myLogger) LogFatal(requestLog utilsmodel.RequestLog) {
	m.log.Fatal(requestLog)
}

func (m *myLogger) LogInfo(requestLog utilsmodel.RequestLog) {
	m.log.Info(requestLog)
}

func (m *myLogger) LogWarn(requestLog utilsmodel.RequestLog) {
	m.log.Warn(requestLog)
}

func NewMyLogger(cfg config.LogConfig) MyLogger {
	return &myLogger{cfg: cfg}
}
