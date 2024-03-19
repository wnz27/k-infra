/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 21:27:10
 * @LastEditTime: 2024-03-19 23:43:00
 * @FilePath: /k-infra/douyin_sdk/entity/api/get_access_token.go
 * @description: type some description
 */
package api

import "github.com/wnz27/k-infra/douyin_sdk/entity/common"

/*
https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-access-token
接口说明：
为了保障应用的数据安全，只能在开发者服务器使用 AppSecret，如果小程序存在泄露 AppSecret 的问题，字节小程序平台将有可能下架该小程序，并暂停该小程序相关服务。​

access_token 是小程序的全局唯一调用凭据，开发者调用小程序支付时需要使用 access_token。 access_token 的有效期为 2 个小时，需要定时刷新 access_token，重复获取会导致之前一次获取的 access_token 的有效期缩短为 5 分钟。

​后续会新增和抖音能力 client_token 相通的接口，各 openapi 支持 client_token 调用，建议开发者逐步迁移新接口。

HTTP URL:
正式地址：https://developer.toutiao.com/api/apps/v2/token
沙盒地址：https://open-sandbox.Douyin.com/api/apps/v2/token
*/

const (
	GetAccessTokenURI               = "/api/apps/v2/token"
	GetAccessTokenProdURL           = common.ToutiaoProdDomain + GetAccessTokenURI
	GetAccessTokenSandBoxURL        = common.DouyinUpperSandboxDomain + GetAccessTokenURI
	GetAccessTokenMethod            = "POST"
	GetAccessTokenContentTypeHeader = "application/json"  // content-type
	GetAccessTokenGrantType         = "client_credential" // accept
)

func BuildGetAccessTokenURL(env common.DouyinDevEnv) string {
	if env == common.DouyinDevEnvTest {
		return GetAccessTokenSandBoxURL
	}
	return GetAccessTokenProdURL
}

type GetAccessTokenRequest struct {
	// APPID  string `json:"appid"`  // 抖音小程序的AppID
	// Secret string `json:"secret"` // 小程序的 APP Secret，可以在开发者后台获取
	common.DouyinBaseRequest
	GrantType string `json:"grant_type"` // 授权类型，此处只需填写 client_credential
}

type GetAccessTokenResData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 过期时间
}

type GetAccessTokenResponse struct {
	// ErrNo   int    `json:"err_no"`
	// ErrTips string `json:"err_tips"`
	common.DouyinBaseResponse
	Data GetAccessTokenResData `json:"data"`
}

type GetAccessTokenErrorResponse struct {
	// ErrNo   int    `json:"err_no"`
	// ErrTips string `json:"err_tips"`
	common.DouyinBaseResponse
	Data string `json:"data"`
}
