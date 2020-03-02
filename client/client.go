package client

import (
	"awesomeProject1/server/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func SaveHumanByPOSTRequest(humanJSONString string) bool {
	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewReader([]byte(humanJSONString)))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	return resp.StatusCode == http.StatusCreated
}
func ReadHumanByGetRequest(humanName string) (*model.Human, bool) {
	queryString := "http://localhost:8080?first_name=" + humanName
	req, err := http.NewRequest("GET", queryString, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	defer resp.Body.Close()
	h := &model.Human{}
	err = json.Unmarshal(body, h)
	if err != nil {
		log.Fatal("Error Unmarshaling body. ")
	}
	return h, true
}
func UpdateByPUTRequest(humanJSONString string) int {
	req, err := http.NewRequest("PUT", "http://localhost:8080/", bytes.NewReader([]byte(humanJSONString)))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	return resp.StatusCode
}
func DeleteHumanByRequest(humanName string) int {
	queryString := "http://localhost:8080?first_name=" + humanName
	req, err := http.NewRequest("DELETE", queryString, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	return resp.StatusCode
}
