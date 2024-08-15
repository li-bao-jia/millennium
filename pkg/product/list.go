package product

type ListProduct struct{}

func (p *ListProduct) GetMethod() string {
	return "product/list"
}

type ListProductParams struct {
	ProductType string `json:"product_type"`
}

type Product struct {
	ID           int    `json:"id"`            // 千禧券商品ID
	ProductName  string `json:"product_name"`  // 商品名称
	ProductType  string `json:"product_type"`  // 商品类型
	Price        string `json:"price"`         // 商品单价
	Discount     string `json:"discount"`      // 折扣(对方技术说这个字段无意义）
	DisplayPrice string `json:"display_price"` // 商品面值
	SalesStatus  string `json:"sales_status"`  // 销售状态：下架、上架、维护中、库存维护
	Storage      int    `json:"storage"`       // 库存状态：断货、警报、充足
	Details      string `json:"details"`       // 商品详情
}
