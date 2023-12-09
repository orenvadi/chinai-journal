package storage

import "errors"

var (
	ErrQRNotFound = errors.New("qr not found")
	ErrQRExpired  = errors.New("qr is expired")
)
