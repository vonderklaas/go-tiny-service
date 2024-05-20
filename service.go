package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}

	fact := &CatFact{}
	defer resp.Body.Close()

	// Decode
	decoder := json.NewDecoder(resp.Body)
	e := decoder.Decode(fact)
	if e != nil {
		return nil, err
	}

	return fact, nil
}
