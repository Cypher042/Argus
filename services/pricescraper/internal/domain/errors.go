package domain

import "errors"

var (
	ErrInvalidPrice  = errors.New("invalid price data")
	ErrSourceTimeout = errors.New("data source timeout")
)
