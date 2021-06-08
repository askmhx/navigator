package model

type BaseResponse struct {
	Version    string `form:"version" binding:"required"`
	MerchantId string `form:"merchant_id" binding:"required"`
	SignKey    int    `form:"sign_key" binding:"required"`
	SignType   string `form:"sign_type" binding:"required"`
	Sign       string `form:"sign" binding:"required"`
}

type ApiResponse struct {
	RequestBaseStruct
	ServiceCode string `form:"service_code" binding:"required"`
	RequestData string `form:"data"`
}

type UploadResponse struct {
	RequestBaseStruct
	FileName string `form:"file_name" binding:"required"`
	File     string `form:"file" binding:"required"`
}

type DownloadResponse struct {
	RequestBaseStruct
	FileType string `form:"file_type" binding:"required"`
	FileDate string `form:"file_date" binding:"required"`
}
