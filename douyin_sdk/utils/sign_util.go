// 生成及验签逻辑
// gen or verify sign logic

// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/signature-algorithm
// doc: https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/signature-algorithm

// 该文件中方法的时间: 2024-03-16 22:41:44
// this file method time: 2024-03-16 22:41:44
package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

/**
 * @description: gen sign string | 构造签名串
 * @param {*} httpMethod: http method support POST GET PUT only Upper | http 请求方法 支持 POST GET PUT 只能大写
 * @param {*} uri: request uri eg. https://xxx/com/api1/biz/xxx -> uri is /api1/biz/xxx | 请求 uri 去除域名部分
 * @param {*} timestamp: request timestamp | 请求时间戳
 * @param {*} nonce: request nonce | 请求随机数
 * @param {string} body: request body | 请求体 GET 无
 * @param {*rsa.PrivateKey} privateKey: rsa private key | rsa 私钥
 * @return {*} sign string | 签名串
 */
func BuildSign(httpMethod, uri, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (string, error) {
	//method内容必须大写，如GET、POST，uri不包含域名，必须以'/'开头
	targetStr := httpMethod + "\n" +
		uri + "\n" +
		timestamp + "\n" +
		nonce + "\n" +
		body + "\n"
	h := sha256.New()
	h.Write([]byte(targetStr))
	digestBytes := h.Sum(nil)

	signBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digestBytes)
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)

	return sign, nil
}

/**
 * @description: verify sign | 验签
 * @param {*} timestamp: request timestamp | 请求时间戳
 * @param {*} nonce: request nonce | 请求随机数
 * @param {*} body: request body | 请求体 GET 无
 * @param {*} signature: request signature | 请求签名串
 * @param {string} pubKeyStr: public key string | 公钥字符串
 * @return {*}
 */
func VerifySign(timestamp, nonce, body, signature, platformPubKeyStr string, isDebug bool) (bool, error) {
	pubKey, err1 := PemToRSAPublicKey(platformPubKeyStr) // 注意验签时publicKey使用平台公钥而非应用公钥
	if err1 != nil {
		return false, err1
	}

	hashed := sha256.Sum256([]byte(
		timestamp + "\n" +
			nonce + "\n" +
			body + "\n"),
	)
	if isDebug {
		fmt.Println("time: ", timestamp, "nonce: ", nonce, "body: ", body, "signature: ", signature)
	}
	signBytes, err2 := base64.StdEncoding.DecodeString(signature)
	if err2 != nil {
		return false, err2
	}
	err3 := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	if err3 != nil {
		return false, err3
	}
	return err3 == nil, nil
}

//   ------------------------    private func   ------------------------
/**
 * @description: use private key string to build rsa private key; secret format is psc8| 使用私钥字符串构造 rsa 私钥; 密码格式 psc8
 * @param {string} privateKeyString
 * @return {*}
 */
func buildPrivateKey(privateKeyString string) (*rsa.PrivateKey, error) {
	// 解析 PEM 格式的私钥
	block, _ := pem.Decode([]byte(privateKeyString))
	if block == nil {
		return nil, errors.New("Failed to decode PEM block containing private key")
	}

	// 解析私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Failed to parse private key:" + err.Error())
	}

	// 打印私钥信息
	// fmt.Println("RSA Private Key:", privateKey)
	return privateKey.(*rsa.PrivateKey), nil
}

/**
 * @description: public key string to rsa.PublicKey | 使用公钥字符串构造 rsa.PublicKey 对象
 * @param {string} publicKeyStr
 * @return {*}
 */
func PemToRSAPublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	block, e1 := pem.Decode([]byte(publicKeyStr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, fmt.Errorf("Empty block in pem string: %v", e1)
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, fmt.Errorf("Not rsa public key")
	}
}
