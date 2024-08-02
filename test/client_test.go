package test

import (
	"encoding/json"
	"fmt"
	"github.com/li-bao-jia/millennium-go-sdk"
	"github.com/li-bao-jia/millennium-go-sdk/pkg/balance"
	"github.com/li-bao-jia/millennium-go-sdk/pkg/order"
	"github.com/li-bao-jia/millennium-go-sdk/pkg/product"
	"testing"
)

/**
 * @Description:测试商品列表获取API
 */
func TestListProduct(t *testing.T) {

	appKey := "8DVI1nCM6SEgi0T3HhUI1J2EQJA4sCKAiLCHU5xAuI5YKXZoEd0ysRQdaHU2DNJc"
	appSecret := "0a091b3aa4324435aab703142518a8f7"

	client := millennium_go_sdk.NewApiClient(appKey, appSecret)

	// 设置是否使用http，ture为http，false为https
	// client.SetHttp(true)

	// 设置开发者模式，true为开发模式，false为正式模式
	// client.SetDevelopment(true)

	// 调用接口
	resp, err := client.CallApi(&product.ListProduct{}, product.ListProductParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Code != 0 {
		fmt.Println(resp.Msg)
		return
	}

	// 转换商品
	var products []product.Product
	if err = json.Unmarshal([]byte(resp.Data), &products); err != nil {
		fmt.Println(resp.Msg)
		return
	}

	fmt.Printf("返回商品数量: %d", len(products))
	fmt.Printf("返回商品数据: %s", resp.Data)
}

/**
 * @Description:测试直充订单提交API
 */
func TestChargeOrder(t *testing.T) {

	appKey := "8DVI1nCM6SEgi0T3HhUI1J2EQJA4sCKAiLCHU5xAuI5YKXZoEd0ysRQdaHU2DNJc"
	appSecret := "0a091b3aa4324435aab703142518a8f7"

	client := millennium_go_sdk.NewApiClient(appKey, appSecret)

	// 设置开发者模式
	client.SetDevelopment(true)

	// 设置请求参数
	data := order.ChargeOrderParams{
		ProductId:  6,             // 千禧券商品ID
		OutOrderNo: "20240801001", // 单号
		Account:    "13188889999", // 充值账号
		BuyNum:     1,             // 购买数量（⚠️：目前只支持单笔充值）
	}

	resp, err := client.CallApi(&order.ChargeOrder{}, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Code != 0 {
		fmt.Println(resp.Msg)
		return
	}

	// 转换订单数据
	var cOrder order.Order
	if err = json.Unmarshal([]byte(resp.Data), &cOrder); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("返回直充订单数据: %s", resp.Data)
	//{"order_no":"20240801160211515353556570","out_order_no":"20240801001","product_id":6,"product_name":"京东E卡100元直充","account":"13188889999","buy_num":1,"order_type":2,"order_price":"100.00","order_state":"processing","create_time":"2024-08-01 16:02:11","finish_time":null,"operator_serial_number":""}
}

/**
 * @Description:测试卡密订单提交API
 */
func TestCardOrder(t *testing.T) {

	appKey := "8DVI1nCM6SEgi0T3HhUI1J2EQJA4sCKAiLCHU5xAuI5YKXZoEd0ysRQdaHU2DNJc"
	appSecret := "0a091b3aa4324435aab703142518a8f7"

	client := millennium_go_sdk.NewApiClient(appKey, appSecret)

	// 设置开发者模式
	client.SetDevelopment(true)

	// 设置请求参数
	data := order.CardOrderParams{
		ProductId:  8,             // 千禧券商品ID
		OutOrderNo: "20240801002", // 单号
		BuyNum:     2,             // 购买数量（⚠️：目前单笔数量最大50）
	}

	resp, err := client.CallApi(&order.CardOrder{}, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Code != 0 {
		fmt.Println(resp.Msg)
		return
	}

	// 转换订单数据
	var cOrder order.Order
	if err = json.Unmarshal([]byte(resp.Data), &cOrder); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("返回卡密订单数据: %s", resp.Data)
	//{"order_no":"20240801160550696865566968","out_order_no":"20240801002","product_id":8,"product_name":"卡密测试商品","buy_num":2,"order_type":1,"order_price":"10.00","order_state":"processing","create_time":"2024-08-01 16:05:50","finish_time":null,"operator_serial_number":""}
}

/**
 * @Description:测试查询订单提交API
 */
func TestQueryOrder(t *testing.T) {

	appKey := "8DVI1nCM6SEgi0T3HhUI1J2EQJA4sCKAiLCHU5xAuI5YKXZoEd0ysRQdaHU2DNJc"
	appSecret := "0a091b3aa4324435aab703142518a8f7"

	client := millennium_go_sdk.NewApiClient(appKey, appSecret)

	// 设置开发者模式
	client.SetDevelopment(true)

	// 设置请求参数
	data := order.QueryOrderParams{
		OutOrderNo: "20240801002",
	}

	resp, err := client.CallApi(&order.QueryOrder{}, data)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Code != 0 {
		t.Fatal(resp.Msg)
	}

	// 转换订单数据
	var cOrder order.Order
	if err = json.Unmarshal([]byte(resp.Data), &cOrder); err != nil {
		t.Fatal(err)
	}

	// 转换卡密数据 千禧券数据1卡密2直充
	if cOrder.OrderType == 1 {
		var cards []order.Card
		if err = json.Unmarshal([]byte(cOrder.Cards), &cards); err != nil {
			return
		}
		t.Logf("返回卡密数量: %d", len(cards))
		t.Logf("返回卡密数据: %s", cOrder.Cards)
		//[{"card_number":"aX0Z5kocbVsW3K0YaS2jNuf8tAaiQfNL7JlorlnOem4=","card_pwd":"vghJjofPK8PjVXeYrFta0y7ugv2ufinxFZtD+BqrvOX5oMllmH2sVlT47MlSPh09","card_deadline":"2045-03-31 23:59:59"},{"card_number":"Oyms4YwVN2tEjZQMaSVIL\/6+S4SXsEAYIRq5jNbWHr4=","card_pwd":"n8jrqZbU8\/oNoRDUh9RHWNLhkiF1QJHBdy\/MR7tKnyL5oMllmH2sVlT47MlSPh09","card_deadline":"2045-03-31 23:59:59"}]

		// 卡密数据使用时，需要解密显示
		for i, card := range cards {
			// 卡密解密
			var CardPwd, CardNumber string
			if CardPwd, err = order.DecryptAES256ECB(card.CardPwd, appSecret); err != nil {
				return
			}
			if CardNumber, err = order.DecryptAES256ECB(card.CardNumber, appSecret); err != nil {
				return
			}
			cards[i].CardPwd = CardPwd
			cards[i].CardNumber = CardNumber
		}
	}

	t.Logf("返回订单数据: %s", resp.Data)
	//{"order_no":"20240801160550696865566968","out_order_no":"20240801002","product_id":8,"product_name":"卡密测试商品","account":"","buy_num":2,"order_type":1,"order_price":"10.00","order_state":"success","order_service_state":"null","create_time":"2024-08-01 16:05:50","finish_time":"2024-08-01 16:08:03","cards":"[{\"card_number\":\"aX0Z5kocbVsW3K0YaS2jNuf8tAaiQfNL7JlorlnOem4=\",\"card_pwd\":\"vghJjofPK8PjVXeYrFta0y7ugv2ufinxFZtD+BqrvOX5oMllmH2sVlT47MlSPh09\",\"card_deadline\":\"2045-03-31 23:59:59\"},{\"card_number\":\"Oyms4YwVN2tEjZQMaSVIL\\\/6+S4SXsEAYIRq5jNbWHr4=\",\"card_pwd\":\"n8jrqZbU8\\\/oNoRDUh9RHWNLhkiF1QJHBdy\\\/MR7tKnyL5oMllmH2sVlT47MlSPh09\",\"card_deadline\":\"2045-03-31 23:59:59\"}]","operator_serial_number":""}
}

/**
 * @Description:测试商品列表获取API
 */
func TestBalance(t *testing.T) {

	appKey := "8DVI1nCM6SEgi0T3HhUI1J2EQJA4sCKAiLCHU5xAuI5YKXZoEd0ysRQdaHU2DNJc"
	appSecret := "0a091b3aa4324435aab703142518a8f7"

	client := millennium_go_sdk.NewApiClient(appKey, appSecret)

	// 设置开发者模式
	client.SetDevelopment(true)

	resp, err := client.CallApi(&balance.QueryBalance{}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Code != 0 {
		fmt.Println(resp.Msg)
		return
	}

	// 转换余额数据
	var credit balance.CreditBalance
	if err = json.Unmarshal([]byte(resp.Data), &credit); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("返回余额查询数据: %s", resp.Data)
}
