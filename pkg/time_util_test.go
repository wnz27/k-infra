/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-01-17 23:05:47
 * @LastEditTime: 2023-01-17 23:20:50
 * @FilePath: /k-go-infra/pkg/time_util_test.go
 * @description: type some description
 */

package pkg

import (
	"fmt"
	"testing"
	"time"
)

func TestMonthTimeFormat(t *testing.T) {
	a := "2021-02"
	a1, e1 := time.Parse(GeneralMonthFormat, a)
	if e1 != nil {
		panic(e1)
	}
	fmt.Println(a1)

}

func TestGetMonthListFromDateScope(t *testing.T) {
	startT, e1 := ToMonthDate("2021-01")
	if e1 != nil {
		panic(e1)
	}
	endT, e2 := ToMonthDate("2021-12")
	if e2 != nil {
		panic(e2)
	}
	res, e3 := GetMonthListFromDateScope(startT, endT)
	if e3 != nil {
		panic(e3)
	}
	fmt.Println(res)
}
