package main

import (
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
	"log"
)

type QxxService interface {
	Getcompany(string) (interface{}, error)
}

type qxxService struct{}

var ErrEmpty = errors.New("empty string")

func (qxxService) Getcompany(s string) (interface{}, error) {

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
		log.Fatalln(err)
	}

	js, err := simplejson.NewJson(body)
	if err != nil {
		log.Fatalln(err)
	}
	data := js.Get("data").Get("Result")

	//for i, v := range data.MustArray(){
	//	fmt.Println(i)
	//	fmt.Println(v)
	//}

	return data, nil
}
