/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-21 10:49:47
 * @LastEditTime: 2024-03-21 15:48:06
 * @FilePath: /k-infra/douyin_sdk/service/pay/pay.go
 * @description: type some description
 */
package pay

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/wnz27/k-infra/douyin_sdk/entity/common"
)

// GinParser implement DouyinPayCallBackReqParser
// gin 框架的解析
type GinParser struct {
}

func NewGinParser() *GinParser {
	return &GinParser{}
}

/**
 * @description: 从请求中获取 post 的请求体以及从 header 中拿到抖音平台过来的回调请求携带的请求头
 * @param {context.Context} ctx
 * @return {*}
 */
func (parser *GinParser) ParseDouyinPayCallBackRequest(ctx context.Context) (*common.DouyinPayCallBackReqAllData, error) {
	ginCtx := ctx.(*gin.Context)
	var douyinPlatformReq common.DouyinPayCallBackRequest
	bindErr := ginCtx.ShouldBind(&douyinPlatformReq)
	if bindErr != nil {
		return nil, bindErr
	}
	// reqCtx := ginCtx.Request.Context()
	// 从请求头中拿数据
	// get data from header
	return &common.DouyinPayCallBackReqAllData{
		DouyinPayCallBackRequest: common.DouyinPayCallBackRequest{
			Msg:     douyinPlatformReq.Msg,
			Type:    douyinPlatformReq.Type,
			Version: douyinPlatformReq.Version,
		},
		DouyinPlatformReqHeader: common.DouyinPlatformReqHeader{
			ByteIdentifyName: ginCtx.Request.Header.Get(common.DouyinIDHeader),
			ByteLogID:        ginCtx.Request.Header.Get(common.DouyinLogIDHeader),
			ByteNonceStr:     ginCtx.Request.Header.Get(common.DouyinNonceHeader),
			ByteSignature:    ginCtx.Request.Header.Get(common.DouyinSignatureHeader),
			ByteTimestamp:    ginCtx.Request.Header.Get(common.DouyinTimeStampHeader),
		},
	}, nil
}
