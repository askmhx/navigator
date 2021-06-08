package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Launch(ctx context.Context, config *AppConfig, router *gin.Engine) {
	log := initLogger(config)
	dbServer := initDatabase(config, log)
	initRoute(config, dbServer,router)
	startService(config, router)
}

func initDatabase(config *AppConfig, logger *Logger) *gorm.DB {
	dbConfig := gorm.Config{Logger: logger.DbLogger()}
	dbConnURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Schema)
	db, err := gorm.Open(mysql.Open(dbConnURL), &dbConfig)
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get raw database connection")
	}
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.ConnMaxLifetime * 1000))
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	return db
}

func startService(config *AppConfig, router *gin.Engine) {
	router.Run(fmt.Sprintf("%s:%d", config.Server.Addr, config.Server.Port))
}
