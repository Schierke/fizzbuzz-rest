package service

import (
	"context"
	"fizzbuzz/internal/domain/entity"
)

type FizzBuzzStorage interface {
	SaveFizzbuzz(context.Context, *entity.FizzBuzz) error
	GetMostFrequentRequest(context.Context) (*entity.FizzBuzz, int, error)
}

type fbServiceImpl struct {
	db FizzBuzzStorage
}

func NewFizzBuzzService(store FizzBuzzStorage) *fbServiceImpl {
	return &fbServiceImpl{
		db: store,
	}
}

func (s *fbServiceImpl) CalculateFizzBuzz(ctx context.Context, entity *entity.FizzBuzz) ([]string, error) {
	res := entity.CalculateFizzBuzz()
	if err := s.db.SaveFizzbuzz(ctx, entity); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *fbServiceImpl) GetMostFrequentRequest(ctx context.Context) (*entity.FizzBuzz, int, error) {
	fb, count, err := s.db.GetMostFrequentRequest(ctx)

	if err != nil {
		return nil, 0, err
	}

	return fb, count, err
}
