/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 23:46:05
 * @LastEditTime: 2024-03-18 00:29:55
 * @FilePath: /k-infra/douyin_sdk/service/token/token_test.go
 * @description: type some description
 */
package token

import (
	"testing"

	"github.com/wnz27/k-infra/douyin_sdk/entity/api"
	"github.com/wnz27/k-infra/douyin_sdk/entity/common"
	"github.com/wnz27/k-infra/pkg"
)

func TestGetAccessToken(t *testing.T) {
	c := pkg.BuildFastHttpClient("500ms", "500ms", "1h")
	tokenService := NewTokenService(common.DouyinDevEnvTest, c)
	apiReq := api.GetAccessTokenRequest{
		DouyinBaseRequest: api.DouyinBaseRequest{
			APPID:  "xxxx",
			Secret: "xxxxx",
		},
		GrantType: api.GetAccessTokenGrantType,
	}
	_, err := tokenService.GetAccessToken(&apiReq)
	t.Log("error is --------> ", err)

	/*
			错误 === RUN   TestGetAccessToken
		------> {"data":"app is not sandboxApp","error":2}
		正确：
		------> {"err_no":0,"err_tips":"success","data":{"access_token":"xxxxxxxxxxxxxxxxxxxxxxx","expires_in":7200,"expiresAt":1710699496}}
	*/
}
