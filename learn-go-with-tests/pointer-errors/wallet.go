package pointer

import "fmt"

type Bitcoin float64

type Wallet struct {
	Amount Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).Amount
}

func (w *Wallet) Deposit(money Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.Amount)
	w.Amount += money
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.Amount -= amount
}

func (b *Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
