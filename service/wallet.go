package service

import (
	"bootcamp/model"
	"bootcamp/repository"
	"errors"
)

type IWalletService interface {
	GetByUsername(username string) (*model.Wallet, error)
	GetAll() ([]*model.Wallet, error)
	Create(username string) error
	Update(username string, amount float32) error
}

type WalletService struct {
	WalletRepository     repository.IWalletRepository
	initialBalanceAmount float32
	minimumBalanceAmount float32
}

func NewWalletService(walletRepository repository.IWalletRepository, initialBalanceAmount, minimumBalanceAmount float32) *WalletService {
	return &WalletService{
		WalletRepository:     walletRepository,
		initialBalanceAmount: initialBalanceAmount,
		minimumBalanceAmount: minimumBalanceAmount,
	}
}

func (w *WalletService) GetByUsername(username string) (*model.Wallet, error) {
	return w.WalletRepository.GetByUsername(username)
}

func (w *WalletService) GetAll() ([]*model.Wallet, error) {
	return w.WalletRepository.GetAll()
}

func (w *WalletService) Create(username string) error {
	user, err := w.WalletRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if user != nil {
		return nil
	}

	return w.WalletRepository.Create(username, w.initialBalanceAmount)
}

func (w *WalletService) Update(username string, balance float32) error {
	user, err := w.WalletRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	newAmount := user.Amount + balance
	// check if amount is valid
	if newAmount < w.minimumBalanceAmount {
		return errors.New("amount is not valid")
	}

	return w.WalletRepository.Update(username, newAmount)
}
