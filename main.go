package main 

import (
	"errors"
	"github.com/sony/gobreaker"
)

func mockService() (string, error) {
	if rand.Intn(100) > 50 {
		return "sucess", nil 
	}

	return "", errors.New("Error trying to process request")
}