/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 01:11:44
 * @LastEditTime: 2024-03-17 01:15:16
 * @FilePath: /k-infra/pkg/str_util_test.go
 * @description: type some description
 */
package pkg

import "testing"

var dLength int = 10

func TestGenRandomStringV1(t *testing.T) {
	t.Log(GenRandomStringV1())
	t.Log(GenRandomStringV1())
	t.Log(GenRandomStringV1())
	t.Log(GenRandomStringV1())
	t.Log(GenRandomStringV1())
	t.Log(GenRandomStringV1())
}

func TestGenRandomStringV2(t *testing.T) {

	t.Log(GenRandomStringV2(20))
	t.Log(GenRandomStringV2(20))
	t.Log(GenRandomStringV2(20))
	t.Log(GenRandomStringV2(20))
}

/*
BenchmarkGenRandomStringV1-10             100480             10852 ns/op              32 B/op          2 allocs/op
PASS
ok      github.com/wnz27/k-infra/pkg    1.618s
*/
func BenchmarkGenRandomStringV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomStringV1()
	}
}

/*
BenchmarkGenRandomStringV2-10            7632860               150.8 ns/op            32 B/op          2 allocs/op
PASS
ok      github.com/wnz27/k-infra/pkg    1.617s
*/
func BenchmarkGenRandomStringV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomStringV2(dLength)
	}
}
