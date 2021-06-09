package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type ConfigRepository interface {
	Create(account *model.ChannelLog) bool
	Update(account *model.ChannelLog) bool
	Query(issNo string, accountNo string) model.ChannelLog
}

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) ConfigRepository {
	return &configRepository{db: db}
}

func (this configRepository) Create(account *model.ChannelLog) bool {
	return this.db.Create(account).RowsAffected == 1
}

func (this configRepository) Update(account *model.ChannelLog) bool {
	return this.db.Save(account).RowsAffected == 1
}

func (this configRepository) Query(issNo string, accountNo string) model.ChannelLog {
	var account model.ChannelLog
	//this.db.Where(model.ChannelLog{InstitutionNo: issNo, AccountNo: accountNo}).First(&account)
	return account
}
