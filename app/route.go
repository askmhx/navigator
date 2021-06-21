package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/controller"
	"rocket.iosxc.com/navigator/v1/repository"
	"rocket.iosxc.com/navigator/v1/service"
)

const URL_API_DOWNLOAD = "/config"

func initRoute(config *AppConfig, dbServer *gorm.DB, router *gin.Engine) {
	appRepository := repository.NewAppRepository(dbServer)
	configRepositry := repository.NewConfigRepository(dbServer)
	configService := service.NewConfigService(appRepository, configRepositry)
	dispatcherCtrl := controller.NewDispatcherCtrl(configService)
	router.POST(URL_API_DOWNLOAD, dispatcherCtrl.Download)
	initTaskManager(configService)
}
