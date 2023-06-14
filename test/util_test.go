package test

import (
	"testing"
)

type Req struct {
	ID int `json:"id"`
}

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestGetJson(t *testing.T) {
	//res, err := gogeneric.HttpGetJson[Response]("http://localhost:8080/api/v1/ping", nil)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(res)
}

func TestPostJson(t *testing.T) {
	//res, err := gogeneric.HttpPostJson[Response]("http://localhost:8080/api/v1/pong", Req{ID: 1}, nil)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(res)
}
