package pointers

import (
	"errors"
	"fmt"
)

//type Name OriginalType

type Bitcoin float64

type Wallet struct {
	balance Bitcoin // lowercase balance is private outside the package pointers
}

func (b Bitcoin) String() string { //implement Stringer on Bitcoin to print %s
	return fmt.Sprintf("%.2f BTC", b)
}

func (w *Wallet) Deposit(d Bitcoin) {
	w.balance += d
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

const ErrInsufficientFunds string = "cannot withdraw, insufficient funds"

func (w *Wallet) Withdraw(d Bitcoin) error {
	if w.balance < d {
		return errors.New(ErrInsufficientFunds)
	}

	w.balance -= d
	return nil
}
