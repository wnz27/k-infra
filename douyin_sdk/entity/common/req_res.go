/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 01:18:23
 * @LastEditTime: 2024-03-20 09:30:55
 * @FilePath: /k-infra/douyin_sdk/entity/common/req_res.go
 * @description: type some description
 */
package common

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

type DouyinPlatformReqHeader struct {
	ByteIdentifyName string // 取自抖音平台过来的请求的请求头 Byte-Identifyname
	ByteLogID        string // 取自抖音平台过来的请求的请求头 Byte-Logid
	ByteNonceStr     string // 取自抖音平台过来的请求的请求头 Byte-Nonce-Str
	ByteSignature    string // 取自抖音平台过来的请求的请求头 Byte-Signature
	ByteTimestamp    string // 取自抖音平台过来的请求的请求头 Byte-Timestamp
}

// base response use const
const (
	SuccessCode    = 0
	SuccessMessage = "success"
	// FailedMsg      = "failed"
)
