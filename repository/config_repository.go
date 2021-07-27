package repository

import (
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/model"
)

type ConfigRepository interface {
	Query(appId string, cluster string, profile string) model.AppEnabledConfig
	QueryAll() []model.AppEnabledConfig
}

type configRepositoryImpl struct {
	db *gorm.DB
}

func (this configRepositoryImpl) QueryAll() []model.AppEnabledConfig {
	var ret []model.AppEnabledConfig
	this.db.Raw("select ai.app_id,ai.app_key,ai.app_name,ai.department,ai.owner,ci.notify_url,ci.cluster, ci.profile, ci.cid, ci.database, ci.created_at,ci.created_by from  (select * from (select max(created_at) as fcreate_at,cid as fcid from t_app_config group by cid ) f left join (select p.app_id,p.cluster ,p.profile,p.notify_url,p.cid as cid,c.database,c.created_at,c.created_by from t_app_profile p left join t_app_config c on p.cid = c.cid where p.status=1) e on e.cid = f.fcid and e.created_at= f.fcreate_at)  ci left join  t_app_inf ai  on ai.app_id = ci.app_id where ai.status = 1").Scan(&ret)
	return ret
}

func (this configRepositoryImpl) Query(appId string, cluster string, profile string) model.AppEnabledConfig {
	var ret model.AppEnabledConfig
	this.db.Raw("select ai.app_id,ai.app_key,ai.app_name,ai.department,ai.owner,ci.notify_url,ci.cluster, ci.profile, ci.cid, ci.database, ci.created_at,ci.created_by from  (select * from (select max(created_at) as fcreate_at,cid as fcid from t_app_config group by cid ) f left join (select p.app_id,p.cluster ,p.profile,p.notify_url,p.cid as cid,c.database,c.created_at,c.created_by from t_app_profile p left join t_app_config c on p.cid = c.cid where p.status=1) e on e.cid = f.fcid and e.created_at= f.fcreate_at)  ci left join  t_app_inf ai  on ai.app_id = ci.app_id where ai.status = 1 and ai.app_id=? and ci.cluster=? and ci.profile=?", appId, cluster, profile).Scan(&ret)
	return ret
}

func NewConfigRepository(db *gorm.DB) *configRepositoryImpl {
	return &configRepositoryImpl{db: db}
}
