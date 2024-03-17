/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 01:03:07
 * @LastEditTime: 2024-03-17 21:25:03
 * @FilePath: /k-infra/pkg/str_util.go
 * @description: type some description
 */
package pkg

import (
	"math/rand"
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
