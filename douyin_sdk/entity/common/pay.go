/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-19 23:38:05
 * @LastEditTime: 2024-03-20 09:03:37
 * @FilePath: /k-infra/douyin_sdk/entity/common/pay.go
 * @description: type some description
 */
package common

import "context"

type CallBackPayStatus string

const (
	CallBackPayStatusSuccess CallBackPayStatus = "SUCCESS" // 支付成功
	CallBackPayStatusCancel  CallBackPayStatus = "CANCEL"  // 支付取消
)

type PayCallBackRequest struct {
	// 订单相关信息的 json 字符串 eg. "{\"app_id\":\"tt07e371xxxxxxx\",\"status\":\"SUCCESS\",\"order_id\":\"ot7057422956397414686\",\"cp_extra\":\"xxx\",\"item_id\":\"xxxxx\",\"seller_uid\":\"xxxxxx\",\"pay_channel\":1,\"message\":\"\",\"extra\":\"{\\\"cps_info\\\":\\\"poi\\\",\\\"share_amount\\\":\\\"299\\\"}\",\"event_time\":1643185090000,\"out_order_no\":\"ext_order_no_1643185079529\",\"total_amount\":1}"
	Msg string `json:"msg"`
	// 回调类型（支付结果回调为 payment）：payment（支付成功/支付取消）
	Type string `json:"type"`
	// 固定值："3.0"。回调版本，用于开发者识别回调参数的变更
	Version string `json:"version"`
}

func (pcbReq *PayCallBackRequest) ToPayCallBackAllData() PayCallBackReqAllData {
	return PayCallBackReqAllData{}
}

type PayCallBackReqAllData struct {
	PayCallBackRequest
}

// 由 PayCallBackRequest 的 msg Unmarshal 而来
// From PayCallBackRequest.Msg Unmarshal
type PayCallBackInfo struct {
	AppID          string            `json:"app_id"`
	OutOrderNo     string            `json:"out_order_no"`
	OrderID        string            `json:"order_id"`
	Status         CallBackPayStatus `json:"status"`
	TotalAmount    int               `json:"total_amount"`
	DiscountAmount int               `json:"discount_amount"`
	MerchantUID    string            `json:"merchant_uid"`
	Message        string            `json:"message"`
	EventTime      int64             `json:"event_time"`
	PayChannel     int               `json:"pay_channel"`
	ChannelPayID   string            `json:"channel_pay_id"`
	UserBillPayID  string            `json:"user_bill_pay_id"`
}

type PayCallBackResponse struct {
	DouyinBaseErrCode
}

type ParsePayCallBackReq interface {
	ParsePayCallBackRequest(ctx context.Context) (*PayCallBackReqAllData, error)
}
