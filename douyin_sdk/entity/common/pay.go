/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-19 23:38:05
 * @LastEditTime: 2024-03-19 23:41:03
 * @FilePath: /k-infra/douyin_sdk/entity/common/pay.go
 * @description: type some description
 */
package common

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

/*
//支付取消回调示例

	{
	    "app_id": "tt07e371xxxxxxx",
	    "out_order_no": "motb52726742593307630520652",
	    "order_id": "ext_order_123",
	    "status": "CANCEL",
	    "total_amount": 1,
	    "discount_amount": 0,
	    "merchant_uid": "1231123",
	    "message": "",
	    "event_time": 1692775192000,
		"pay_channel": 1,
		"channel_pay_id": "2iu2082897r9hflquf",
		"user_bill_pay_id": "DPTS12031230128124421",
	}
*/
// 由 msg unmarshall来
type PayCallBackInfo struct {
	Status CallBackPayStatus `json:"status"`
}

type PayCallBackResponse struct {
	DouyinBaseErrCode
}
