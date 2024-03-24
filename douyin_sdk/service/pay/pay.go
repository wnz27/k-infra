/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-21 10:49:47
 * @LastEditTime: 2024-03-24 09:04:26
 * @FilePath: /k-infra/douyin_sdk/service/pay/pay.go
 * @description: type some description
 */
package pay

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"strings"

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
	// 创建一个缓冲区来读取请求体
	var buf bytes.Buffer
	scanner := bufio.NewScanner(ginCtx.Request.Body)
	for scanner.Scan() {
		// 去除每行末尾的换行符和空格
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// 将处理后的行追加到缓冲区
		buf.WriteString(line)
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("scanner err:" + scanner.Err().Error())
	}
	// 将缓冲区内容转换为字符串
	bodyString := buf.String()
	// 重置请求体
	ginCtx.Request.Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))
	var douyinPlatformReq common.DouyinPayCallBackRequest
	bindErr := ginCtx.ShouldBind(&douyinPlatformReq)
	if bindErr != nil {
		return nil, bindErr
	}
	// reqCtx := ginCtx.Request.Context()
	// 从请求头中拿数据
	// get data from header
	return &common.DouyinPayCallBackReqAllData{
		BodyString: bodyString,
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
