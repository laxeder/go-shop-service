package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ClientRequest struct {
}

// criando uma postagem
func (c *ClientRequest) Get(url string) (body any, err error) {

	body = make(map[string]interface{})

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	fmt.Println("GET - ", resp.Status, ": ", url)

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var bodyResp = make(map[string]interface{})
	errBodyResp := json.Unmarshal(bodyByte, &bodyResp)
	if errBodyResp == nil {
		body = bodyResp
		return
	}

	var bodyRespSlice = make([]map[string]interface{}, 0)
	errBodyRespSlice := json.Unmarshal(bodyByte, &bodyRespSlice)
	if errBodyRespSlice == nil {
		body = bodyRespSlice
		return
	}

	return

}

// criando uma postagem
func (c *ClientRequest) Post(url string, payload any, user string, password string) (body any, err error) {

	body = make(map[string]interface{})

	payloadByte := make([]byte, 0)

	byteData, err := json.Marshal(payload)
	if err == nil {
		payloadByte = byteData
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadByte))
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(user, password))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var bodyResp = make(map[string]interface{})
	errBodyResp := json.Unmarshal(bodyByte, &bodyResp)
	if errBodyResp == nil {
		body = bodyResp
		return
	}

	var bodyRespSlice = make([]map[string]interface{}, 0)
	errBodyRespSlice := json.Unmarshal(bodyByte, &bodyRespSlice)
	if errBodyRespSlice == nil {
		body = bodyRespSlice
		return
	}

	return
}

// Criando o base de auth
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
