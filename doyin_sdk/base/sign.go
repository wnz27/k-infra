/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-16 22:36:34
 * @LastEditTime: 2024-03-16 22:53:05
 * @FilePath: /k-infra/doyin_sdk/base/sign.go
 * @description: type some description
 */
// 生成及验签逻辑
// gen or verify sign logic

// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/signature-algorithm
// doc: https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/signature-algorithm

// 该文件中方法的时间: 2024-03-16 22:41:44
// this file method time: 2024-03-16 22:41:44
package base

import "crypto/rsa"

func BuildSign(httpMethod, uri, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (string, error) {
	// TODO
	return "", nil
}

func VerifySign(timestamp, nonce, body, signature, pubKeyStr string) (bool, error) {
	// TODO
	return false, nil
}
