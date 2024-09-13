//write a test, for depositing 10 bitcoins

package bitcoin

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{balance: 0, name: "Wen's wallet"}
	t.Run("deposit", func(t *testing.T) {
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		withdrawOperationReturn := wallet.Withdraw(200)
		assertNoError(t, withdrawOperationReturn)
		got := wallet.Balance()
		want := Bitcoin(0)
		assertBalance(t, got, want)
	})

	t.Run("overdraw", func(t *testing.T) {
		got := wallet.Withdraw(10)
		want := generateOverdrawError(Bitcoin(10), wallet.balance)
		assertError(t, got, want)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error, but didn't get one")
	}
	if got.Error() != want.Error() {
		t.Errorf("Got ErrorString: %s, Want ErrorString: %s", got.Error(), want.Error())
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Error("Got and error, but didn't want one")
	}
}

func assertBalance(t testing.TB, got, want Bitcoin) {
	if got != want {
		t.Errorf("Got %s, Want %s", got.String(), want.String())
	}
}
