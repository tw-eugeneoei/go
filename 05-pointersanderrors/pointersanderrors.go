package pointersanderrors

import (
	"errors"
	"fmt"
)

// * var keyword allows us to define values global to the package
var insufficientFundsErr = errors.New("Insufficient funds.")

type Bitcoin int

// * syntax for creating a method on a type declaration is the same as it is on a struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return insufficientFundsErr
	}
	w.balance -= amt
	return nil
}
