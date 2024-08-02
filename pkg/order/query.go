package order

type QueryOrderParams struct {
	OutOrderNo string `json:"out_order_no"`
}

type Order struct {
	Cards                string `json:"cards"`                  // 卡密信息，仅卡密订单返回
	OrderNo              string `json:"order_no"`               // 订单编号
	OutOrderNo           string `json:"out_order_no"`           // 外部订单号，每次请求必须唯一
	ProductId            int    `json:"product_id"`             // 商品Id
	ProductName          string `json:"product_name"`           // 商品名称
	BuyNum               int    `json:"buy_num"`                // 购买数量
	OrderType            int    `json:"order_type"`             // 订单类型：1-卡密 2-直充
	OrderPrice           string `json:"order_price"`            // 交易单价
	OrderState           string `json:"order_state"`            // 订单状态： （success：成功，processing：处理中，failed：失败）
	OrderServiceState    string `json:"order_service_state"`    // 订单售后状态：（null：无售后；processing：售后处理中；finished：处理完成
	Account              string `json:"account"`                // 充值账号
	CreateTime           string `json:"create_time"`            // 创建时间
	FinishTime           string `json:"finish_time"`            // 订单完成时间，查单接口返回
	OperatorSerialNumber string `json:"operator_serial_number"` // 运营商流水号
}

type QueryOrder struct{}

func (o *QueryOrder) GetMethod() string {
	return "order/query"
}
