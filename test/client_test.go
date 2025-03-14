package test

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/balance"
	"github.com/li-bao-jia/millennium/pkg/order"
	"github.com/li-bao-jia/millennium/pkg/product"
	"testing"
)

var (
	domain    = ""
	appId     = ""
	appSecret = ""
)

/**
 * @Description:测试余额查询接口
 */
func TestBalanceQuery(t *testing.T) {
	client := millennium.NewApiClient(domain, appId, appSecret)

	var (
		err      error
		query    balance.Query
		response balance.QueryResponse
	)
	if err, response = query.Handle(client); err != nil {
		fmt.Println(err)
		return
	}

	if response.Code != "00" {
		fmt.Println(response.Msg)
		return
	}

	fmt.Printf("余额: %s", response.Balance)
}

/**
 * @Description:测试商品同步接口
 */
func TestProductQuery(t *testing.T) {
	client := millennium.NewApiClient(domain, appId, appSecret)

	var (
		err      error
		query    product.Query
		response product.QueryResponse
	)
	if err, response = query.Handle(client); err != nil {
		fmt.Println(err)
		return
	}

	if response.Code != "00" {
		fmt.Println(response.Msg)
		return
	}

	fmt.Printf("返回商品数量: %d", len(response.Data))
}

/**
 * @Description:测试通用直充接口
 */
func TestOrderSubmit(t *testing.T) {
	client := millennium.NewApiClient(domain, appId, appSecret)

	var params = order.SubmitParams{
		OutOrderID:  "20240801001",
		UUID:        "751818588",
		ItemID:      "10001",
		ItemFace:    "100.00",
		Amount:      "1",
		CallbackURL: "",
		SMSCode:     "",
		Ext1:        "",
		Ext2:        "",
		Ext3:        "",
	}

	var (
		err      error
		query    order.Submit
		response order.SubmitResponse
	)
	if err, response = query.Handle(client, params); err != nil {
		fmt.Println(err)
		return
	}

	if response.Code != "00" {
		fmt.Println(response.Msg)
		return
	}

	fmt.Printf("返回平台订单号: %s", response.OrderID)
}

/**
 * @Description:测试话费直充接口
 */
func TestHfSubmit(t *testing.T) {
	client := millennium.NewApiClient(domain, appId, appSecret)

	var params = order.HfSubmitParams{
		OutOrderID:   "20240801001",
		UUID:         "751818588",
		ItemID:       "10001",
		ItemFace:     "100.00",
		CallbackURL:  "",
		Isp:          "",
		ProvinceCode: "",
		ProvinceName: "",
		Timeout:      "",
		SMSCode:      "",
	}

	var (
		err      error
		query    order.HfSubmit
		response order.HfSubmitResponse
	)
	if err, response = query.Handle(client, params); err != nil {
		fmt.Println(err)
		return
	}

	if response.Code != "00" {
		fmt.Println(response.Msg)
		return
	}

	fmt.Printf("返回平台订单号: %s", response.OrderID)
}

/**
 * @Description:测试卡券提取接口
 */
func TestCardSubmit(t *testing.T) {
	client := millennium.NewApiClient(domain, appId, appSecret)

	var params = order.CardSubmitParams{
		OutOrderID:  "20240801001",
		UUID:        "751818588",
		ItemID:      "10001",
		ItemFace:    "100.00",
		Amount:      "1",
		SupplyMode:  "",
		CallbackURL: "",
		PhoneNo:     "",
		Ext1:        "",
		Ext2:        "",
		Ext3:        "",
	}

	var (
		err      error
		query    order.CardSubmit
		response order.CardSubmitResponse
	)
	if err, response = query.Handle(client, params); err != nil {
		fmt.Println(err)
		return
	}

	if response.Code != "00" {
		fmt.Println(response.Msg)
		return
	}

	fmt.Printf("返回平台订单号: %s", response.OrderID)
}

/**
 * @Description:测试查询订单提交API
 */
func TestOrderQuery(t *testing.T) {
	client := millennium.NewApiClient(domain, appId, appSecret)

	// 查询参数
	var params = order.QueryParams{
		OrderID:    "201701010101010001", // orderId与outOrderId至少填写1项，如都填写以orderId进行查询
		OutOrderNo: "cy1017010101010101",
	}

	var (
		err      error
		query    order.Query
		response order.QueryResponse
	)
	if err, response = query.Handle(client, params); err != nil {
		fmt.Println(err)
		return
	}

	if response.Code != "00" {
		fmt.Println(response.Msg)
		return
	}

	fmt.Printf("返回平台订单号: %s", response.OrderID)

	// 卡密解密
	if response.OrderStatus == "2" && len(response.CardData) > 0 {
		for i, card := range response.CardData {
			switch card.CardLink {
			case "LINK":
				_, link := client.DecryptAES256ECB(card.CardLink)

				fmt.Println("第%d张卡密CardLink：%s", i+1, link)
			case "PICTURE":
				_, link := client.DecryptAES256ECB(card.CardLink)
				_, pwd := client.DecryptAES256ECB(card.CardPwd)

				fmt.Println("第%d张卡密CardLink：%s", i+1, link)
				fmt.Println("第%d张卡密CardPwd：%s", i+1, pwd)
			case "PASSWORD":
				_, pwd := client.DecryptAES256ECB(card.CardPwd)

				fmt.Println("第%d张卡密CardPwd：%s", i+1, pwd)
			case "NUMBER_PASSWORD":
				_, no := client.DecryptAES256ECB(card.CardNo)
				_, pwd := client.DecryptAES256ECB(card.CardPwd)

				fmt.Println("第%d张卡密CardNo：%s", i+1, no)
				fmt.Println("第%d张卡密CardPwd：%s", i+1, pwd)
			}
		}
	}
}
