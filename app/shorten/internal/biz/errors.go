package biz

import "errors"

var (
	ErrParamInvalid      = errors.New("invalid param")
	ErrNotFoundFromDB    = errors.New("data not found from db")
	ErrNotFoundFromRedis = errors.New("data not found from redis")
)
