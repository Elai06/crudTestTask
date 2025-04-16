package server

import (
	"bytes"
	"crudTestTask/internal/repository"
	"crudTestTask/internal/repository/mocks"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockDataBaseHandler(ctrl)
	mockRepo := New(mockHandler)
	fakeData := repository.Data{
		Id:   gofakeit.Int64(),
		Name: gofakeit.FirstName()}

	tests := []struct {
		name           string
		expectedResult repository.Data
		inputData      repository.Data
		setupMock      func()
		expectedError  bool
	}{
		{
			name:           "success create user",
			expectedResult: fakeData,
			inputData:      fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Create(fakeData).Return(&fakeData, nil)
			},
			expectedError: false,
		},
		{
			name:           "failed create user",
			expectedResult: fakeData,
			inputData:      fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Create(fakeData).Return(&repository.Data{}, nil)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.inputData)
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBody))
			rec := httptest.NewRecorder()

			if tt.setupMock != nil {
				tt.setupMock()
			}

			mockRepo.createUser(rec, req)
			result := repository.Data{}
			err := json.Unmarshal(rec.Body.Bytes(), &result)

			if tt.expectedError {
				assert.NotEqual(t, result, tt.inputData)
				assert.Error(t, fmt.Errorf("failed to create user"))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockDataBaseHandler(ctrl)
	mockRepo := New(mockHandler)
	fakeData := repository.Data{
		Id:   gofakeit.Int64(),
		Name: gofakeit.FirstName()}

	tests := []struct {
		name           string
		expectedResult repository.Data
		inputData      int64
		setupMock      func()
		expectedError  bool
	}{
		{
			name:           "success get user",
			inputData:      fakeData.Id,
			expectedResult: fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Get(fakeData.Id).Return(&fakeData, nil)
			},
			expectedError: false,
		},
		{
			name:           "failed get user",
			expectedResult: fakeData,
			inputData:      fakeData.Id,
			setupMock: func() {
				mockHandler.EXPECT().Get(fakeData.Id).Return(&repository.Data{}, nil)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/users?user_id="+fmt.Sprint(tt.inputData), nil)
			rec := httptest.NewRecorder()

			if tt.setupMock != nil {
				tt.setupMock()
			}

			mockRepo.getUser(rec, req)
			result := repository.Data{}
			err := json.Unmarshal(rec.Body.Bytes(), &result)

			if tt.expectedError {
				assert.NotEqual(t, result, tt.expectedResult)
				assert.Error(t, fmt.Errorf("failed to create user"))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockDataBaseHandler(ctrl)
	mockRepo := New(mockHandler)
	fakeData := repository.Data{
		Id:   gofakeit.Int64(),
		Name: gofakeit.FirstName()}

	tests := []struct {
		name           string
		expectedResult repository.Data
		inputData      repository.Data
		setupMock      func()
		expectedError  bool
	}{
		{
			name:           "success update user",
			expectedResult: fakeData,
			inputData:      fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Update(fakeData).Return(&fakeData, nil)
			},
			expectedError: false,
		},
		{
			name:           "failed update user",
			expectedResult: repository.Data{Name: fakeData.Name, Id: fakeData.Id},
			inputData:      fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Update(fakeData).Return(&repository.Data{}, nil)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.inputData)
			req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewReader(reqBody))
			rec := httptest.NewRecorder()

			if tt.setupMock != nil {
				tt.setupMock()
			}

			mockRepo.updateUser(rec, req)
			result := repository.Data{}
			err := json.Unmarshal(rec.Body.Bytes(), &result)

			if tt.expectedError {
				assert.NotEqual(t, result, tt.inputData)
				assert.Error(t, fmt.Errorf("failed to create user"))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := mocks.NewMockDataBaseHandler(ctrl)
	mockRepo := New(mockHandler)
	fakeData := repository.Data{
		Id: gofakeit.Int64()}

	tests := []struct {
		name          string
		inputData     repository.Data
		setupMock     func()
		expectedError bool
	}{
		{
			name:      "success delete user",
			inputData: fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Delete(fakeData.Id).Return(nil)
			},
			expectedError: false,
		},
		{
			name:      "failed delete user",
			inputData: fakeData,
			setupMock: func() {
				mockHandler.EXPECT().Delete(fakeData.Id).Return(fmt.Errorf("failed to delete user"))
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/users?user_id="+fmt.Sprint(tt.inputData.Id), nil)
			rec := httptest.NewRecorder()

			if tt.setupMock != nil {
				tt.setupMock()
			}

			mockRepo.deleteUser(rec, req)
			result := ""
			err := json.Unmarshal(rec.Body.Bytes(), &result)

			if tt.expectedError {
				assert.Error(t, fmt.Errorf("failed to create user"))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
