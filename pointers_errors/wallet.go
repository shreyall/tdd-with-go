package banking

import (
	"errors"
	"fmt"
)

// Bitcoin - same type as int
type Bitcoin int

// Stringer - stringer method on the Bitcoin type
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct to hold our mulla $
type Wallet struct {
	balance Bitcoin
}

// Deposit into a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance of a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds - error string for excess withdrawal
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw - removes mulla from your wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
