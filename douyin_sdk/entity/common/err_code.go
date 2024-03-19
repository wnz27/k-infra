/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 23:22:51
 * @LastEditTime: 2024-03-19 23:39:29
 * @FilePath: /k-infra/douyin_sdk/entity/common/err_code.go
 * @description: type some description
 */
package common

// V1 版本
var ErrorCodeMapString = map[int]string{
	0:     "请求成功",
	-1:    "系统错误",
	40015: "appid 错误",
	40017: "secret 错误",
	40020: "grant_type 不是 client_credential",
}

type DouyinBaseErrCode struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
}

func (ec DouyinBaseErrCode) ErrMsg() string {
	return ErrorCodeMapString[ec.ErrNo]
}

type DouyinErrCode struct {
	DouyinBaseErrCode
	Error int `json:"error"`
}
