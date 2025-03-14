package product

import (
	"encoding/json"
	"github.com/li-bao-jia/millennium"
	"net/url"
)

/**
 * 商品同步接口
 */

type Query struct{}

type QueryResponse struct {
	Code string    `json:"code"` // 请求结果
	Msg  string    `json:"msg"`  // 结果描述
	Data []Product `json:"data"`
}

type Product struct {
	ItemID   string `json:"itemId"`   // 商品编码
	ItemName string `json:"itemName"` // 商品名称
	ItemType string `json:"itemType"` // 商品类型: 0-直充, 1-卡券
	ItemFace string `json:"itemFace"` // 商品面额: 月卡、周卡、100元等
	ItemVal  string `json:"itemVal"`  // 商品面值: 标准价或官方原价
	Status   string `json:"status"`   // 商品状态: 1-正常, 0-维护
	Price    string `json:"price"`    // 商品价格: 单位元
}

func (q *Query) getMethod() string {
	return "api/item/query"
}

func (q *Query) Handle(c *millennium.ApiClient) (error, QueryResponse) {
	// 请求参数
	params := url.Values{
		"appId":     {c.GetAppId()},
		"timestamp": {c.GetTimestamp()},
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
