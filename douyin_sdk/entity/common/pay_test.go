/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-23 14:57:24
 * @LastEditTime: 2024-03-23 14:59:04
 * @FilePath: /k-infra/douyin_sdk/entity/common/pay_test.go
 * @description: type some description
 */
package common

import (
	"fmt"
	"testing"
)

func TestToBodyString(t *testing.T) {
	pcbReq := &DouyinPayCallBackReqAllData{
		DouyinPayCallBackRequest: DouyinPayCallBackRequest{
			Msg:     "msg",
			Type:    "type",
			Version: "version",
		},
		DouyinPlatformReqHeader: DouyinPlatformReqHeader{
			ByteIdentifyName: "xxx",
		},
	}
	a, err := pcbReq.ToBodyString()
	if err != nil {
		t.Errorf("ToBodyString() error = %v", err)
	}
	fmt.Println(a)
}
