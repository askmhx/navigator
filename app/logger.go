package app

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
)

func initLogger(config *AppConfig) *Logger {

	cPath, _ := filepath.Abs(fmt.Sprintf("%s%s", config.Home, config.Log.FilePath))
	logPath := fmt.Sprintf("%s%s", cPath, os.Executable())
	file, err := os.Open(logPath)
	if err != nil {
		file, _ = os.Create(logPath)
	}
	defer file.Close()
	return &Logger{LogFile: file}
}

type Logger struct {
	LogFile *os.File
}

func (this *Logger) DbLogger() logger.Interface {
	logConfig := logger.Config{LogLevel: logger.Info, Colorful: false}
	return logger.New(log.New(this.LogFile, "DbLog:", log.Ldate|log.Ltime|log.Lshortfile), logConfig)
}

func (this *Logger) SysLogger() *log.Logger {
	return log.New(this.LogFile, "SysLog:", log.Ldate|log.Ltime|log.Lshortfile)
}
