/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-01-17 23:32:00
 * @LastEditTime: 2023-01-18 00:07:59
 * @FilePath: /k-go-infra/pkg/gjj/gjj_test.go
 * @description: type some description
 */

package gjj

import (
	"fmt"
	"testing"
	"time"

	"github.com/wnz27/k-go-infra/pkg"
)

func Test_buildDiscontinuousMonthsMap(t *testing.T) {
	a1, _ := pkg.ToMonthDate("2021-01")
	a2, _ := pkg.ToMonthDate("2021-05")
	a := []time.Time{a1, a2}
	res := buildDiscontinuousMonthsMap(a)
	fmt.Println(res)
}

func TestCalculateGJJLoanQuota(t *testing.T) {
	w1, _ := pkg.ToMonthDate("2021-11")
	w2, _ := pkg.ToMonthDate("2023-01")

	t1, _ := pkg.ToMonthDate("2020-05")
	t2, _ := pkg.ToMonthDate("2021-10")
	periods := []*GJJCalculatePeriod{
		&GJJCalculatePeriod{
			StartMonthIdx: 1,
			StartTime:     w1,
			EndTime:       w2,
			MoneyPerMonth: 1344,
		},
		&GJJCalculatePeriod{
			StartMonthIdx: 16,
			StartTime:     t1,
			EndTime:       t2,
			MoneyPerMonth: 2100,
		},
	}
	res1, err1 := CalculateGJJLoanQuota(periods, []time.Time{})
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("llll ----> ", res1)
}
