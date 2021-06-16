package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type AppRepository interface {
	Query(appId string) model.AppInf
	QueryAll() []model.AppInf
}

type appRepositoryImpl struct {
	db *gorm.DB
}

func (this appRepositoryImpl) Query(appId string) model.AppInf {
	var ret model.AppInf
	this.db.Where(model.AppInf{AppId: appId}).First(&ret)
	return ret
}

func (this appRepositoryImpl) QueryAll() []model.AppInf {
	var rets []model.AppInf
	this.db.Find(&rets)
	return rets
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return &appRepositoryImpl{db: db}
}
