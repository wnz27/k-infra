/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 23:22:51
 * @LastEditTime: 2024-03-18 00:19:20
 * @FilePath: /k-infra/douyin_sdk/entity/api/err_code.go
 * @description: type some description
 */
package api

// V1 版本
var ErrorCodeMapString = map[int]string{
	0:     "请求成功",
	-1:    "系统错误",
	40015: "appid 错误",
	40017: "secret 错误",
	40020: "grant_type 不是 client_credential",
}

type DouyinErrCode struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Error   int    `json:"error"`
}

func (ec DouyinErrCode) ErrMsg() string {
	return ErrorCodeMapString[ec.ErrNo]
}
