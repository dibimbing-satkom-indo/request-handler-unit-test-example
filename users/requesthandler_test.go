package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"request-handler-unit-test-example/helper"
	mocks "request-handler-unit-test-example/test/mocks/users"
	"request-handler-unit-test-example/users/dto"
	"testing"
)

func TestRequestHandler_GetUsers(t *testing.T) {
	type fields struct {
		controller Controller
	}

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "error",
			expectedStatusCode: http.StatusInternalServerError,
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/users", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				err := errors.New("error")
				mockController.EXPECT().
					GetUsers().
					Return(dto.GetUsersResponse{}, err).
					Once()
				return fields{controller: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, dto.MessageResponse{Message: "error"}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				controller: f.controller,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/users", h.GetUsers)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_CreateUser(t *testing.T) {
	type fields struct {
		controller Controller
	}

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "error",
			expectedStatusCode: http.StatusInternalServerError,
			makeRequest: func() *http.Request {
				body, _ := json.Marshal(dto.CreateUserRequest{Name: "john"})
				request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
				request.Header.Set("Content-Type", "application/json")
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				err := errors.New("error")
				mockController.EXPECT().
					CreateUser(dto.CreateUserRequest{Name: "john"}).
					Return(dto.CreateUserResponse{}, err).
					Once()
				return fields{controller: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, dto.MessageResponse{Message: "error"}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				controller: f.controller,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/users", h.CreateUser)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}
