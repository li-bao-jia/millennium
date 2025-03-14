package order

import (
	"encoding/json"
	"github.com/li-bao-jia/millennium"
	"net/url"
)

/**
 * 话费直充接口
 */

type HfSubmit struct{}

type HfSubmitParams struct {
	OutOrderID   string `json:"outOrderId"`   // 接入方订单ID，长度<=50，幂等性保证
	UUID         string `json:"uuid"`         // 充值号码/帐号/卡号
	ItemID       string `json:"itemId"`       // 商品ID
	ItemFace     string `json:"itemFace"`     // 商品面值（单位：元），用于校验，为空则不校验
	ItemPrice    string `json:"itemPrice"`    // 结算单价（单位：元，精确到小数点后3位），用于校验，为空则不校验
	CallbackURL  string `json:"callbackUrl"`  // 订单状态回调地址，为空则不回调
	Isp          string `json:"isp"`          // 运营商(移动：yd，联通：lt，电信：dx，广电：gd)如不传则使用本系统号码库识别，不保障携号转网数据准确性，携号问题自理
	ProvinceCode string `json:"provinceCode"` // 归属地省份代码，取值见省份代码表（与provinceName传其中一个即可，都传则以provinceCode为准）
	ProvinceName string `json:"provinceName"` // 归属地省份名称，取值见省份代码表（与provinceCode传其中一个即可，都传则以provinceCode为准）
	Timeout      string `json:"timeout"`      // 订单超时时间，单位：秒（超过该时间订单不再失败重试，如果已提交上级则需等待上级返回）
	SMSCode      string `json:"smsCode"`      // 短信验证码，部分商品需要

}

type HfSubmitResponse struct {
	Code       string `json:"code"`       // 请求结果
	Msg        string `json:"msg"`        // 结果描述
	Cost       string `json:"cost"`       // 订单总消费: 单位元
	OrderID    string `json:"orderId"`    // 平台订单号
	OutOrderID string `json:"outOrderId"` // 接入方订单号
}

func (s *HfSubmit) getMethod() string {
	return "api/hf/order/submit"
}

func (s *HfSubmit) Handle(c *millennium.ApiClient, p HfSubmitParams) (error, HfSubmitResponse) {
	// 请求参数
	params := url.Values{
		"appId":        {c.GetAppId()},
		"outOrderId":   {p.OutOrderID},
		"uuid":         {p.UUID},
		"itemId":       {p.ItemID},
		"itemFace":     {p.ItemFace},
		"itemPrice":    {p.ItemPrice},
		"callbackUrl":  {p.CallbackURL},
		"isp":          {p.Isp},
		"provinceCode": {p.ProvinceCode},
		"provinceName": {p.ProvinceName},
		"timeout":      {p.Timeout},
		"smsCode":      {p.SMSCode},
		"timestamp":    {c.GetTimestamp()},
	}
	params.Set("sign", c.GenerateSign(params))

	// 发送请求
	res, err := c.SendPostRequest(c.GetFullUrl(s.getMethod()), params)
	if err != nil {
		return err, HfSubmitResponse{}
	}

	// 解析响应
	var response HfSubmitResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return err, HfSubmitResponse{}
	}

	// 返回结果
	return nil, response
}
