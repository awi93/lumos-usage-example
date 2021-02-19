package model

type BaseResponse struct {
	Id   string      `json:"id"`
	Data interface{} `json:"data"`
}
