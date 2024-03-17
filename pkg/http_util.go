/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-03-17 23:53:29
 * @LastEditTime: 2024-03-18 00:00:31
 * @FilePath: /k-infra/pkg/http_util.go
 * @description: type some description
 */
package pkg

import (
	"errors"
	"reflect"
	"time"

	"github.com/valyala/fasthttp"
)

func BuildFastHttpClient(readTimeoutStr, writeTimeoutStr, maxIdleConnDurationStr string) *fasthttp.Client {
	// You may read the timeouts from some config
	// readTimeout, _ := time.ParseDuration("500ms")
	readTimeout, _ := time.ParseDuration(readTimeoutStr)
	// writeTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration(writeTimeoutStr)
	// maxIdleConnDuration, _ := time.ParseDuration("1h")
	maxIdleConnDuration, _ := time.ParseDuration(maxIdleConnDurationStr)
	client := &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
	return client
}

func HTTPConnError(err error) (string, bool) {
	var (
		errName string
		known   = true
	)

	switch {
	case errors.Is(err, fasthttp.ErrTimeout):
		errName = "timeout"
	case errors.Is(err, fasthttp.ErrNoFreeConns):
		errName = "conn_limit"
	case errors.Is(err, fasthttp.ErrConnectionClosed):
		errName = "conn_close"
	case reflect.TypeOf(err).String() == "*net.OpError":
		errName = "timeout"
	default:
		known = false
	}

	return errName, known
}
