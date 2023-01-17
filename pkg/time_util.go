/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-01-17 22:25:26
 * @LastEditTime: 2023-01-17 23:21:58
 * @FilePath: /k-go-infra/pkg/time_util.go
 * @description: type some description
 */

package pkg

import (
	"errors"
	"time"
)

// ToDateMonthStr 时间格式 转为 年月的字符串 time to
func ToDateMonthStr(dt time.Time) string {
	return dt.Format(GeneralMonthFormat)
}

// ToMonthDate 把年月字符串变为 时间格式; monthDateStr must match 2021-02
func ToMonthDate(monthDateStr string) (time.Time, error) {
	return time.Parse(GeneralMonthFormat, monthDateStr)
}

// GetMonthListFromDateScope startTime endTime 只包含年月的时间 only contain year and month's time
// eg: 2021-02-01 00:00:00 +0000 UTC
func GetMonthListFromDateScope(startTime, endTime time.Time) ([]string, error) {
	if startTime.After(endTime) {
		return nil, errors.New("起始时间不能晚于结束时间 (start time don't later than end time!)")
	}

	res := []string{
		startTime.Format(GeneralMonthFormat),
	}
	for startTime.Before(endTime) {
		startTime = startTime.AddDate(0, 1, 0)
		res = append(res, startTime.Format(GeneralMonthFormat))
	}
	return res, nil
}
