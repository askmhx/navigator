package model

const RESULT_CODE_SUCCESS = "0000"
const RESULT_CODE_FAIL = "9999"
const RESULT_CODE_FILE_NOT_FOUND = "4000"

type CommonResult struct {
	Code    string
	Message string
	Data    interface{}
}
