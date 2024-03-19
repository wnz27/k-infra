/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 23:22:51
 * @LastEditTime: 2024-03-19 23:55:14
 * @FilePath: /k-infra/douyin_sdk/entity/common/error.go
 * @description: type some description
 */
/*
 error code for douyin sdk
 抖音的错误码定义
*/
package common

// V1 版本
var ErrorCodeMapString = map[int]string{
	0:     "请求成功",
	-1:    "系统错误",
	40015: "appid 错误",
	40017: "secret 错误",
	40020: "grant_type 不是 client_credential",
	10000: "参数错误",  // 对照错误提示和接口字段定义，检查对应的参数
	11001: "访问未授权", // 1. 订单不属于该小程序，无法查询到该订单信息，检查订单是否属于该小程序 2. 没有该接口的请求权限，需要接入通用交易解决方案开通接口权限
	13000: "系统错误",  // 请重试，若多次重试仍然报错，请联系oncall
	20000: "订单不存在", // 查不到该订单，请检查订单号是否传错，是否把order_id 当 out_order_no 传入
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
