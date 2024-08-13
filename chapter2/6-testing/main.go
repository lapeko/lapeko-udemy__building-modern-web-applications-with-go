package main

import (
	"errors"
)

func Divide(dividend, divider float32) (float32, error) {
	if divider == 0 {
		return 0.0, errors.New("divider can not be null")
	}
	return dividend / divider, nil
}
