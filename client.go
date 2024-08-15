package millennium

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
	"github.com/li-bao-jia/millennium/pkg"
)

type ApiClient struct {
	appKey    string
	secret    string
	isDev     bool   // 是否为开发模式 ture为开发模式，false为正式模式
	isHttp    bool   // 是否使用http协议 ture为http，false为https
	version   string // 请求版本
	domain    string // 正式域名
	devDomain string // 开发域名
}

/**
 * @Description:创建一个新的ApiClient
 */

func NewApiClient(appKey, secret string) *ApiClient {
	return &ApiClient{
		appKey:    appKey,
		secret:    secret,
		isDev:     false,
		isHttp:    false,
		version:   "v1",
		domain:    "openapi.qianxiquan.com",
		devDomain: "testopen.qianxiquan.com",
	}
}

/**
 * @Description:请求api并返回结果
 */

func (a *ApiClient) CallApi(o pkg.IOperate, data interface{}) (res pkg.ApiResponse, err error) {
	var paramStr string
	if paramStr, err = pkg.PostParams(a.appKey, a.secret, data); err != nil {
		return
	}

	res, err = pkg.Post(a.getUrl()+o.GetMethod(), paramStr)
	return
}

/**
 * @Description:定义请求的模式，true为开发模式，false为正式模式
 */

func (a *ApiClient) SetDev(d bool) {
	a.isDev = d
}

/**
 * @Description:定义请求的http协议 ture为http，false为https
 */

func (a *ApiClient) SetHttp(h bool) {
	a.isHttp = h
}

/**
 * @Description:定义请求的版本
 */

func (a *ApiClient) SetVersion(v string) {
	a.version = v
}

/**
 * @Description:DecryptAES256ECB 实现 AES-256 ECB 模式解密
 */

func (a *ApiClient) DecryptAES256ECB(pass string) (str string, err error) {
	cardNumberRes, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return
	}

	r, err := openssl.AesECBDecrypt(cardNumberRes, []byte(a.secret), openssl.PKCS7_PADDING)
	if err != nil {
		return
	}

	str = string(r)

	return
}

/**
 * @Description:获取请求的域名
 */

func (a *ApiClient) getUrl() string {
	https := "https"
	if a.isHttp {
		https = "http"
	}

	domain := a.domain
	if a.isDev {
		domain = a.devDomain
	}

	return https + "://" + domain + "/" + a.version + "/" // fmt.Sprintf("%s://%s/%s/", https, domain, Version)
}
