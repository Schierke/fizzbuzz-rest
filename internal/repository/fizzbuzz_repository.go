package repository

import (
	"context"
	"fizzbuzz/internal/domain/entity"
	"time"
)

const (
	insertFizzBuzzQuery         = `INSERT INTO fizzbuzz (int1, int2, limit_number, str1, str2, created_at) VALUES($1, $2, $3, $4, $5, NOW());`
	getMostFrequentRequestQuery = `
	SELECT COUNT(*) as count, int1, int2, limit_number, str1, str2 
	FROM fizzbuzz
	GROUP BY int1, int2, limit_number, str1, str2
	ORDER BY count DESC
	LIMIT 1;
	`
)

type FizzBuzzRepository struct {
	pool    PgxIface
	timeout time.Duration
}

func NewFizzBuzzRepository(c PgxIface) *FizzBuzzRepository {
	return &FizzBuzzRepository{
		pool:    c,
		timeout: time.Duration(120) * time.Second,
	}
}

func (fbRepo *FizzBuzzRepository) SaveFizzbuzz(ctx context.Context, entity *entity.FizzBuzz) error {
	ctx, cancel := context.WithTimeout(ctx, fbRepo.timeout)
	defer cancel()
	_, err := fbRepo.pool.Exec(ctx, insertFizzBuzzQuery, entity.Int1, entity.Int2, entity.Limit, entity.Str1, entity.Str2)

	if err != nil {
		return err
	}
	return nil
}

func (fbRepo *FizzBuzzRepository) GetMostFrequentRequest(ctx context.Context) (*entity.FizzBuzz, int, error) {
	ctx, cancel := context.WithTimeout(ctx, fbRepo.timeout)
	defer cancel()
	var count, int1, int2, limit int
	var str1, str2 string
	err := fbRepo.pool.QueryRow(ctx, getMostFrequentRequestQuery).Scan(&count, &int1, &int2, &limit, &str1, &str2)

	if err != nil {
		return nil, 0, err
	}

	return &entity.FizzBuzz{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}, count, nil
}
