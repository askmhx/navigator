package model

type ConfigRequest struct {
	AppId   string `form:"AppId" binding:"required"`
	Cluster string `form:"Cluster" binding:"required"`
	Profile string `form:"Profile" binding:"required"`
	Sign    string `form:"Sign" binding:"required"`
}

type ConfigResponse struct {
	AppId    string
	Cluster  string
	Profile  string
	Config   string
	CreateAt string
	CreateBy string
	Sign     string
}
