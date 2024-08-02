package order

type ChargeOrderParams struct {
	ProductId  int    `json:"product_id"`   // 商品编号
	OutOrderNo string `json:"out_order_no"` // 外部订单号
	Account    string `json:"account"`      // 充值账号
	BuyNum     int    `json:"buy_num"`      // 购买数量
}

type ChargeOrder struct{}

func (o *ChargeOrder) GetMethod() string {
	return "order/charge"
}
