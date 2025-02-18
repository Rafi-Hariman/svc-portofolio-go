package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const ENV = "ENV"
const USER = "USER"
const LOCAL = "local"
const PASSWORD = "PASSWORD"
const CLOUD_API = "CLOUD_API"
const LOGIN_URL = "/v1/login"
const CONTENT_TYPE = "Content-Type"
const APPLICATION_JSON = "application/json"

type result struct {
	Authorization struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	} `json:"authorization"`
}

func getBearer() string {
	type login struct {
		KodeMember string `json:"kd_member"`
		Password   string `json:"password"`
	}
	data, err := json.Marshal(login{
		KodeMember: os.Getenv(USER),
		Password:   os.Getenv(PASSWORD),
	})
	if err != nil {
		return ""
	}
	result := result{}
	err = HttpPost(map[string]interface{}{
		"url":  os.Getenv(CLOUD_API) + LOGIN_URL,
		"data": data,
	}, &result)
	if err != nil {
		return ""
	}
	return result.Authorization.AccessToken
}

func HttpGet(param map[string]interface{}, result any) (err error) {
	request, err := http.NewRequest(http.MethodGet, param["url"].(string), nil)

	if err != nil {
		return
	}

	query := request.URL.Query()

	for i, x := range param["query"].(map[string]string) {
		query.Add(i, x)
	}

	// FOR LOCAL ONLY
	if os.Getenv(ENV) == LOCAL {
		request.Header.Set("Authorization", getBearer())
	}

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return
	}

	if response.StatusCode > 300 {
		err = fmt.Errorf(string(body))
		return
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return
	}

	return
}

func HttpPost(param map[string]interface{}, result any) (err error) {
	url := param["url"].(string)
	data := param["data"].([]byte)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	if err != nil {
		return
	}

	request.Header.Set(CONTENT_TYPE, APPLICATION_JSON)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return
	}

	if response.StatusCode > 300 {
		err = fmt.Errorf(string(body))
		return
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return
	}

	return
}

func HttpPut(param map[string]interface{}, result any) (err error) {
	url := param["url"].(string)
	data := param["data"].([]byte)

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))

	if err != nil {
		return
	}

	request.Header.Set(CONTENT_TYPE, APPLICATION_JSON)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return
	}

	if response.StatusCode > 300 {
		err = fmt.Errorf(string(body))
		return
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return
	}

	return
}
