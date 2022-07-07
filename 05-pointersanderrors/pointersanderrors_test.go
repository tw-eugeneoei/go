package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("desposit", func(t *testing.T) {
		expected := Bitcoin(10)

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw sufficient amount", func(t *testing.T) {
		expected := Bitcoin(5)

		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient amount", func(t *testing.T) {
		expected := Bitcoin(10)
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, insufficientFundsErr.Error())
		assertBalance(t, wallet, expected)
	})

}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()
	if expected != result {
		t.Errorf("expected '%s' got '%s'", expected, result)
	}
}

func assertError(t testing.TB, err error, expected string) {
	t.Helper()
	// * "nil" is synonymous with "null"
	// * Errors can be "nil" because the return type of Withdraw will be error, which is an interface.
	// * If you see a function that takes arguments or returns values that are interfaces, they can be nillable.
	if err == nil {
		// ! t.Fatal will stop the test if it is called
		t.Fatal("should throw an error")
	}

	if err.Error() != expected {
		t.Errorf("expected '%s' got '%s'", expected, err.Error())
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("should not throw error")
	}
}
