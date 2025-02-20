package millennium

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
	"github.com/li-bao-jia/millennium/pkg"
)

type ApiClient struct {
	appKey    string
	secret    string
	isHttp    bool   // 是否使用http协议 ture为http，false为https
	version   string // 请求版本
	domain    string // 正式域名
	devDomain string // 开发域名
}

/**
 * @Description:创建一个新的ApiClient
 * @param appKey
 * @param secret
 * @param domain 根据域名控制控制正式环境 或 测试环境
 */

func NewApiClient(appKey, secret, domain string) *ApiClient {
	return &ApiClient{
		appKey:  appKey,
		secret:  secret,
		isHttp:  false,
		version: "v1",
		domain:  domain,
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

	return https + "://" + a.domain + "/" + a.version + "/" // fmt.Sprintf("%s://%s/%s/", https, domain, Version)
}
