package order

import (
	"encoding/json"
	"github.com/li-bao-jia/millennium"
	"net/url"
)

/**
 * 通用直充接口
 */

type Submit struct{}

type SubmitParams struct {
	OutOrderID  string `json:"outOrderId"`            // 接入方订单ID，长度<=50，幂等性保证
	UUID        string `json:"uuid"`                  // 充值号码/帐号/卡号
	ItemID      string `json:"itemId"`                // 商品ID
	ItemFace    string `json:"itemFace,omitempty"`    // 商品面值（单位：元），用于校验，为空则不校验
	ItemPrice   string `json:"itemPrice,omitempty"`   // 结算单价（单位：元，精确到小数点后3位），用于校验，为空则不校验
	Amount      string `json:"amount"`                // 充值数量，默认值1
	CallbackURL string `json:"callbackUrl,omitempty"` // 订单状态回调地址，为空则不回调
	SMSCode     string `json:"smsCode,omitempty"`     // 短信验证码，部分商品需要
	Ext1        string `json:"ext1,omitempty"`        // 扩展参数1（Q币/游戏：终端ip、中石化：手机号、电费：省份|市）
	Ext2        string `json:"ext2,omitempty"`        // 扩展参数2（游戏：区、中石化/电费：身份证号后6位）
	Ext3        string `json:"ext3,omitempty"`        // 扩展参数3（游戏：服、中石化：姓名，电费：1-住宅，2-店铺，3-企业）
}

type SubmitResponse struct {
	Code       string `json:"code"`       // 请求结果
	Msg        string `json:"msg"`        // 结果描述
	Cost       string `json:"cost"`       // 订单总消费: 单位元
	OrderID    string `json:"orderId"`    // 平台订单号
	OutOrderID string `json:"outOrderId"` // 接入方订单号
}

func (s *Submit) getMethod() string {
	return "api/order/submit"
}

func (s *Submit) Handle(c *millennium.ApiClient, p SubmitParams) (error, SubmitResponse) {
	// 请求参数
	params := url.Values{
		"appId":       {c.GetAppId()},
		"outOrderId":  {p.OutOrderID},
		"uuid":        {p.UUID},
		"itemId":      {p.ItemID},
		"itemFace":    {p.ItemFace},
		"itemPrice":   {p.ItemPrice},
		"amount":      {p.Amount},
		"callbackUrl": {p.CallbackURL},
		"smsCode":     {p.SMSCode},
		"ext1":        {p.Ext1},
		"ext2":        {p.Ext2},
		"ext3":        {p.Ext3},
		"timestamp":   {c.GetTimestamp()},
	}
	params.Set("sign", c.GenerateSign(params))

	// 发送请求
	res, err := c.SendPostRequest(c.GetFullUrl(s.getMethod()), params)
	if err != nil {
		return err, SubmitResponse{}
	}

	// 解析响应
	var response SubmitResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return err, SubmitResponse{}
	}

	// 返回结果
	return nil, response
}
