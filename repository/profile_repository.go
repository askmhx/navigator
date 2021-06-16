package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type ProfileRepository interface {
	Query(appId string, cluster string, profile string) model.AppProfile
	QueryAll() []model.AppProfile
}

type profileRepositoryImpl struct {
	db *gorm.DB
}

func (this profileRepositoryImpl) Query(appId string, cluster string, profile string) model.AppProfile {
	var ret model.AppProfile
	this.db.Where(model.AppProfile{AppId: appId, Cluster: cluster, Profile: profile}).First(&ret)
	return ret
}

func (this profileRepositoryImpl) QueryAll() []model.AppProfile {
	var ret []model.AppProfile
	this.db.Where(model.AppProfile{}).Find(&ret)
	return ret
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepositoryImpl{db: db}
}
