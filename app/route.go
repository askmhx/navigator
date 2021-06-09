package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rocket.iosxc.com/navigator/v1/controller"
)

const URL_API_DOWNLOAD = "/api/download"

func initRoute(config *AppConfig, dbServer *gorm.DB, router *gin.Engine) {
	dispatcherCtrl := controller.NewDispatcherCtrl()
	router.POST(URL_API_DOWNLOAD, dispatcherCtrl.Download)
}
