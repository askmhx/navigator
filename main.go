package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"rocket.iosxc.com/navigator/v1/app"
	"runtime"
	"time"
)

var (
	AppBanner  = "Navigator By askmhx@gmail.com %s Date: %s Build: %s"
	AppVersion = "1.0.0"
	AppDate    = time.Now().Format("2006-01-02 15:04:05")
	GoVersion  = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

func main() {
	fmt.Println(fmt.Sprintf(AppBanner, AppVersion, AppDate, GoVersion))
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "NOT SET", "app cfg path")
	flag.Parse()
	if cfgPath == "NOT SET" {
		cfgPath = "./config.json"
	}
	config := app.InitConfig(cfgPath)
	gin := gin.Default()
	app.Launch(config, gin)
}
