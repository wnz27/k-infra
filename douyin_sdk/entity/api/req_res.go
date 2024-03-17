/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 01:18:23
 * @LastEditTime: 2024-03-17 23:26:51
 * @FilePath: /k-infra/douyin_sdk/entity/api/req_res.go
 * @description: type some description
 */
package api

type DouyinBaseRequest struct {
	APPID  string `json:"appid"`  // 抖音小程序的AppID
	Secret string `json:"secret"` // 小程序的 APP Secret，可以在开发者后台获取
}

type DouyinBaseResponse struct {
	// ErrNo   int    `json:"err_no"`
	// ErrTips string `json:"err_tips"`
	DouyinErrCode
	// Data    interface{} `json:"data"`
}
