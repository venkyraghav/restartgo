package pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientBalance = fmt.Errorf("cannot withdraw. insufficient balance")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientBalance
	}
	w.balance -= amount
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
