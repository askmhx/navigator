package app

import (
	"fmt"
	"reflect"
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/service"
	"rocket.iosxc.com/navigator/v1/util"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type taskManager struct {
	service   service.ConfigService
	configMap map[string]model.AppEnabledConfig
}

var tm *taskManager

func initTaskManager(service service.ConfigService) *taskManager {
	if tm == nil {
		lock.Lock()
		defer lock.Unlock()
		if tm == nil {
			tm = &taskManager{service: service, configMap: map[string]model.AppEnabledConfig{}}
		}
	}
	return tm
}

func getTaskManager() *taskManager {
	return tm
}

func (this *taskManager) notify(m model.AppEnabledConfig) {
	if !util.IsValidUrl(m.NotifyUrl) {
		fmt.Printf("URL ERROR For:%s Cluster:%s Profile:%s CreatedAt:%s\n", m.AppId, m.Cluster, m.Profile, m.CreatedAt)
	}
	datas := map[string]string{}
	datas["AppId"] = m.AppId
	datas["Cluster"] = m.Cluster
	datas["Profile"] = m.Profile
	datas["Config"] = util.AESEncrypt(m.Config, m.AppKey)
	datas["CreateAt"] = m.CreatedAt.Format(util.DATE_FORMAT_YMDHMS)
	datas["CreateBy"] = m.CreatedBy
	util.HttpPost(m.NotifyUrl, datas, m.AppKey)
}

func (this *taskManager) runTask(period int) {
	if period <= 0 {
		return
	}
	ticker := time.NewTicker(time.Second * time.Duration(period))
	for t := range ticker.C {
		fmt.Printf("Timer Start AT:%s\n", t)
		for _, m := range this.service.QueryAll() {
			oldConfig, exist := this.configMap[m.Cid]
			if exist && !reflect.DeepEqual(oldConfig, m) {
				fmt.Printf("Find New Config For:%s Cluster:%s Profile:%s CreatedAt:%s\n", m.AppId, m.Cluster, m.Profile, m.CreatedAt)
				this.notify(oldConfig)
			} else {
				this.configMap[m.Cid] = oldConfig
			}
		}
	}
}
