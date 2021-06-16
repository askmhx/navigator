package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/service"
	"sync"
)

type DispatcherCtrl struct {
	ConfigService service.ConfigService
}

func NewDispatcherCtrl() DispatcherCtrl {
	return DispatcherCtrl{}
}

var once sync.Once

func (this *DispatcherCtrl) Download(context *gin.Context) {
	var requestData model.ConfigRequest
	err := context.ShouldBind(&requestData)
	if err != nil {
		fmt.Println("error happened:", err)
	}
	result := this.ConfigService.Download(requestData)
	this.response(result.Data, context)
	once.Do(func() {
		go this.ConfigService.RunNotifyTask()
	})
}

func (this *DispatcherCtrl) response(data interface{}, context *gin.Context) {
	json_str, _ := json.Marshal(data)
	context.Writer.Write(json_str)
}
