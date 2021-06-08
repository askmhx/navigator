package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"rocket.iosxc.com/navigator/v1/model"
	"rocket.iosxc.com/navigator/v1/service"
)

type DispatcherCtrl struct {
	MerchantService service.MerchantService
	FileService     service.FileService
	AccountService  service.AccountService
	ChannelService  service.ChannelService
}

func NewDispatcherCtrl() DispatcherCtrl {
	return DispatcherCtrl{}
}

func (this *DispatcherCtrl) Handle(context *gin.Context) {
	var requestData model.ApiRequest
	err := context.ShouldBind(&requestData)
	if err != nil {
		fmt.Println("error happened:", err)
	}
	var vErr = this.verify(requestData)
	if vErr != nil {

		this.response(nil, context)
	} else {
		this.response(nil, context)
	}
}

func (this *DispatcherCtrl) Upload(context *gin.Context) {
	var requestData model.UploadRequest
	err := context.ShouldBind(&requestData)
	if err != nil {
		fmt.Println("error happened:", err)
	}
	var vErr = this.verify(requestData)
	result := this.FileService.Upload(requestData)
	if vErr != nil && result.Code == model.RESULT_CODE_SUCCESS {
		this.response(nil, context)
	} else {
		this.response(nil, context)
	}
}

func (this *DispatcherCtrl) Download(context *gin.Context) {
	var requestData model.DownloadRequest
	err := context.ShouldBind(&requestData)
	if err != nil {
		fmt.Println("error happened:", err)
	}
	var vErr = this.verify(requestData)
	result := this.FileService.Download(requestData)
	if vErr != nil {
		this.response(result.Data, context)
	} else {
		this.response(nil, context)
	}
}

func (this *DispatcherCtrl) verify(request model.RequestInterface) error {
	key := this.MerchantService.GetMerchantKey(request.GetRequest())
	if request.Verify(key) {
		return errors.New("verify sign failed")
	}
	return nil
}

func (this *DispatcherCtrl) response(data interface{}, context *gin.Context) {
	json_str, _ := json.Marshal(data)
	context.Writer.Write(json_str)
}
