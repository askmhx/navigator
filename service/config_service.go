package service

import (
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/repository"
	"rocket.iosxc.com/navigator/v1/util"
)

type ConfigService interface {
	Download(data model.ConfigRequest) model.CommonResult
	QueryAll() []model.AppEnabledConfig
}

type configService struct {
	configRepostiry repository.ConfigRepository
}

func (this *configService) QueryAll() []model.AppEnabledConfig {
	return this.configRepostiry.QueryAll()
}

func (this *configService) Download(data model.ConfigRequest) model.CommonResult {
	appConfig := this.configRepostiry.Query(data.AppId, data.Cluster, data.Profile)
	dataMap := map[string]string{}
	dataMap["AppId"] = data.AppId
	dataMap["Cluster"] = data.Cluster
	dataMap["Profile"] = data.Profile
	if util.GetSign(dataMap, appConfig.AppKey) != data.Sign {
		ret := model.CommonResult{
			Code:    model.RESULT_CODE_FAIL,
			Message: "sign error",
			Data:    nil,
		}
		return ret
	}
	config := util.AESEncrypt(appConfig.Config, appConfig.AppKey)
	createAt := appConfig.CreatedAt.Format(util.DATE_FORMAT_YMDHMS)
	dataMap["Config"] = config
	dataMap["NotifyUrl"] = appConfig.NotifyUrl
	dataMap["CreateAt"] = createAt
	dataMap["CreateBy"] = appConfig.CreatedBy
	sign := util.GetSign(dataMap, appConfig.AppKey)
	response := model.ConfigResponse{AppId: appConfig.AppId, Profile: appConfig.Profile, Cluster: appConfig.Cluster, Config: config, CreateAt: createAt, CreateBy: appConfig.CreatedBy, Sign: sign}
	ret := model.CommonResult{
		Code:    model.RESULT_CODE_SUCCESS,
		Message: "success",
		Data:    response,
	}
	return ret
}

func NewConfigService(configRepostiry repository.ConfigRepository) *configService {
	return &configService{
		configRepostiry: configRepostiry,
	}
}
