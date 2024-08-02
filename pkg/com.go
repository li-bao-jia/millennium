package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type IOperate interface {
	GetMethod() string
}

type ApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
	Sign string `json:"sign"`
}

func Post(url string, data string) (res ApiResponse, err error) {
	client := &http.Client{Timeout: 15 * time.Second}

	var req *http.Request
	if req, err = http.NewRequest("POST", url, strings.NewReader(data)); err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	// 反序列化 res
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}

	return
}

func PostParams(appKey, appSecret string, data interface{}) (string, error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	params := generateParams(appKey, appSecret, dataByte)

	var paramByte []byte
	if paramByte, err = json.Marshal(params); err != nil {
		return "", err
	}

	return string(paramByte), nil
}

func generateParams(appKey, appSecret string, dataByte []byte) map[string]interface{} {
	params := map[string]interface{}{
		"app_key":   appKey,
		"timestamp": time.Now().Unix(),
		"version":   "1.0",
		"format":    "json",
		"charset":   "utf-8",
		"sign_type": "md5",
		"data":      string(dataByte),
	}
	params["sign"] = Sign(params, appSecret)

	return params
}
