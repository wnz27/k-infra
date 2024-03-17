/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-16 23:20:01
 * @LastEditTime: 2024-03-17 23:38:27
 * @FilePath: /k-infra/douyin_sdk/entity/common/const.go
 * @description: type some description
 */
package common

/*
一些常量定义
some const value define
*/

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

const (
	DouyinSignResponseTimestampHeader         string = "Byte-Timestamp"
	DouyinSignResponseNonceHeader             string = "Byte-Nonce-Str"
	DouyinSignResponseSignatureHeader         string = "Byte-Signature"
	DouyinSignReqAuthHeaderKey                string = "Byte-Authorization"
	DouyinSignReqAuthHeaderAlgorithmType      string = "SHA256-RSA2048"
	DouyinSignReqAuthHeaderValueKeyAppID      string = "appid"
	DouyinSignReqAuthHeaderValueKeyNonce      string = "nonce_str"
	DouyinSignReqAuthHeaderValueKeyTimestamp  string = "timestamp"
	DouyinSignReqAuthHeaderValueKeyKeyVersion string = "key_version" // 公钥版本必须填写计算签名值时采用的应用私钥对应的应用公钥版本，应用公钥版本可通过「开发管理-开发设置-密钥设置」处获取。
	DefaultAppPublicKeyVersion                string = "1"
	DouyinSignReqAuthHeaderValueKeySignature  string = "signature"
)
