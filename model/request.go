package model

type ConfigRequest struct {
	AppId     string `form:"AppId" binding:"required"`
	Cluster   string `form:"Cluster" binding:"required"`
	Profile   string `form:"Profile" binding:"required"`
	NotifyUrl string `form:"NotifyUrl" binding:"required"`
	Sign      string `form:"Sign" binding:"required"`
}

type NotifyRequest struct {
	AppId      string
	Cluster    string
	Profile    string
	Config     string
	UpdateTime string
	UpdateBy   string
	Sign       string
}
