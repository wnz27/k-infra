/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 23:08:40
 * @LastEditTime: 2024-03-18 00:36:05
 * @FilePath: /k-infra/douyin_sdk/service/token/token.go
 * @description: type some description
 */
package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/wnz27/k-infra/douyin_sdk/entity/api"
	"github.com/wnz27/k-infra/douyin_sdk/entity/common"
	"github.com/wnz27/k-infra/pkg"
)

type TokenService struct {
	client *fasthttp.Client
	// 支持一个加一个
	// support one and add a one service call
	getAccessTokenURL string
}

func NewTokenService(env common.DouyinDevEnv, fastHttpClient *fasthttp.Client) *TokenService {
	tokenSrv := &TokenService{
		getAccessTokenURL: api.BuildGetAccessTokenURL(env),
		client:            fastHttpClient,
	}
	return tokenSrv
}

func (srv *TokenService) GetAccessToken(getAccessTokenReq *api.GetAccessTokenRequest) (*api.GetAccessTokenResponse, error) {
	// 使用 fasthttp 发送 POST 请求
	// per-request timeout
	reqTimeout := time.Duration(500) * time.Millisecond

	reqBytes, _ := json.Marshal(getAccessTokenReq)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(srv.getAccessTokenURL)
	req.Header.SetMethod(fasthttp.MethodPost)

	var headerContentTypeJson = []byte(api.GetAccessTokenContentTypeHeader)
	req.Header.SetContentTypeBytes(headerContentTypeJson)
	// set body
	req.SetBodyRaw(reqBytes)

	resp := fasthttp.AcquireResponse()

	err := srv.client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		errName, known := pkg.HTTPConnError(err)
		if known {
			// fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
			return nil, errors.New(errName)
		} else {
			// fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
			return nil, err
		}
	}

	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("ERR invalid HTTP response code: %d", statusCode))
	}

	respBody := resp.Body()
	// 跟文档不太一致，错误的话返回的是这种结构体
	// fmt.Println("=======>", string(respBody))
	respError := api.GetAccessTokenErrorResponse{}
	err1 := json.Unmarshal(respBody, &respError)
	if err1 == nil {
		return nil, errors.New(respError.Data)
	}

	respEntity := api.GetAccessTokenResponse{}
	err = json.Unmarshal(respBody, &respEntity)
	if err == nil {
		return &respEntity, nil
	} else {
		// if errors.Is(err, io.EOF) {
		// 	fmt.Printf("DEBUG Parsed Response: %v\n", respEntity)
		// }
		return nil, errors.New(fmt.Sprintf("ERR failed to parse response: %v\n", err))
	}
}
