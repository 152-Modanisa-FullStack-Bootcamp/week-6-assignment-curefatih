package handler

import (
	"bootcamp/service"
	"encoding/json"
	"net/http"
)

type IWalletHandler interface {
	GetByUsername(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	HandleFunc(w http.ResponseWriter, r *http.Request)
}

type WalletHandler struct {
	walletService service.IWalletService
}

func NewHandler(walletService service.IWalletService) IWalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {

	// get path /:username
	username := r.URL.Path[len("/"):]
	rawResponse, err := h.walletService.GetByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// struct to json
	response, err := json.Marshal(rawResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// write response
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *WalletHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rawResponse, err := h.walletService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// struct to json
	response, err := json.Marshal(rawResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// write response
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *WalletHandler) Create(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/"):]
	err := h.walletService.Create(username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// write response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func (h *WalletHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/"):]

	var body struct {
		Balance float32 `json:"balance"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.walletService.Update(username, body.Balance)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated"))
}

// wrapper func for request type
func (h *WalletHandler) HandleFunc(w http.ResponseWriter, r *http.Request) {
	// if only /
	if r.URL.Path == "/" && r.Method == "GET" {
		h.GetAll(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetByUsername(w, r)
	case http.MethodPut:
		h.Create(w, r)
	case http.MethodPost:
		h.Update(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
