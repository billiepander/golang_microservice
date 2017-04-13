package main

import (
	"errors"
	"io/ioutil"
	"fmt"
	"net/http"
)

type QxxService interface {
	Getcompany(string) (string, error)
}

type qxxService struct{}

var ErrEmpty = errors.New("empty string")

func (qxxService) Getcompany(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	url := fmt.Sprintf("http://open.qichacha.com/link/search?key=%s&province=&token=712c2f98c26c284df43c9c888ef67d52", s)
	resp, err := http.Get(url)
	if err != nil {
		return "", ErrEmpty
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", ErrEmpty
	}

	return string(body), nil
}
