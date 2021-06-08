package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/controller"
)

const URL_API_REQUEST = "/api/request"
const URL_API_UPLOAD = "/api/upload"
const URL_API_DOWNLOAD = "/api/download"

func initRoute(config *AppConfig, dbServer *gorm.DB, router *gin.Engine) {
	dispatcherCtrl := controller.NewDispatcherCtrl()

	//merchantRepository := repository.NewMerchantRepository(db)
	//merchantService := service.NewMerchantService(merchantRepository, cacheRespository)
	//
	//merchantRepository := repository.NewMerchantRepository(db)
	//accountService := service.NewMerchantService(merchantRepository, cacheRespository)
	//
	//dispatcherCtrl.MerchantService = merchantService
	//dispatcherCtrl.FileService = service.NewAccountService()
	//dispatcherCtrl.AccountService = service.NewAccountService()
	//dispatcherCtrl.ChannelService = service.NewChannelService()

	router.POST(URL_API_REQUEST, dispatcherCtrl.Handle)
	router.POST(URL_API_UPLOAD, dispatcherCtrl.Upload)
	router.POST(URL_API_DOWNLOAD, dispatcherCtrl.Download)
}
