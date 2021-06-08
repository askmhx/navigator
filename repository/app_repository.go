package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type AppRepository interface {
	Create(account *model.ChannelLog) bool
	Update(account *model.ChannelLog) bool
	Query(issNo string, accountNo string) model.ChannelLog
}

type appRepository struct {
	db *gorm.DB
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return &appRepository{db: db}
}

func (this appRepository) Create(account *model.ChannelLog) bool {
	return this.db.Create(account).RowsAffected == 1
}

func (this appRepository) Update(account *model.ChannelLog) bool {
	return this.db.Save(account).RowsAffected == 1
}

func (this appRepository) Query(issNo string, accountNo string) model.ChannelLog {
	var account model.ChannelLog
	//this.db.Where(model.ChannelLog{InstitutionNo: issNo, AccountNo: accountNo}).First(&account)
	return account
}
