/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-16 23:51:32
 * @LastEditTime: 2024-03-23 11:46:25
 * @FilePath: /k-infra/douyin_sdk/utils/sign_util_test.go
 * @description: type some description
 */
package utils

import (
	"testing"
)

var mockPublicKey string = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5BsOF7lcjl/lwFZZwWxf
magH0kwR56yzbBbhethXS1csmInwVSmABjLl8fSTi80QL9cNkzFUDshN7l/gqtDT
xxmFQkopcxG+d7qllfMYqlxhssgwVQJSd1DUIuf67QazFR7a5ZJH0bY2DBFnSy+C
KZNFUlkgHJwGS1+9xlopgsTDKVSbiIkD84mgIAaHnn7bJmd1Zk8WURdp5nrfhjQd
zAsRO0t9SvukCqnYwrE9H3ahoqo7AlUzESm67WiwRC0AvSZkZWFQ6W5rDzf7gZib
KP5eVcLoh8W2wFwI+K3DJBJi/iVRzbRKP4uGepReoozCYJgkzl+jzHW7i3D+wl8t
awIDAQAB
-----END PUBLIC KEY-----
`

var mockPrivateKey string = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDkGw4XuVyOX+XA
VlnBbF+ZqAfSTBHnrLNsFuF62FdLVyyYifBVKYAGMuXx9JOLzRAv1w2TMVQOyE3u
X+Cq0NPHGYVCSilzEb53uqWV8xiqXGGyyDBVAlJ3UNQi5/rtBrMVHtrlkkfRtjYM
EWdLL4Ipk0VSWSAcnAZLX73GWimCxMMpVJuIiQPziaAgBoeeftsmZ3VmTxZRF2nm
et+GNB3MCxE7S31K+6QKqdjCsT0fdqGiqjsCVTMRKbrtaLBELQC9JmRlYVDpbmsP
N/uBmJso/l5VwuiHxbbAXAj4rcMkEmL+JVHNtEo/i4Z6lF6ijMJgmCTOX6PMdbuL
cP7CXy1rAgMBAAECggEBANLAaIHk4i4tTjIp1h3OynlGdYuIexaJSvD4JvYAQo86
iNqav7F/eWjAyBGH/naxTV5WHJW9gsNxsAcpEIPiE3kmDChWKvvFDizDB1CG3Wgx
mJa9PWKdlaHlzUo++WjbwsQl0OtA/xg0eLUnsz8VMvbEucy+yduFEI+9crJ8BVRq
RrhtqDG0A90BDDIhnhj3KPrfbATHb9KagFZrpO+xmh0t4KzspV+6AzfrkoI9vzky
2Mll7sTMkLjnbcIdUlJCEH9MH0FHBEnpEkTFV1kd1tQzuFU3N0503q/NuUOqLpds
C+f0YxHKhNag73uNN1zIUC9eEyBCxlx+tLn6EmmSBlECgYEA8jnmXjHrE1E7i8cB
aq4CSKci+l/iIIw0LHX17tfT7hWqLFDHimkNkFuPlp7z1W69EeWQgKtWRxgvIfnx
DMYmX90Nmn3tRO1526FkrUsz6uiyRGNnOKPQyPw9lCsyMp/tQsBPxmGUrBbvEcIJ
HB6TY8wH0jqNZbDPQ89LGD8OzBkCgYEA8ROaOTKg8Y47SS/At+IRjrAOGj+8fvDz
/JoIVNzWYXZcTap4gcgwczMdEoFcLxJBsZxxDaEiQGrgVNEsy6msKGseSWoDGhSD
aOjphbawyjOT3d+joUnz4ka08AEpc0cyr/4LEr37AIXg1z6ScXuMpGeLa6A/T1Bu
KzmdQ3AsNiMCgYAPhuiedyKzfUyM3DfaB8d7ssMKO6U6IuKhSvp10f3y0A61goQX
+j31V/kvVYcZ0lxqTkXiCZmhOwqiaewqvnTtRjU+Bv5zoaljC8hxV1W/pCTxP1H5
jn6us4Sa/93a4ueJlNxIQi8OjPXMNJzy4X7fMc/6iOhRcXEHzrzok/o12QKBgQCM
BmjD5EZbR9PjtJrps6OjD1uBn5eq2+W7yPQh5ouW3JrMecG5EEAkCYJPZ1fV93K0
6Ts5QWiVpf5bBYxRV2Ipr95NogffNB8H5pENG4ogSEkQzH9MhZnkylD6PpKG5Mnq
M1LXNgX+zcRFAZEp3StZqtLuVouvU/ZJoRNZQmRLpQKBgQCJnM1Pzk58meTEXsXQ
qpvI6zKjIPBjuZAmFPVw2YCpQxgmal1qi00pkoRRGc2kOyYoQ7gPvZ9/gVUUX0sb
Onqd5ld76/6J+regYJBee+7ZZUk8I258ZCPyqT4gBg7j7EMtsBEbuYN/CqNpbZbH
CSayosY3+XXx0MsJDpKG2FLpYg==
-----END PRIVATE KEY-----
`

func TestBuildPrivateKey(t *testing.T) {
	pk, err := buildPrivateKey(mockPrivateKey)
	if err != nil {
		t.Errorf("buildPrivateKey() error = %v", err)
	}
	if pk == nil {
		t.Error("build private key is failed")
	}
}

func TestPemToRSAPublicKey(t *testing.T) {
	pubK, err := PemToRSAPublicKey(mockPublicKey)
	if err != nil {
		t.Errorf("pemToRSAPublicKey() error = %v", err)
	}
	if pubK == nil {
		t.Error("pem to rsa public key is failed")
	}
}

func TestVerifySign(t *testing.T) {

}
