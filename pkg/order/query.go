package order

import (
	"encoding/json"
	"github.com/li-bao-jia/millennium"
	"net/url"
)

/**
 * 订单查询接口
 */

type Query struct{}

type QueryParams struct {
	OrderID    string `json:"orderId"`    // 平台订单号(orderId与outOrderId至少填写1项，如都填写以orderId进行查询)
	OutOrderNo string `json:"outOrderNo"` // 接入方订单号
}

type QueryResponse struct {
	Code         string `json:"code"`         // 请求结果
	Msg          string `json:"msg"`          // 结果描述
	OrderID      string `json:"orderId"`      // 平台订单号
	OutOrderID   string `json:"outOrderId"`   // 接入方订单号
	OrderStatus  string `json:"orderStatus"`  // 订单状态: 1-处理中, 2-充值成功, 3-充值失败, 4-未查询到订单
	OrderDesc    string `json:"orderDesc"`    // 结果描述: 充值成功/失败原因
	CompleteTime string `json:"completeTime"` // 订单完成时间: 格式 yyyyMMddHHmmss
	Cost         string `json:"cost"`         // 订单总消费: 单位元
	CardData     []Card `json:"cardData"`     // 卡券信息(仅卡密订单返回)
	Ext1         string `json:"ext1"`         // 扩展字段1: 透传上游数据, 流水号
	Ext2         string `json:"ext2"`         // 扩展字段2: 透传上游数据
	Ext3         string `json:"ext3"`         // 扩展字段3: 透传上游数据
}

func (q *Query) getMethod() string {
	return "api/order/query"
}

func (q *Query) Handle(c *millennium.ApiClient, p QueryParams) (error, QueryResponse) {
	// 请求参数
	params := url.Values{
		"appId":      {c.GetAppId()},
		"orderId":    {p.OrderID},
		"outOrderId": {p.OutOrderNo},
		"timestamp":  {c.GetTimestamp()},
	}
	params.Set("sign", c.GenerateSign(params))

	// 发送请求
	res, err := c.SendPostRequest(c.GetFullUrl(q.getMethod()), params)
	if err != nil {
		return err, QueryResponse{}
	}

	// 解析响应
	var response QueryResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return err, QueryResponse{}
	}

	// 返回结果
	return nil, response
}
