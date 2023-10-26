package tutorial

import (
	"context"
	"testing"
	"time"

	"github.com/salmaqnsGH/my-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	got, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, account1.ID, got.ID)
	require.Equal(t, account1.Owner, got.Owner)
	require.Equal(t, account1.Balance, got.Balance)
	require.Equal(t, account1.Currency, got.Currency)

	require.WithinDuration(t, account1.CreatedAt, got.CreatedAt, time.Second)
}