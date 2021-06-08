package model

import "os"

type RequestInterface interface {
	Verify(merchantKey MerchantKey) bool
	GetRequest() RequestBaseStruct
}

type RequestBaseStruct struct {
	Version    string `form:"version" binding:"required"`
	MerchantId string `form:"merchant_id" binding:"required"`
	SignKey    int    `form:"sign_key" binding:"required"`
	SignType   string `form:"sign_type" binding:"required"`
	Sign       string `form:"sign" binding:"required"`
}

type ApiRequest struct {
	RequestBaseStruct
	ServiceCode string `form:"service_code" binding:"required"`
	RequestData string `form:"data"`
}

type UploadRequest struct {
	RequestBaseStruct
	FileName string  `form:"file_name" binding:"required"`
	File     os.File `form:"file" binding:"required"`
}

type DownloadRequest struct {
	RequestBaseStruct
	FileType string `form:"file_type" binding:"required"`
	FileDate string `form:"file_date" binding:"required"`
}

func (this RequestBaseStruct) Verify(merchantKey MerchantKey) bool {
	return true
}

func (this RequestBaseStruct) GetRequest() RequestBaseStruct {
	return this
}
