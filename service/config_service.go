package service

import (
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/repository"
)

type ConfigService interface {
	Download(data model.DownloadRequest) model.CommonResult
	RunNotifyTask()
}

type configService struct {
	appRepository   repository.AppRepository
	configRepostiry repository.ConfigRepository
}

func (this configService) RunNotifyTask() {
	panic("implement me")
}

func (this configService) Download(data model.DownloadRequest) model.CommonResult {

	//attach := model.MerchantAttach{
	//	MerchantId: data.MerchantId,
	//	AttachName: data.FileName,
	//	AttachType: "FILE",
	//	AttachPath: data.FileName,
	//	CreatedBy:  "system",
	//}
	//this.appRepository.Save(attach)

	ret := model.CommonResult{
		Code:    model.RESULT_CODE_SUCCESS,
		Message: "upload success",
		Data:    nil,
	}
	return ret
}

func NewConfigService(appRepository repository.AppRepository, configRepostiry repository.ConfigRepository) ConfigService {
	return &configService{
		appRepository:   appRepository,
		configRepostiry: configRepostiry,
	}
}
