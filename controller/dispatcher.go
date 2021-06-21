package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/service"
)

type DispatcherCtrl struct {
	ConfigService service.ConfigService
}

func NewDispatcherCtrl(service service.ConfigService) DispatcherCtrl {
	return DispatcherCtrl{ConfigService: service}
}


func (this *DispatcherCtrl) Download(context *gin.Context) {
	var requestData model.ConfigRequest
	err := context.ShouldBind(&requestData)
	if err != nil {
		fmt.Println("error happened:", err)
	}
	result := this.ConfigService.Download(requestData)
	this.response(result.Data, context)
}

func (this *DispatcherCtrl) response(data interface{}, context *gin.Context) {
	json_str, _ := json.Marshal(data)
	context.Writer.Write(json_str)
}
