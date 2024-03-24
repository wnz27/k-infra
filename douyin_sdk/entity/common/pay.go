/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-19 23:38:05
 * @LastEditTime: 2024-03-24 09:19:39
 * @FilePath: /k-infra/douyin_sdk/entity/common/pay.go
 * @description: type some description
 */
/*
参考文档 / reference document
https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trade-system/general/order/notify-payment-result
https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/trade-system/general/common-param
*/
package common

import (
	"context"
	"errors"

	jsoniter "github.com/json-iterator/go"

	"github.com/wnz27/k-infra/douyin_sdk/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type DouyinCallBackPayStatus string

const (
	DouyinCallBackPayStatusSuccess DouyinCallBackPayStatus = "SUCCESS" // 支付成功
	DouyinCallBackPayStatusCancel  DouyinCallBackPayStatus = "CANCEL"  // 支付取消
)

type DouyinCallBackPayChannel int32

const (
	DouyinCallBackPayChannelUnPay     DouyinCallBackPayChannel = iota // 未支付
	DouyinCallBackPayChannelWeChat                                    // 1: 微信
	DouyinCallBackPayChannelAliPay                                    // 2: 支付宝
	DouyinCallBackPayChannelDouYinPay DouyinCallBackPayChannel = 10   // 10: 抖音支付
)

type DouyinPayCallBackRequest struct {
	// 订单相关信息的 json 字符串 eg. "{\"app_id\":\"tt07e371xxxxxxx\",\"status\":\"SUCCESS\",\"order_id\":\"ot7057422956397414686\",\"cp_extra\":\"xxx\",\"item_id\":\"xxxxx\",\"seller_uid\":\"xxxxxx\",\"pay_channel\":1,\"message\":\"\",\"extra\":\"{\\\"cps_info\\\":\\\"poi\\\",\\\"share_amount\\\":\\\"299\\\"}\",\"event_time\":1643185090000,\"out_order_no\":\"ext_order_no_1643185079529\",\"total_amount\":1}"
	Msg string `json:"msg"`
	// 回调类型（支付结果回调为 payment）：payment（支付成功/支付取消）
	Type string `json:"type"`
	// 固定值："3.0"。回调版本，用于开发者识别回调参数的变更
	Version string `json:"version"`
}

func (pcbReq *DouyinPayCallBackRequest) ToPayCallBackAllData(ctx context.Context, parser DouyinPayCallBackReqParser) (*DouyinPayCallBackReqAllData, error) {
	return parser.ParseDouyinPayCallBackRequest(ctx)
}

type DouyinPayCallBackReqAllData struct {
	DouyinPayCallBackRequest
	DouyinPlatformReqHeader
	BodyString string // 解析 body 变为串
}

func (reqData *DouyinPayCallBackReqAllData) VerifySign(
	publicKeyTokenStr string, isDebug bool) (bool, error) {
	return utils.VerifySign(
		reqData.ByteTimestamp, reqData.ByteNonceStr,
		reqData.BodyString, reqData.ByteSignature,
		publicKeyTokenStr, isDebug,
	)
}

func (reqData *DouyinPayCallBackReqAllData) ParseDouyinOrderInfo() (DouyinPayCallBackInfo, error) {
	infoBytes := []byte(reqData.Msg)
	var payCBInfo DouyinPayCallBackInfo
	err := json.Unmarshal(infoBytes, &payCBInfo)
	if err != nil {
		return payCBInfo, err
	}
	return payCBInfo, nil
}

type BaseDouyinOrderInfo struct {
	OrderID        string                   `json:"order_id"`         //  抖音开平侧订单id，长度 <= 64byte
	Status         DouyinCallBackPayStatus  `json:"status"`           //  支付结果状态，目前有两种状态： "SUCCESS" （支付成功 ） "CANCEL" （支付取消）
	TotalAmount    int64                    `json:"total_amount"`     //  订单总金额，单位分支付金额为 = total_amount - discount_amount
	DiscountAmount int64                    `json:"discount_amount"`  //  订单优惠金额，单位分，接入营销时请关注这个字段
	PayChannel     DouyinCallBackPayChannel `json:"pay_channel"`      //  支付渠道枚举（支付成功时才有）：1：微信 2：支付宝 10：抖音支付
	ChannelPayID   string                   `json:"channel_pay_id"`   //  渠道支付单号，如微信/支付宝的支付单号，长度 <= 64byte 注：status="SUCCESS"时一定有值
	UserBillPayID  string                   `json:"user_bill_pay_id"` //  对应用户抖音账单里的"支付单号" 注：status="SUCCESS"时一定有值
	MerchantUID    string                   `json:"merchant_uid"`     //  该笔交易的卖家商户号 注：status="SUCCESS"时一定有值
	Message        string                   `json:"message"`          //  该笔交易取消原因，如："USER_CANCEL"：用户取消 "TIME_OUT"：超时取消
	EventTime      int64                    `json:"event_time"`       //  用户支付成功/支付取消时间戳，单位为毫秒
}

func (order *BaseDouyinOrderInfo) IsPaid() bool {
	return order.Status == DouyinCallBackPayStatusSuccess
}

// 由 PayCallBackRequest 的 msg Unmarshal 而来
// From PayCallBackRequest.Msg Unmarshal
type DouyinPayCallBackInfo struct {
	// 跟开发者相关的信息
	// developer info
	AppID      string `json:"app_id"`
	OutOrderNo string `json:"out_order_no"` //  开发者系统生成的订单号，与抖音开平交易单号 order_id 唯一关联，长度 <= 64byte
	// 抖音开发平台内部的信息
	// douyin platform info
	BaseDouyinOrderInfo
}

// 尽量不要强依赖该方法的字符串
func (payCBInfo *DouyinPayCallBackInfo) ParamCheck() error {
	if payCBInfo.Status == DouyinCallBackPayStatusSuccess {
		// if payCBInfo.PayChannel != CallBackPayChannelDouYinPay &&
		// 	payCBInfo.PayChannel != CallBackPayChannelAliPay &&
		// 	payCBInfo.PayChannel != CallBackPayChannelWeChat {
		if payCBInfo.PayChannel == DouyinCallBackPayChannelUnPay {
			return errors.New("Douyin platform error: pay_channel is not valid")
		}
		if payCBInfo.ChannelPayID == "" {
			return errors.New("Douyin platform error: channel_pay_id is empty")
		}
		if payCBInfo.UserBillPayID == "" {
			return errors.New("Douyin platform error: user_bill_pay_id is empty")
		}
		if payCBInfo.MerchantUID == "" {
			return errors.New("Douyin platform error: merchant_uid is empty")
		}
	}
	return nil
}

type DouyinPayCallBackResponse struct {
	DouyinBaseErrCode
}

// 从请求中获取 post 的请求体以及从 header 中拿到抖音平台过来的回调请求携带的请求头
// get post req body and get douyin platform header from request
type DouyinPayCallBackReqParser interface {
	ParseDouyinPayCallBackRequest(ctx context.Context) (*DouyinPayCallBackReqAllData, error)
}
