package main

import (
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

var ErrEmpty = errors.New("empty string")

func (stringService stringService) Uppercase(s string) (string, error) {
	if s == "" {
		// return "", errors.New("Empty String")
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService stringService) Count(s string) int {
	return len(s)
}
