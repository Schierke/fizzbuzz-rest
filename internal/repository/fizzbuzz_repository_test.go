package repository_test

import (
	"context"
	"errors"
	"fizzbuzz/internal/domain/entity"
	"fizzbuzz/internal/repository"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

func Test_SaveFizzbuzz(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		fb        *entity.FizzBuzz
		mockCall  func(*entity.FizzBuzz, pgxmock.PgxConnIface)
		wantError string
	}{
		{
			name: "OK",
			ctx:  context.Background(),
			fb: &entity.FizzBuzz{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			mockCall: func(fb *entity.FizzBuzz, mock pgxmock.PgxConnIface) {
				mock.ExpectExec("INSERT INTO fizzbuzz").
					WithArgs(fb.Int1, fb.Int2, fb.Limit, fb.Str1, fb.Str2).
					WillReturnResult(pgxmock.NewResult("INSERT", 1))
			},
		},
		{
			name: "Failed: db failed",
			ctx:  context.Background(),
			fb: &entity.FizzBuzz{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			mockCall: func(fb *entity.FizzBuzz, mock pgxmock.PgxConnIface) {
				mock.ExpectExec("INSERT INTO fizzbuzz").
					WithArgs(fb.Int1, fb.Int2, fb.Limit, fb.Str1, fb.Str2).
					WillReturnError(errors.New("failed to exec query"))
			},
			wantError: "failed to exec query",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock, err := pgxmock.NewConn()

			if err != nil {
				t.Errorf("Error on creating mock: %s", err)
			}

			test.mockCall(test.fb, mock)
			repo := repository.NewFizzBuzzRepository(mock)

			err = repo.SaveFizzbuzz(test.ctx, test.fb)
			if err != nil {
				assert.EqualError(t, err, test.wantError)
			} else {
				assert.NoError(t, mock.ExpectationsWereMet())
			}
		})
	}
}

func Test_GetMostFrequentRequest(t *testing.T) {
	type result struct {
		fb    *entity.FizzBuzz
		count int
	}

	tests := []struct {
		name       string
		ctx        context.Context
		mockCall   func(pgxmock.PgxConnIface)
		wantResult result
		wantError  string
	}{
		{
			name: "OK",
			ctx:  context.Background(),
			mockCall: func(mock pgxmock.PgxConnIface) {
				rows := mock.NewRows([]string{"count", "int1", "int2", "limit", "str1", "str2"}).
					AddRow(3, 3, 5, 15, "fizz", "buzz")
				mock.ExpectQuery("SELECT (.+) FROM fizzbuzz").WillReturnRows(rows)
			},
			wantResult: result{
				fb: &entity.FizzBuzz{
					Int1:  3,
					Int2:  5,
					Limit: 15,
					Str1:  "fizz",
					Str2:  "buzz",
				},
				count: 3,
			},
		},
		{
			name: "Failed: db failed",
			ctx:  context.Background(),
			mockCall: func(mock pgxmock.PgxConnIface) {
				mock.ExpectQuery("SELECT (.+) FROM fizzbuzz").WillReturnError(errors.New("failed to exec query"))
			},
			wantError: "failed to exec query",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock, err := pgxmock.NewConn()

			if err != nil {
				t.Errorf("Error on creating mock: %s", err)
			}

			test.mockCall(mock)
			repo := repository.NewFizzBuzzRepository(mock)

			fb, count, err := repo.GetMostFrequentRequest(test.ctx)
			if err != nil {
				assert.EqualError(t, err, test.wantError)
			} else {
				assert.NoError(t, mock.ExpectationsWereMet())

				assert.Equal(t, fb, test.wantResult.fb)
				assert.Equal(t, count, test.wantResult.count)
			}
		})
	}
}
