package millennium

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type ApiClient struct {
	domain    string // 请求域名（https://www.aaa.com 或 http://www.aaa.com）
	appId     string
	appSecret string
}

/**
 * @Description:创建一个新的ApiClient
 *
 * @param domain
 * @param appId
 * @param appSecret
 */

func NewApiClient(domain, appId, appSecret string) *ApiClient {
	return &ApiClient{
		domain:    domain,
		appId:     appId,
		appSecret: appSecret,
	}
}

func (a *ApiClient) GetAppId() string {
	return a.appId
}

func (a *ApiClient) GetTimestamp() string {
	return time.Now().Format("20060102150405000")
}

func (a *ApiClient) GetFullUrl(method string) string {
	return a.domain + "/" + method
}

func (a *ApiClient) GenerateSign(p url.Values) string {
	var params []string

	// 遍历参数并添加到切片（排除空值）
	for k, v := range p {
		if len(v) > 0 && v[0] != "" {
			params = append(params, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}

	// 添加 appSecret
	params = append(params, fmt.Sprintf("appSecret=%s", a.appSecret))

	// 按参数名升序排序
	sort.Strings(params)

	// 连接参数
	signStr := strings.Join(params, "&")

	// 计算 MD5 并返回小写32位结果
	sum := md5.Sum([]byte(signStr))

	return fmt.Sprintf("%x", sum)
}

// 发送 HTTP POST 请求的封装

func (a *ApiClient) SendPostRequest(apiURL string, params url.Values) ([]byte, error) {
	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应数据
	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return body, nil
}

/**
 * @Description:DecryptAES256ECB 实现 AES-256 ECB 模式解密
 */

func (a *ApiClient) DecryptAES256ECB(pass string) (err error, str string) {
	cardNumberRes, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return
	}

	r, err := openssl.AesECBDecrypt(cardNumberRes, []byte(a.appSecret), openssl.PKCS7_PADDING)
	if err != nil {
		return
	}

	str = string(r)

	return
}
