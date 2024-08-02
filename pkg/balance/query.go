package balance

type CreditBalance struct {
	Balance string `json:"balance"`
}

type QueryBalance struct{}

func (b *QueryBalance) GetMethod() string {
	return "balance"
}
