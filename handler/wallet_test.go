package handler

import (
	"bootcamp/mock"
	"bootcamp/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerWallet_Create(t *testing.T) {

	testData := &model.Wallet{
		Username: "test",
		Amount:   100,
	}

	testDataJSON, _ := json.Marshal(testData)

	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().Create(testData.Username).Return(nil)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodPut, "/"+testData.Username, bytes.NewReader(testDataJSON))
	w := httptest.NewRecorder()
	handler.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestHandlerWallet_Create_Error(t *testing.T) {

	err := errors.New("error")
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().Create(gomock.Any()).Return(err)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	handler.Create(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Error(t, err)
}

func TestHandlerWallet_GetByUsername(t *testing.T) {

	testData := &model.Wallet{
		Username: "test",
		Amount:   100,
	}

	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().GetByUsername(testData.Username).Return(testData, nil)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/"+testData.Username, nil)
	w := httptest.NewRecorder()
	handler.GetByUsername(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response *model.Wallet
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, testData, response)
}

func TestHandlerWallet_GetByUsername_Error(t *testing.T) {

	testUsername := "test"
	err := errors.New("error")
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().GetByUsername(testUsername).Return(nil, err)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/"+testUsername, nil)
	w := httptest.NewRecorder()
	handler.GetByUsername(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandlerWallet_GetAll(t *testing.T) {

	testData := []*model.Wallet{
		{
			Username: "test",
			Amount:   100,
		},
	}

	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().GetAll().Return(testData, nil)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAll(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []*model.Wallet
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, testData, response)
}

func TestHandlerWallet_GetAll_Error(t *testing.T) {

	err := errors.New("error")
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().GetAll().Return(nil, err)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.GetAll(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandlerWallet_Update(t *testing.T) {

	testUsername := "test"
	testData := map[string]interface{}{
		"balance": float32(100),
	}

	testDataJSON, _ := json.Marshal(testData)

	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().Update(testUsername, testData["balance"]).Return(nil)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUsername, bytes.NewReader(testDataJSON))
	w := httptest.NewRecorder()
	handler.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestHandlerWallet_Update_Error(t *testing.T) {

	testUsername := "test"
	testData := map[string]interface{}{
		"balance": float32(100),
	}

	testDataJSON, _ := json.Marshal(testData)

	err := errors.New("error")
	service := mock.NewMockIWalletService(gomock.NewController(t))
	service.EXPECT().Update(testUsername, testData["balance"]).Return(err)

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUsername, bytes.NewReader(testDataJSON))
	w := httptest.NewRecorder()
	handler.Update(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandlerWallet_HandleFunc(t *testing.T) {

	service := mock.NewMockIWalletService(gomock.NewController(t))

	testCases := []struct {
		name   string
		method string
		path   string
		body   []byte
		status int
		EXPECT *gomock.Call
	}{
		{
			name:   "Create",
			method: http.MethodPut,
			path:   "/test",
			body:   []byte(`{"username":"test","amount":100}`),
			status: http.StatusCreated,
			EXPECT: service.EXPECT().Create("test"),
		},
		{
			name:   "GetByUsername",
			method: http.MethodGet,
			path:   "/test",
			body:   nil,
			status: http.StatusOK,
			EXPECT: service.EXPECT().GetByUsername("test"),
		},
		{
			name:   "GetAll",
			method: http.MethodGet,
			path:   "/",
			body:   nil,
			status: http.StatusOK,
			EXPECT: service.EXPECT().GetAll(),
		},
		{
			name:   "Update",
			method: http.MethodPost,
			path:   "/test",
			body:   []byte(`{"balance":100}`),
			status: http.StatusOK,
			EXPECT: service.EXPECT().Update("test", float32(100)),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			handler := NewHandler(service)
			r := httptest.NewRequest(testCase.method, testCase.path, bytes.NewReader(testCase.body))
			w := httptest.NewRecorder()

			handler.HandleFunc(w, r)

			assert.Equal(t, testCase.status, w.Code)
		})
	}

	handler := NewHandler(service)
	r := httptest.NewRequest(http.MethodDelete, "/", nil)
	w := httptest.NewRecorder()

	handler.HandleFunc(w, r)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

}
