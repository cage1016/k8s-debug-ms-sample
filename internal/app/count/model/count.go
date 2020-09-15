package model

import "context"

type Count struct {
	Value int64 `json:"value" db:"value"`
}

type CountRespository interface {
	Add(context.Context) error
	Get(context.Context) (int64, error)
}
