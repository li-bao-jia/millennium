package order

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
)

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

// DecryptAES256ECB 实现 AES-256 ECB 模式解密
func DecryptAES256ECB(pass, secret string) (str string, err error) {
	cardNumberRes, err := base64.StdEncoding.DecodeString(pass)
	if err != nil {
		return
	}

	r, err := openssl.AesECBDecrypt(cardNumberRes, []byte(secret), openssl.PKCS7_PADDING)
	if err != nil {
		return
	}

	str = string(r)

	return
}
