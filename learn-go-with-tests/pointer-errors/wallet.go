package pointer

type Wallet struct {
	Amount float64 `json:"amout"`
}

func (w *Wallet) Balance() float64 {
	return w.Amount
}

func (w *Wallet) Deposit(money float64) {
	w.Amount += money
}
