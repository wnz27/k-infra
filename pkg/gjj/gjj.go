/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2023-01-17 22:32:24
 * @LastEditTime: 2023-01-17 23:56:46
 * @FilePath: /k-go-infra/pkg/gjj/gjj.go
 * @description: type some description
 */

package gjj

import (
	"time"

	"github.com/wnz27/k-go-infra/pkg"
)

type GJJCalculatePeriod struct {
	StartMonthIdx int       // 该阶段 开始的序号数
	StartTime     time.Time // 年月 只到年月的时间
	EndTime       time.Time // 年月 只到年月的时间
	MoneyPerMonth float32   // 每月缴存金额
}

func currMonthSum(moneyPerMonth float32, currMonthIdx int) float32 {
	return moneyPerMonth * float32(currMonthIdx) * pkg.GJJ_Factor_CD
}

func (gjjP *GJJCalculatePeriod) Sum(discontinuousMonthsMap map[string]int) (float32, error) {
	sumAmount := float32(0)
	monthDateList, err1 := pkg.GetMonthListFromDateScope(gjjP.StartTime, gjjP.EndTime)
	if err1 != nil {
		return 0, err1
	}
	for i, dtStr := range monthDateList {
		currIdx := gjjP.StartMonthIdx + i
		_, ok := discontinuousMonthsMap[dtStr]
		if !ok {
			sumAmount += currMonthSum(gjjP.MoneyPerMonth, currIdx)
		}

	}
	return sumAmount, nil
}

func buildDiscontinuousMonthsMap(discontinuousMonths []time.Time) map[string]int {
	res := make(map[string]int)
	for _, t := range discontinuousMonths {
		res[pkg.ToDateMonthStr(t)] = 1
	}
	return res
}

func CalculateGJJLoanQuota(gjjPeriods []*GJJCalculatePeriod, discontinuousMonths []time.Time) (float32, error) {
	discontinuousMonthsMap := buildDiscontinuousMonthsMap(discontinuousMonths)
	allPeriodsSum := float32(0)
	for _, p := range gjjPeriods {
		currSum, err1 := p.Sum(discontinuousMonthsMap)
		if err1 != nil {
			return 0, err1
		}
		allPeriodsSum += currSum
	}
	return allPeriodsSum, nil
}
