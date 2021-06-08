package service

import (
	"context"
	"rocket.iosxc.com/navigator/v1/app"
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/repository"
)

type AppConfigService interface {
	Upload(data model.UploadRequest) model.CommonResult
	Download(data model.DownloadRequest) model.CommonResult
}

type appConfigService struct {
	ctx   context.Context
	appRepository  repository.AppRepository
	configRepostiry  repository.ConfigRepository
}

func (this appConfigService) Download(data model.UploadRequest) model.CommonResult {

	attach := model.MerchantAttach{
		MerchantId: data.MerchantId,
		AttachName: data.FileName,
		AttachType: "FILE",
		AttachPath: data.FileName,
		CreatedBy:  "system",
	}
	this.appRepository.Save(attach)

	ret := model.CommonResult{
		Code:    model.RESULT_CODE_SUCCESS,
		Message: "upload success",
		Data:    nil,
	}
	return ret
}

func NewFileService(repo repository.MerchantAttachRepository, cache app.CacheManager) AppConfigService {
	return &fileService{
		repo:  repo,
		cache: cache,
	}
}
