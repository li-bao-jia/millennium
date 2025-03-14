<div align=center>
  <p align="center">millennium SDK</p>
  <p align="center">千禧券GO SDK，用于对接千禧券供应链平台</p>
</div>

<div align=center>
  <p align="center">
    <a href="https://github.com/li-bao-jia">
      <img src="https://img.shields.io/badge/go-1.21.8-blue" alt="Build Status">
    </a>
    <a href="https://github.com/li-bao-jia">
      <img src="https://img.shields.io/github/license/li-bao-jia/millennium" alt="License">
    </a>
  </p>
</div>


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

### 查询余额接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/balance"
)

func main() {
	client := millennium.NewApiClient("domain", "appId", "appSecret")

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

```

### 获取商品列表接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/product"
)

func main() {
	client := millennium.NewApiClient("domain", "appId", "appSecret")

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

```

### 直充下单接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("domain", "appId", "appSecret")
	
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

```

### 话费直充接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("domain", "appId", "appSecret")
	
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

```

### 卡密下单接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("domain", "appId", "appSecret")

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

```

### 订单查询接口

```go
package main

import (
	"fmt"
	"github.com/li-bao-jia/millennium"
	"github.com/li-bao-jia/millennium/pkg/order"
)

func main() {
	client := millennium.NewApiClient("domain", "appId", "appSecret")

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

```


### 联系方式

- 开发者: BaoJia Li

- QQ: 751818588

- QQ群: 232185834

- 邮箱: livsyitian@163.com
