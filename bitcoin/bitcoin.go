package bitcoin

import (
	"fmt"
)

type Wallet struct {
	balance Bitcoin
	name    string
}

type Bitcoin int

func (w *Wallet) Balance() Bitcoin {
	return w.balance

}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return generateOverdrawError(amt, w.balance)
	}
	w.balance -= amt
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func generateOverdrawError(amt, balance Bitcoin) error {
	return fmt.Errorf("withdrawal amount %s is greater than wallet balance %s", amt.String(), balance.String())
}
