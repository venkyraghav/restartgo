package pointers

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		assert.Equal(t, want, got)
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with sufficient", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Error("No error expected but got one")
		}
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with insufficient", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		err := wallet.Withdraw(Bitcoin(10))
		assertError(t, err, ErrInsufficientBalance)
		assertBalance(t, wallet, Bitcoin(5))
	})
}
