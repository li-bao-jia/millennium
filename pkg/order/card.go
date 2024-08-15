package order

type CardOrderParams struct {
	ProductId  int    `json:"product_id"`   // 商品编号
	OutOrderNo string `json:"out_order_no"` // 外部订单号
	BuyNum     int    `json:"buy_num"`      // 购买数量
}

type Card struct {
	CardNumber   string `json:"card_number"`
	CardPwd      string `json:"card_pwd"`
	CardDeadline string `json:"card_deadline"`
}

type CardOrder struct{}

func (o *CardOrder) GetMethod() string {
	return "order/card"
}
