/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 01:03:07
 * @LastEditTime: 2024-03-20 13:31:48
 * @FilePath: /k-infra/pkg/str_util.go
 * @description: type some description
 */
package pkg

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// 定义包含所有可能字符的字符串
var chars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

/**
 * @description: 生成指定长度的随机字符串 | generate random string with specified length
 * @param {int} length
 * @return {*}
 */
func GenRandomStringV2(length int) string {

	rand.Seed(time.Now().UnixNano())
	// 生成随机字符串
	randomString := make([]byte, length)
	for i := range randomString {
		randomByte := make([]byte, 1)
		rand.Read(randomByte)
		randomString[i] = chars[int(randomByte[0])%len(chars)]
	}

	// 输出随机字符串
	return string(randomString)
}

/**
 * @description: 生成指定长度的随机字符串 | generate random string with specified length
 * @param {int} length
 * @return {*}
 */
func GenRandomStringV1() string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 定义生成的随机字符串的长度
	length := 10

	// 生成随机字符串
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = chars[rand.Intn(len(chars))]
	}

	// 输出随机字符串
	return string(randomString)
}

func panicErr(errMsgKey string, err error) {
	errMsg := errMsgKey + ":" + err.Error()
	panic(errors.New(errMsg))
}

// F64ToString float64 --> string, precision 精度 -1 表示保留所有位, 产生截断会四舍五入
func F64ToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}

func ParseInt64FromStr(str string, errMsgKey string) int64 {
	res, e1 := strconv.ParseInt(str, 10, 64)
	if e1 != nil {
		panicErr(errMsgKey, e1)
	}
	return res
}
