package handler_test

import (
	"errors"
	"fizzbuzz/internal/domain/entity"
	"fizzbuzz/internal/handler"
	"fizzbuzz/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateFizzBuzzString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		mockCall   func(*mocks.MockFizzBuzzService)
		url        string
		wantStatus int
		wantError  error
	}{
		{
			name: "OK",
			mockCall: func(m *mocks.MockFizzBuzzService) {
				m.EXPECT().CalculateFizzBuzz(gomock.Any(), gomock.Any()).Return(
					[]string{}, nil,
				)
			},
			url:        "/fizzbuzz/v1/fizzbuzz?int1=3&int2=5&limit=18&str1=fizz&str2=buzz",
			wantStatus: http.StatusOK,
			wantError:  nil,
		},
		{
			name:       "Bad parameter: invalid input 1",
			mockCall:   func(m *mocks.MockFizzBuzzService) {},
			url:        "/fizzbuzz/v1/fizzbuzz?int1=not_integer&int2=5&limit=18&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  errors.New(handler.ErrInvalidFizzBuzzParams + "\n"),
		},
		{
			name:       "Bad parameter: invalid input 2",
			mockCall:   func(m *mocks.MockFizzBuzzService) {},
			url:        "/fizzbuzz/v1/fizzbuzz?int1=3&int2=5&limit=not_integer&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  errors.New(handler.ErrInvalidFizzBuzzParams + "\n"),
		},
		{
			name: "Bad parameter: invalid input 3",
			mockCall: func(m *mocks.MockFizzBuzzService) {
				m.EXPECT().CalculateFizzBuzz(gomock.Any(), gomock.Any()).Return(
					nil, errors.New("can't validate the input\n"),
				)
			},
			url:        "/fizzbuzz/v1/fizzbuzz?int1=3&int2=-5&limit=15&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  errors.New(handler.ErrFizzBuzzNotValidate + "\n"),
		},
		{
			name: "Failed to calculate fizzbuzz",
			mockCall: func(m *mocks.MockFizzBuzzService) {
				m.EXPECT().CalculateFizzBuzz(gomock.Any(), gomock.Any()).Return(
					nil, errors.New("failed to calculate fizzbuzz"),
				)
			},
			url:        "/fizzbuzz/v1/fizzbuzz?int1=3&int2=19&limit=15&str1=fizz&str2=buzz",
			wantStatus: http.StatusInternalServerError,
			wantError:  errors.New("failed to calculate fizzbuzz\n"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			serviceMock := mocks.NewMockFizzBuzzService(ctrl)
			test.mockCall(serviceMock)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, test.url, nil)

			router := chi.NewRouter()
			handler.NewFizzBuzzHandler(serviceMock, router)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, test.wantStatus, recorder.Code)
			if recorder.Code != http.StatusOK {
				assert.EqualError(t, test.wantError, recorder.Body.String())
			}
		})
	}

}

func Test_GetStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		mockCall   func(*mocks.MockFizzBuzzService)
		url        string
		wantStatus int
		wantError  error
	}{
		{
			name: "OK",
			mockCall: func(m *mocks.MockFizzBuzzService) {
				m.EXPECT().GetMostFrequentRequest(gomock.Any()).Return(
					&entity.FizzBuzz{}, 5, nil,
				)
			},
			url:        "/fizzbuzz/v1/stats",
			wantStatus: http.StatusOK,
			wantError:  nil,
		},
		{
			name: "Not found",
			mockCall: func(m *mocks.MockFizzBuzzService) {
				m.EXPECT().GetMostFrequentRequest(gomock.Any()).Return(
					nil, 0, errors.New("can't find the result"),
				)
			},
			url:        "/fizzbuzz/v1/stats",
			wantStatus: http.StatusInternalServerError,
			wantError:  errors.New("can't find the result\n"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			serviceMock := mocks.NewMockFizzBuzzService(ctrl)
			test.mockCall(serviceMock)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, test.url, nil)

			router := chi.NewRouter()
			handler.NewFizzBuzzHandler(serviceMock, router)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, test.wantStatus, recorder.Code)
			if recorder.Code != http.StatusOK {
				assert.EqualError(t, test.wantError, recorder.Body.String())
			}
		})
	}

}
