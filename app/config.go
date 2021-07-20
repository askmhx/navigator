package app

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	Home    string
	IrisYML string

	Server struct {
		Port    uint
		Addr    string
		CharSet string
		Mode    string
	}
	Database struct {
		User            string
		Password        string
		Schema          string
		Host            string
		Port            uint
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime int
	}
	Log struct {
		FilePath    string
		RotateType  string
		RotateValue string
		RotateCount uint
		Level       string
	}
	Task struct {
		NotifyDuration int
	}
}

var config *AppConfig

func InitConfig(configPath string) *AppConfig {
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
