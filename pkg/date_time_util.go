/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-20 13:09:20
 * @LastEditTime: 2024-03-20 13:11:10
 * @FilePath: /k-infra/pkg/date_time_util.go
 * @description: type some description
 */
package pkg

import (
	"errors"
	"time"
)

func DefaultZeroDateTime() (time.Time, error) {
	var defaultZeroTime time.Time
	// 加载本地时区
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return defaultZeroTime, errors.New("无法加载本地时区")
	}

	// 创建默认时间
	defaultZeroTime = time.Date(0, 0, 0, 0, 0, 0, 0, loc)
	return defaultZeroTime, nil
}
