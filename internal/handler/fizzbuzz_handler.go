package handler

import (
	"context"
	"encoding/json"
	"fizzbuzz/internal/domain/entity"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type FizzBuzzService interface {
	CalculateFizzBuzz(context.Context, *entity.FizzBuzz) ([]string, error)
	GetMostFrequentRequest(context.Context) (*entity.FizzBuzz, int, error)
}
type handler struct {
	service FizzBuzzService
}

func NewFizzBuzzHandler(service FizzBuzzService, r *chi.Mux) {
	handler := &handler{
		service: service,
	}
	r.Route("/fizzbuzz/v1", func(r chi.Router) {
		r.Get("/fizzbuzz", handler.CreateFizzBuzzString)
		r.Get("/stats", handler.GetStats)
	})

}

func (h *handler) CreateFizzBuzzString(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error
	var int1, int2, limit int
	var str1, str2 string
	if int1, err = strconv.Atoi(r.URL.Query().Get("int1")); err != nil {
		log.Error().Err(err).Msg("invalid request: can't extract int1")
		http.Error(w, errInvalidFizzBuzzParams, http.StatusBadRequest)
		return
	}

	if int2, err = strconv.Atoi(r.URL.Query().Get("int2")); err != nil {
		log.Error().Err(err).Msg("invalid request: can't extract int2")
		http.Error(w, errInvalidFizzBuzzParams, http.StatusBadRequest)
		return
	}

	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil {
		log.Error().Err(err).Msg("invalid request: can't extract limit")
		http.Error(w, errInvalidFizzBuzzParams, http.StatusBadRequest)
		return
	}

	str1 = r.URL.Query().Get("str1")
	str2 = r.URL.Query().Get("str2")

	fizzbuzz, err := entity.NewFizzBuzzEntity(int1, int2, limit, str1, str2)
	if err != nil {
		log.Error().Err(err).Msg("invalid request: can't extract limit")
		http.Error(w, errFizzBuzzNotValidate, http.StatusBadRequest)
	}

	ret, err := h.service.CalculateFizzBuzz(ctx, fizzbuzz)
	if err != nil {
		log.Error().Msg(err.Error())
		select {
		case <-ctx.Done():
			http.Error(w, timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		return

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) GetStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fb, count, err := h.service.GetMostFrequentRequest(ctx)

	if err != nil {
		log.Error().Msg(err.Error())
		select {
		case <-ctx.Done():
			http.Error(w, timeout, http.StatusGatewayTimeout)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ret := struct {
		Request *entity.FizzBuzz `json:"request"`
		Count   int              `json:"count"`
	}{
		Request: fb,
		Count:   count,
	}
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		log.Error().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
