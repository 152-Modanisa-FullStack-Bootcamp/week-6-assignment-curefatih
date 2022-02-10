package repository

import (
	"bootcamp/model"
	"errors"
)

type IWalletRepository interface {
	GetByUsername(username string) (*model.Wallet, error)
	GetAll() ([]*model.Wallet, error)
	Create(username string, amount float32) error
	Update(username string, amount float32) error
}

type WalletRepository struct {
	wallet map[string]float32
}

func NewWalletRepository() IWalletRepository {
	return &WalletRepository{
		wallet: make(map[string]float32),
	}
}

func newWalletForTest(wallet map[string]float32) *WalletRepository {
	return &WalletRepository{
		wallet: wallet,
	}
}

func (w *WalletRepository) GetByUsername(username string) (*model.Wallet, error) {
	amount, ok := w.wallet[username]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &model.Wallet{
		Username: username,
		Amount:   amount,
	}, nil
}

func (w *WalletRepository) GetAll() ([]*model.Wallet, error) {
	var wallets []*model.Wallet

	for username, amount := range w.wallet {
		wallets = append(wallets, &model.Wallet{
			Username: username,
			Amount:   amount,
		})
	}

	return wallets, nil
}

func (w *WalletRepository) Create(username string, amount float32) error {
	w.wallet[username] = amount
	return nil
}

func (w *WalletRepository) Update(username string, amount float32) error {
	w.wallet[username] = amount
	return nil
}
