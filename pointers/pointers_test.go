package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Make a deposit of 10", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBallance(t, wallet, want)

	})

	t.Run("Make a withdraw of 5", func(t *testing.T) {

		wallet := Wallet{}

		wallet.balance = Bitcoin(10)
		wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)

		assertBallance(t, wallet, want)

	})

	t.Run("Throw error if insufficient fonds", func(t *testing.T) {

		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(50))

		assertError(t, err, ErrInsufficientFonds.Error())
		assertBallance(t, wallet, startingBalance)

	})

}
