/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-21 10:49:47
 * @LastEditTime: 2024-03-21 11:41:44
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
