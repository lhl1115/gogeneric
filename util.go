package gogeneric

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
)

// HttpGetJson get json data from api
// apiURL: api url
// headers: http headers
// return: json data
// T: json data struct
func HttpGetJson[T any](apiURL string, headers map[string]string) (*T, error) {

	var response T
	// check Value Type == reflect.Struct
	typeT := reflect.TypeOf(response)
	if typeT == nil {
		return nil, errors.New("response type is nil")
	}
	if typeT.Kind() != reflect.Struct {
		return nil, errors.New("response type is not struct")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// HttpPostJson post json data to api
// apiURL: api url
// request: request data struct (json) can be nil
// headers: http headers
// return: json data
// T: json data struct
func HttpPostJson[T any](apiURL string, request any, headers map[string]string) (response T, err error) {

	// check Value Type == reflect.Struct
	typeT := reflect.TypeOf(response)
	if typeT == nil {
		return response, errors.New("response type is nil")
	}
	if typeT.Kind() != reflect.Struct {
		return response, errors.New("response type is not struct")
	}

	client := &http.Client{}
	var req *http.Request
	if request != nil {
		data, errJson := json.Marshal(request)
		if errJson != nil {
			return response, errJson
		}
		req, err = http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest("POST", apiURL, nil)
	}

	if err != nil {
		return response, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return
}
