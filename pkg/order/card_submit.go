package order

import (
	"encoding/json"
	"github.com/li-bao-jia/millennium"
	"net/url"
)

/**
 * 卡券提取接口
 */

type CardSubmit struct{}

type CardSubmitParams struct {
	OutOrderID  string `json:"outOrderId"`  // 接入方订单ID，长度<=50，幂等性保证
	UUID        string `json:"uuid"`        // 充值号码/帐号/卡号
	ItemID      string `json:"itemId"`      // 商品ID
	ItemFace    string `json:"itemFace"`    // 商品面值（单位：元），用于校验，为空则不校验
	ItemPrice   string `json:"itemPrice"`   // 结算单价（单位：元，精确到小数点后3位），用于校验，为空则不校验
	Amount      string `json:"amount"`      // 充值数量，默认值1
	SupplyMode  string `json:"supplyMode"`  // 供货模式(可选参数),0：异步供货，1：同步供货
	CallbackURL string `json:"callbackUrl"` // 订单状态回调地址，为空则不回调
	PhoneNo     string `json:"phoneNo"`     // 手机号，可选参数，少部分卡券商品需要用户号码
	Ext1        string `json:"ext1"`        // 扩展参数1（Q币/游戏：终端ip、中石化：手机号、电费：省份|市）
	Ext2        string `json:"ext2"`        // 扩展参数2（游戏：区、中石化/电费：身份证号后6位）
	Ext3        string `json:"ext3"`        // 扩展参数3（游戏：服、中石化：姓名，电费：1-住宅，2-店铺，3-企业）
}

type CardSubmitResponse struct {
	Code       string `json:"code"`       // 请求结果
	Msg        string `json:"msg"`        // 结果描述
	Cost       string `json:"cost"`       // 订单总消费: 单位元
	OrderID    string `json:"orderId"`    // 平台订单号
	OutOrderID string `json:"outOrderId"` // 接入方订单号
	CardData   []Card `json:"cardData"`   // 卡券信息json数组
}

type Card struct {
	CardName   string `json:"cardName"`   // 卡卷名称
	CardType   string `json:"cardType"`   // 卡券形式： LINK 链接 PICTURE 券码+链接 NUMBER_PASSWORD 卡号+密码 PASSWORD 密码
	CardNo     string `json:"cardNo"`     // 卡号(卡券形式为NUMBER_PASSWORD时有值)，
	CardPwd    string `json:"cardPwd"`    // 密码(卡券类型为NUMBER_PASSWORD、PASSWORD、PICTURE时有值)
	CardLink   string `json:"cardLink"`   // 链接(卡券形式为LINK或PICTURE时有值)
	ExpireTime string `json:"expireTime"` // 卡券有效期，格式：yyyy-MM-dd HH:mm:ss
}

func (s *CardSubmit) getMethod() string {
	return "api/card/get"
}

func (s *CardSubmit) Handle(c *millennium.ApiClient, p CardSubmitParams) (error, CardSubmitResponse) {
	// 请求参数
	params := url.Values{
		"appId":       {c.GetAppId()},
		"outOrderId":  {p.OutOrderID},
		"uuid":        {p.UUID},
		"itemId":      {p.ItemID},
		"itemFace":    {p.ItemFace},
		"itemPrice":   {p.ItemPrice},
		"amount":      {p.Amount},
		"supplyMode":  {p.SupplyMode},
		"callbackUrl": {p.CallbackURL},
		"phoneNo":     {p.PhoneNo},
		"ext1":        {p.Ext1},
		"ext2":        {p.Ext2},
		"ext3":        {p.Ext3},
		"timestamp":   {c.GetTimestamp()},
	}
	params.Set("sign", c.GenerateSign(params))

	// 发送请求
	res, err := c.SendPostRequest(c.GetFullUrl(s.getMethod()), params)
	if err != nil {
		return err, CardSubmitResponse{}
	}

	// 解析响应
	var response CardSubmitResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return err, CardSubmitResponse{}
	}

	// 返回结果
	return nil, response
}
