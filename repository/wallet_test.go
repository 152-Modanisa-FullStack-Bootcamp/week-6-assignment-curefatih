package repository

import (
	"bootcamp/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {

	username := "test"
	repo := NewWalletRepository()
	err := repo.Create(username, 0)

	assert.NoError(t, err)

	response, err := repo.GetByUsername(username)

	assert.NoError(t, err)
	assert.Equal(t, username, response.Username)
}

func TestRepository_GetByUsername(t *testing.T) {

	wallet := make(map[string]float32)
	username := "test"
	wallet[username] = 100
	repo := newWalletForTest(wallet)

	response, err := repo.GetByUsername(username)

	assert.NoError(t, err)
	assert.Equal(t, username, response.Username)
}

func TestRepository_GetByUsername_NotFound(t *testing.T) {

	username := "test"
	repo := NewWalletRepository()

	response, err := repo.GetByUsername(username)

	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestRepository_GetAll(t *testing.T) {

	actual := []*model.Wallet{
		{
			Username: "test",
			Amount:   0,
		},
		{
			Username: "test2",
			Amount:   100,
		},
	}

	wallet := make(map[string]float32)
	for _, w := range actual {
		wallet[w.Username] = w.Amount
	}
	repo := newWalletForTest(wallet)

	response, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, response, actual)
}

func TestRepository_Update(t *testing.T) {

	wallet := make(map[string]float32)
	username := "test"
	wallet[username] = 100
	repo := newWalletForTest(wallet)

	err := repo.Update(username, float32(100))

	assert.NoError(t, err)

	response, err := repo.GetByUsername(username)

	assert.NoError(t, err)
	assert.Equal(t, float32(100), response.Amount)
}
