package millennium

import (
	"github.com/li-bao-jia/millennium/pkg"
)

var (
	Http        = false                     // 是否使用http协议 ture为http，false为https
	Development = false                     // 是否为开发模式 ture为开发模式，false为正式模式
	Version     = "v1"                      // 请求的版本
	Domain      = "openapi.qianxiquan.com"  // 正式域名
	DevDomain   = "testopen.qianxiquan.com" // 开发域名
)

type ApiClient struct {
	AppKey string
	Secret string
}

/**
 * @Description:创建一个新的ApiClient
 */

func NewApiClient(appKey, secret string) *ApiClient {
	return &ApiClient{AppKey: appKey, Secret: secret}
}

/**
 * @Description:请求api并返回结果
 */

func (a *ApiClient) CallApi(o pkg.IOperate, data interface{}) (res pkg.ApiResponse, err error) {
	var paramStr string
	if paramStr, err = pkg.PostParams(a.AppKey, a.Secret, data); err != nil {
		return
	}

	res, err = pkg.Post(a.GetUrl()+o.GetMethod(), paramStr)
	return
}

/**
 * @Description:定义请求的http协议 ture为http，false为https
 */

func (a *ApiClient) SetHttp(h bool) {
	Http = h
}

/**
 * @Description:定义请求的版本
 */

func (a *ApiClient) SetVersion(v string) {
	Version = v
}

/**
 * @Description:定义请求的模式，true为开发模式，false为正式模式
 */

func (a *ApiClient) SetDevelopment(d bool) {
	Development = d
}

/**
 * @Description:获取请求的域名
 */

func (a *ApiClient) GetUrl() string {
	https := "https"
	if Http {
		https = "http"
	}

	domain := Domain
	if Development {
		domain = DevDomain
	}

	return https + "://" + domain + "/" + Version + "/" // fmt.Sprintf("%s://%s/%s/", https, domain, Version)
}
