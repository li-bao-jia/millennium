<p align="center">millennium SDK</p>
<p align="center">千禧券GO SDK，用于对接千禧券供应链平台</p>


### 项目概述
- 初衷：对接时发现没有 GO 的 SDK，反正要写，直接封装了
- 设计：不喜欢一件事做两次，封装了这个SDK提供给所有开发者
- 特点：开箱即用，有完整的示例代码，直接复制就可以完成对接
- 功能：用于对接千禧券供应链平台API，商品查询、订单推送、订单查询，余额查询等


## 安装 Installation

你可以直接使用 go get 安装：

```
go get github.com/li-bao-jia/millennium@latest
```

## 快速开始 Quick Start

### 获取商品列表接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/product"
)

func main() {
	client := millennium.NewApiClient("appKey", "appSecret")

	// 设置是否使用http，ture为http，false为https
	// client.SetHttp(true)

	// 设置开发者模式，true为开发模式，false为正式模式
	// client.SetDev(true)

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

```

### 直充下单接口

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("appKey", "appSecret")

	// 设置开发者模式
	// client.SetDev(true)

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
}

```

### 卡密下单接口

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("appKey", "appSecret")

	// 设置开发者模式
	// client.SetDev(true)
	
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
}

```

### 订单查询接口

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("appKey", "appSecret")

	// 设置开发者模式
	// client.SetDev(true)
	
	// 设置请求参数
	data := order.QueryOrderParams{
		OutOrderNo: "20240801002", // 单号
	}

	resp, err := client.CallApi(&order.QueryOrder{}, data)
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

	// 转换卡密数据 千禧券数据1卡密2直充
	if cOrder.OrderType == 1 {
		var cards []order.Card
		if err = json.Unmarshal([]byte(cOrder.Cards), &cards); err != nil {
			return
		}
		fmt.Printf("返回卡密数量: %d", len(cards))
		fmt.Printf("返回卡密数据: %s", cOrder.Cards)

		// 卡密数据使用时，需要解密显示
		for i, card := range cards {
			// 卡密解密
			var CardPwd, CardNumber string
			if CardPwd, err = client.DecryptAES256ECB(card.CardPwd); err != nil {
				return
			}
			if CardNumber, err = client.DecryptAES256ECB(card.CardNumber); err != nil {
				return
			}
			cards[i].CardPwd = CardPwd
			cards[i].CardNumber = CardNumber
		}
	}
	fmt.Printf("返回订单查询数据: %s", resp.Data)
}

```

### 查询余额接口

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/balance"
)

func main() {
	client := millennium.NewApiClient("appKey", "appSecret")

	// 设置开发者模式
	// client.SetDev(true)
	
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

```

### 联系方式

- DEVELOPER: BaoJia Li

- QQ: 751818588

- QQ群: 232185834

- EMAIL: livsyitian@163.com
