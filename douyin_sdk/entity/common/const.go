/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-16 23:20:01
 * @LastEditTime: 2024-03-21 11:50:10
 * @FilePath: /k-infra/douyin_sdk/entity/common/const.go
 * @description: type some description
 */
package common

/*
一些常量定义
some const value define
*/

const (
	GeneralDateTimeLayout = "2006-01-02 15:04:05"
	GeneralTimeFormat     = "15:04:05"
	GeneralDateFormat     = "2006-01-02"
)

// 基础常量定义
type DouyinDevEnv string

const (
	DouyinDevEnvTest DouyinDevEnv = "DY_TEST"
	DouyinDevEnvProd DouyinDevEnv = "DY_PROD"
)

// 沙盒环境的域名为
const (
	DouyinSandboxDomain      = "https://open-sandbox.douyin.com"
	DouyinUpperSandboxDomain = "https://open-sandbox.Douyin.com" // 暂不知区别，如果后续发现一样会替换掉
	DouyinProdDomain         = "https://open.douyin.com"
	ToutiaoProdDomain        = "https://developer.toutiao.com"
)

// Header
const (
	DouyinAuthHeader string = "Byte-Authorization"
	// 开发者回调路径，不包含域名
	DouyinIDHeader string = "Byte-Identifyname"
	// 请求包签名
	DouyinSignatureHeader string = "Byte-Signature"
	// 抖音开平统一日志id，当出现问题时可以提供此id给抖音研发人员协助定位问题
	DouyinLogIDHeader string = "Byte-Logid"
	// 随机字符串，由字母、数字、下划线组成，区分大小写，len(Byte-Nonce-Str) <= 128
	DouyinNonceHeader string = "Byte-Nonce-Str"
	// 请求时间戳，精度：秒
	DouyinTimeStampHeader string = "Byte-Timestamp"
)

// HeaderContent
const (
	DouyinSignReqAuthHeaderAlgorithmType      string = "SHA256-RSA2048"
	DouyinSignReqAuthHeaderValueKeyAppID      string = "appid"
	DouyinSignReqAuthHeaderValueKeyNonce      string = "nonce_str"
	DouyinSignReqAuthHeaderValueKeyTimestamp  string = "timestamp"
	DouyinSignReqAuthHeaderValueKeyKeyVersion string = "key_version" // 公钥版本必须填写计算签名值时采用的应用私钥对应的应用公钥版本，应用公钥版本可通过「开发管理-开发设置-密钥设置」处获取。
	DefaultAppPublicKeyVersion                string = "1"
	DouyinSignReqAuthHeaderValueKeySignature  string = "signature"
)
