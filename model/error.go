package model

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")

	ErrOpendJobSessionNotFound = errors.New("opend job session not found")
)
