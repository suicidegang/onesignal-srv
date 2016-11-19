package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/jmcvetta/napping"
)

func RequestPost(endpoint string, body map[string]interface{}) (map[string]interface{}, error) {
	var response map[string]interface{}

	session := napping.Session{
		Header: requestHeaders(),
	}

	body["app_id"] = AppID
	log.Printf("Sending %+v", body)
	res, err := session.Post(ONESIGNAL_ENDPOINT+endpoint, &body, &response, nil)

	if err != nil {
		return response, err
	}

	if res.Status() != 200 {
		return response, errors.New(res.RawText())
	}

	return response, nil
}

func requestHeaders() *http.Header {
	headers := &http.Header{}
	headers.Add("Content-Type", "application/json; charset=utf-8")
	headers.Add("Authorization", "Basic "+ApiKey)
	return headers
}
