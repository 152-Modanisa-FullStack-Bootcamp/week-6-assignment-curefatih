package service

import (
	"bootcamp/mock"
	"bootcamp/model"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceWallet_Create(t *testing.T) {

	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}
	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().Create(response.Username, response.Amount).Return(nil).Times(1)
	repo.EXPECT().GetByUsername(response.Username).Return(nil, nil).Times(1)

	service := NewWalletService(repo, 0, 0)
	err := service.Create(response.Username)

	assert.NoError(t, err)
}

func TestServiceWallet_ShouldReturnNilOnExist(t *testing.T) {
	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}

	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().GetByUsername(response.Username).Return(response, nil).Times(1)

	service := NewWalletService(repo, 0, 0)
	err := service.Create(response.Username)

	// should err nil
	assert.Nil(t, err)
}

func TestServiceWallet_GetByUsername(t *testing.T) {
	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}

	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().GetByUsername(response.Username).Return(response, nil).Times(1)

	service := NewWalletService(repo, 0, 0)
	wallet, err := service.GetByUsername(response.Username)

	assert.NoError(t, err)
	assert.Equal(t, response, wallet)
}

func TestServiceWallet_GetAll(t *testing.T) {
	response := []*model.Wallet{
		{
			Username: "test",
			Amount:   0,
		},
	}

	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().GetAll().Return(response, nil).Times(1)

	service := NewWalletService(repo, 0, 0)
	wallets, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, response, wallets)
}

func TestServiceWallet_Update(t *testing.T) {
	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}

	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().Update(response.Username, response.Amount).Return(nil).Times(1)
	repo.EXPECT().GetByUsername(response.Username).Return(response, nil).Times(2)

	service := NewWalletService(repo, 0, 0)
	err := service.Update(response.Username, response.Amount)

	assert.NoError(t, err)

	wallet, err := service.GetByUsername(response.Username)

	assert.NoError(t, err)
	assert.Equal(t, response, wallet)
}

func TestServiceWallet_GetByUsername_NotFound(t *testing.T) {
	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}

	errResponse := errors.New("user not found")
	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().GetByUsername(response.Username).Return(nil, errResponse).Times(1)

	service := NewWalletService(repo, 0, 0)
	wallet, err := service.GetByUsername(response.Username)

	assert.Error(t, err)
	assert.Nil(t, wallet)
	assert.Equal(t, "user not found", errResponse.Error())
}

func TestServiceWallet_Update_NotFound(t *testing.T) {
	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}

	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().GetByUsername(response.Username).Return(nil, nil).Times(1)

	service := NewWalletService(repo, 0, 0)
	err := service.Update(response.Username, response.Amount)

	// should user not found
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestServiceWallet_Update_AmountNotValid(t *testing.T) {
	response := &model.Wallet{
		Username: "test",
		Amount:   0,
	}

	repo := mock.NewMockIWalletRepository(gomock.NewController(t))
	repo.EXPECT().GetByUsername(response.Username).Return(response, nil).Times(1)

	service := NewWalletService(repo, 0, 0)
	err := service.Update(response.Username, -1)

	// should amount not valid
	assert.Error(t, err)
	assert.Equal(t, "amount is not valid", err.Error())
}
