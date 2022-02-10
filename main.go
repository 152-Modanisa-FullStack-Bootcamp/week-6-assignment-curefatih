package main

import (
	"bootcamp/config"
	"bootcamp/handler"
	"bootcamp/repository"
	"bootcamp/service"
	"net/http"
)

func main() {
	config := config.Get()

	repo := repository.NewWalletRepository()
	service := service.NewWalletService(repo, config.InitialBalanceAmount, config.MinimumBalanceAmount)
	h := handler.NewHandler(service)

	http.HandleFunc("/", h.HandleFunc)
	err := http.ListenAndServe(config.ServerAddr, nil)
	if err != nil {
		panic(err)
	}
}
