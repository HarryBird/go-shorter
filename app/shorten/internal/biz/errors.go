package biz

import "errors"

var (
	ErrNotFoundFromDB    = errors.New("data not found from db")
	ErrNotFoundFromRedis = errors.New("data not found from redis")
)
