/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-19 23:38:05
 * @LastEditTime: 2024-03-20 15:25:44
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
)

type CallBackPayStatus string

const (
	CallBackPayStatusSuccess CallBackPayStatus = "SUCCESS" // 支付成功
	CallBackPayStatusCancel  CallBackPayStatus = "CANCEL"  // 支付取消
)

type CallBackPayChannel int32

const (
	CallBackPayChannelUnPay     CallBackPayChannel = iota // 未支付
	CallBackPayChannelWeChat                              // 1: 微信
	CallBackPayChannelAliPay                              // 2: 支付宝
	CallBackPayChannelDouYinPay CallBackPayChannel = 10   // 10: 抖音支付
)

type PayCallBackRequest struct {
	// 订单相关信息的 json 字符串 eg. "{\"app_id\":\"tt07e371xxxxxxx\",\"status\":\"SUCCESS\",\"order_id\":\"ot7057422956397414686\",\"cp_extra\":\"xxx\",\"item_id\":\"xxxxx\",\"seller_uid\":\"xxxxxx\",\"pay_channel\":1,\"message\":\"\",\"extra\":\"{\\\"cps_info\\\":\\\"poi\\\",\\\"share_amount\\\":\\\"299\\\"}\",\"event_time\":1643185090000,\"out_order_no\":\"ext_order_no_1643185079529\",\"total_amount\":1}"
	Msg string `json:"msg"`
	// 回调类型（支付结果回调为 payment）：payment（支付成功/支付取消）
	Type string `json:"type"`
	// 固定值："3.0"。回调版本，用于开发者识别回调参数的变更
	Version string `json:"version"`
}

func (pcbReq *PayCallBackRequest) ToPayCallBackAllData(ctx context.Context, parser PayCallBackReqParser) (*PayCallBackReqAllData, error) {
	return parser.ParsePayCallBackRequest(ctx)
}

type PayCallBackReqAllData struct {
	PayCallBackRequest
	DouyinPlatformReqHeader
}

type BaseDouyinOrderInfo struct {
	OrderID        string             `json:"order_id"`         //  抖音开平侧订单id，长度 <= 64byte
	Status         CallBackPayStatus  `json:"status"`           //  支付结果状态，目前有两种状态： "SUCCESS" （支付成功 ） "CANCEL" （支付取消）
	TotalAmount    int64              `json:"total_amount"`     //  订单总金额，单位分支付金额为 = total_amount - discount_amount
	DiscountAmount int64              `json:"discount_amount"`  //  订单优惠金额，单位分，接入营销时请关注这个字段
	PayChannel     CallBackPayChannel `json:"pay_channel"`      //  支付渠道枚举（支付成功时才有）：1：微信 2：支付宝 10：抖音支付
	ChannelPayID   string             `json:"channel_pay_id"`   //  渠道支付单号，如微信/支付宝的支付单号，长度 <= 64byte 注：status="SUCCESS"时一定有值
	UserBillPayID  string             `json:"user_bill_pay_id"` //  对应用户抖音账单里的"支付单号" 注：status="SUCCESS"时一定有值
	MerchantUID    string             `json:"merchant_uid"`     //  该笔交易的卖家商户号 注：status="SUCCESS"时一定有值
	Message        string             `json:"message"`          //  该笔交易取消原因，如："USER_CANCEL"：用户取消 "TIME_OUT"：超时取消
	EventTime      int64              `json:"event_time"`       //  用户支付成功/支付取消时间戳，单位为毫秒
}

// 由 PayCallBackRequest 的 msg Unmarshal 而来
// From PayCallBackRequest.Msg Unmarshal
type PayCallBackInfo struct {
	// 跟开发者相关的信息
	// developer info
	AppID      string `json:"app_id"`
	OutOrderNo string `json:"out_order_no"` //  开发者系统生成的订单号，与抖音开平交易单号 order_id 唯一关联，长度 <= 64byte
	// 抖音开发平台内部的信息
	// douyin platform info
	BaseDouyinOrderInfo
}

// 尽量不要强依赖该方法的字符串
func (payCBInfo *PayCallBackInfo) ParamCheck() error {
	if payCBInfo.Status == CallBackPayStatusSuccess {
		// if payCBInfo.PayChannel != CallBackPayChannelDouYinPay &&
		// 	payCBInfo.PayChannel != CallBackPayChannelAliPay &&
		// 	payCBInfo.PayChannel != CallBackPayChannelWeChat {
		if payCBInfo.PayChannel == CallBackPayChannelUnPay {
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

type PayCallBackResponse struct {
	DouyinBaseErrCode
}

type PayCallBackReqParser interface {
	ParsePayCallBackRequest(ctx context.Context) (*PayCallBackReqAllData, error)
}
