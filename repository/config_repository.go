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
	this.db.Raw("select AI.APP_ID,AI.APP_KEY,AI.APP_NAME,AI.DEPARTMENT,AI.OWNER,CI.NOTIFY_URL,CI.CLUSTER, CI.PROFILE, CI.CID, CI.CONFIG, CI.CREATED_AT,CI.CREATED_BY from  (select * from (select MAX(CREATED_AT) as FCREATE_AT,CID as FCID from T_APP_CONFIG group by CID ) F left join (select P.APP_ID,P.CLUSTER ,P.PROFILE,P.NOTIFY_URL,P.CID AS CID,C.CONFIG,C.CREATED_AT,C.CREATED_BY from T_APP_PROFILE P left join T_APP_CONFIG C on P.CID = C.CID where P.STATUS=1) E on E.CID = F.FCID and E.CREATED_AT= F.FCREATE_AT)  CI left join  T_APP_INF AI  on AI.APP_ID = CI.APP_ID where AI.STATUS = 1").Scan(&ret)
	return ret
}

func (this configRepositoryImpl) Query(appId string, cluster string, profile string) model.AppEnabledConfig {
	var ret model.AppEnabledConfig
	this.db.Raw("select AI.APP_ID,AI.APP_KEY,AI.APP_NAME,AI.DEPARTMENT,AI.OWNER,CI.NOTIFY_URL,CI.CLUSTER, CI.PROFILE, CI.CID, CI.CONFIG, CI.CREATED_AT,CI.CREATED_BY from  (select * from (select MAX(CREATED_AT) as FCREATE_AT,CID as FCID from T_APP_CONFIG group by CID ) F left join (select P.APP_ID,P.CLUSTER ,P.PROFILE,P.NOTIFY_URL,P.CID AS CID,C.CONFIG,C.CREATED_AT,C.CREATED_BY from T_APP_PROFILE P left join T_APP_CONFIG C on P.CID = C.CID where P.STATUS=1) E on E.CID = F.FCID and E.CREATED_AT= F.FCREATE_AT)  CI left join  T_APP_INF AI  on AI.APP_ID = CI.APP_ID where AI.STATUS = 1 and AI.APP_ID = ? and CI.CLUSTER=? and CI.PROFILE=?", appId, cluster, profile).Scan(&ret)
	return ret
}

func NewConfigRepository(db *gorm.DB) *configRepositoryImpl {
	return &configRepositoryImpl{db: db}
}
