package service

import (
	"fmt"
	"reflect"
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/repository"
	"time"
)

type ConfigService interface {
	Download(data model.ConfigRequest) model.CommonResult
	RunNotifyTask()
}

var configMap = new(map[string]model.AppConfig)
var appMap = new(map[string]model.AppEnabledConfig)

type configService struct {
	appRepository   repository.AppRepository
	configRepostiry repository.ConfigRepository
}

func (this configService) RunNotifyTask() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for t := range ticker.C {
		fmt.Printf("Timer Start AT:%s\n", t)
		for _, m := range this.configRepostiry.QueryAll() {
			oldConfig := configMap[m.AppId]
			if oldConfig != nil&!reflect.DeepEqual(oldConfig, m) {
				fmt.Printf("Find New Config For:%s Cluster:%s Profile:%s CreatedAt:%s\n", m.AppId, m.Cluster, m.Profile, m.CreatedAt)
				notify()
			} else {
				configMap[m.AppId] = oldConfig
			}
		}
	}
}

func notify() {

}

func (this configService) Download(data model.ConfigRequest) model.CommonResult {

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
