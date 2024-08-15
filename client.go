package millennium

import (
	"github.com/li-bao-jia/millennium/pkg"
)

type ApiClient struct {
	AppKey    string
	Secret    string
	Version   string // 请求的版本
	IsDev     bool   // 是否为开发模式 ture为开发模式，false为正式模式
	IsHttp    bool   // 是否使用http协议 ture为http，false为https
	Domain    string // 正式域名
	DevDomain string // 开发域名
}

/**
 * @Description:创建一个新的ApiClient
 */

func NewApiClient(appKey, secret string) *ApiClient {
	return &ApiClient{
		AppKey:    appKey,
		Secret:    secret,
		IsHttp:    false,
		IsDev:     false,
		Version:   "v1",
		Domain:    "openapi.qianxiquan.com",
		DevDomain: "testopen.qianxiquan.com",
	}
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
	a.IsHttp = h
}

/**
 * @Description:定义请求的版本
 */

func (a *ApiClient) SetVersion(v string) {
	a.Version = v
}

/**
 * @Description:定义请求的模式，true为开发模式，false为正式模式
 */

func (a *ApiClient) SetDevelopment(d bool) {
	a.IsDev = d
}

/**
 * @Description:获取请求的域名
 */

func (a *ApiClient) GetUrl() string {
	https := "https"
	if a.IsHttp {
		https = "http"
	}

	domain := a.Domain
	if a.IsDev {
		domain = a.DevDomain
	}

	return https + "://" + domain + "/" + a.Version + "/" // fmt.Sprintf("%s://%s/%s/", https, domain, Version)
}
