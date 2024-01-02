package service_test

import (
	"context"
	"errors"
	"fizzbuzz/internal/domain/entity"
	"fizzbuzz/internal/domain/service"
	"fizzbuzz/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CalculateFizzBuzz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		mockCalls func(m *mocks.MockFizzBuzzStorage)
		entity    *entity.FizzBuzz
		want      []string
		wantError string
	}{
		{
			name: "OK",
			mockCalls: func(m *mocks.MockFizzBuzzStorage) {
				m.EXPECT().SaveFizzbuzz(gomock.Any(), gomock.Any()).Return(nil)
			},
			entity: &entity.FizzBuzz{
				Int1:  3,
				Int2:  5,
				Limit: 5,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			want: []string{
				"1",
				"2",
				"fizz",
				"4",
				"buzz",
			},
		},
		{
			name: "KO: failed to savefizzbuzz",
			mockCalls: func(m *mocks.MockFizzBuzzStorage) {
				m.EXPECT().SaveFizzbuzz(gomock.Any(), gomock.Any()).Return(errors.New("failed"))
			},
			entity: &entity.FizzBuzz{
				Int1:  3,
				Int2:  5,
				Limit: 5,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			want:      nil,
			wantError: "failed",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			m := mocks.NewMockFizzBuzzStorage(ctrl)

			test.mockCalls(m)
			s := service.NewFizzBuzzService(m)
			got, err := s.CalculateFizzBuzz(ctx, test.entity)

			if err != nil {
				assert.EqualError(t, err, test.wantError)
			}

			assert.Equal(t, test.want, got)
		})
	}
}
