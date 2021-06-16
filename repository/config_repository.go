package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type ConfigRepository interface {
	Query(appId string, cluster string, profile string) model.AppConfig
	QueryAll() []model.AppConfig
}

type configRepositoryImpl struct {
	db *gorm.DB
}

func (this configRepositoryImpl) Query(appId string, cluster string, profile string) model.AppConfig {
	var ret model.AppConfig
	this.db.Where(model.AppConfig{AppId: appId, Cluster: cluster, Profile: profile}).First(&ret)
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
