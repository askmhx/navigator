package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type ConfigRepository interface {
	Query(cid string) model.AppConfig
	QueryAll() []model.AppConfig
}

type configRepositoryImpl struct {
	db *gorm.DB
}

func (this configRepositoryImpl) Query(cid string) model.AppConfig {
	var ret model.AppConfig
	this.db.Where(model.AppConfig{Cid: cid}).First(&ret)
	return ret
}

func (this configRepositoryImpl) QueryAll() []model.AppConfig {
	var ret []model.AppConfig
	this.db.Where(model.AppConfig{}).Find(&ret)
	return ret
}

func NewConfigRepository(db *gorm.DB) ConfigRepository {
	return &configRepositoryImpl{db: db}
}
