/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 01:03:07
 * @LastEditTime: 2024-03-17 01:16:34
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

// 性能更好 | better performance
// result of benchmark in test file: str_util_test.go
func GenRandomStringV2(length int) string {
	// 定义生成的随机字符串的长度
	// length := 10

	// 定义包含所有可能字符的字符串
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
