package handlers

import (
	"bytes"
	"context"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestHandler_FetchJobById(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &Handler{
				service: ms,
			}

			h.FetchJobById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func TestHandler_FetchJobByCompanyId(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.FetchJobByCompanyId(tt.args.c)
		})
	}
}

func TestHandler_ProcessJobApp(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
				rr := httptest.NewRecorder()
				//response of the request
				c, _ := gin.CreateTestContext(rr)
				//its just a mock context bcz we didn't use the original context
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", nil)
				//give the input of the for the method
				c.Request = httpRequest
				//updating the request
				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "decoder",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
				rr := httptest.NewRecorder()
				//response of the request
				c, _ := gin.CreateTestContext(rr)
				//its just a mock context bcz we didn't use the original context
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", bytes.NewBufferString(``))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				//give the input of the for the method
				c.Request = httpRequest
				//updating the request
				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "error in validate",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", strings.NewReader(`[
					{
					  "id": 1,
					  "role": "Software Engineer",
					  "notice_period": 30,
					  "budget": 8000,
					  "locations": [1, 2],
					  "technology": [1],
					  "workmode": [1, 2],
					  "experience": 4,
					  "qualification": [1, 2],
					  "shift": [1, 2],
					  "jobtype": [1, 2]
					},
					{
					  "role": "Data Scientist",
					  "notice_period": 15,
					  "budget": 8000,
					  "locations": [1,2],
					  "technology": [1,2],
					  "workmode": [1, 2],
					  "experience": 4,
					  "qualification": [2],
					  "shift": [1,2],
					  "jobtype": [1, 2]
					}
				  ]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				// mc := gomock.NewController(t)
				// ms := .NewMockUserService(mc)
				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"please provide job role and Description"}`,
		},
		{
			name: "data not found",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", strings.NewReader(`[
					{
					  "id": 1,
					  "role": "Software Engineer",
					  "notice_period": 30,
					  "budget": 8000,
					  "locations": [1, 2],
					  "technology": [1],
					  "workmode": [1, 2],
					  "experience": 4,
					  "qualification": [1, 2],
					  "shift": [1, 2],
					  "jobtype": [1, 2]
					},
					{
						"id": 1,
						"role": "Data Scientist",
						"notice_period": 15,
						"budget": 8000,
						"locations": [1,2],
						"technology": [1,2],
						"workmode": [1, 2],
						"experience": 4,
						"qualification": [2],
						"shift": [1,2],
						"jobtype": [1, 2]
					}
				  ]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := services.NewMockServiceMethod(mc)
				ms.EXPECT().ProcessJob(gomock.Any()).Return(nil, nil).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"data is not same"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", strings.NewReader(`[
					{
					  "id": 1,
					  "role": "Software Engineer",
					  "notice_period": 30,
					  "budget": 8000,
					  "locations": [1, 2],
					  "technology": [1],
					  "workmode": [1, 2],
					  "experience": 4,
					  "qualification": [1, 2],
					  "shift": [1, 2],
					  "jobtype": [1, 2]
					},
					{
						"id": 1,
						"role": "Data Scientist",
						"notice_period": 15,
						"budget": 8000,
						"locations": [1,2],
						"technology": [1,2],
						"workmode": [1, 2],
						"experience": 4,
						"qualification": [2],
						"shift": [1,2],
						"jobtype": [1, 2]
					}
				  ]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := services.NewMockServiceMethod(mc)
				ms.EXPECT().ProcessJob(gomock.Any()).Return([]models.RequestJob{}, nil).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &Handler{
				service: ms,
			}

			h.ProcessJobApp(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
