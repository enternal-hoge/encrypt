package main

import (
	"encoding/json"
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

type RequestResult struct {
	Result  interface{}						`msgpack:"result" json:"result"`
	Error   map[string]interface{}			`msgpack:"error" json:"error"`
}


var router = setupRouter()

func requestJson(url string, data string) (*RequestResult, *httptest.ResponseRecorder){
	w := httptest.NewRecorder()

	var jsonStr = []byte(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	requestResult := new(RequestResult)

	_ = json.Unmarshal(body, &requestResult)

	return requestResult, w
}

func requestMsgPack(url string, data interface{}) (*RequestResult, *httptest.ResponseRecorder){
	w := httptest.NewRecorder()

	msgpackB, _ := msgpack.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(msgpackB))
	req.Header.Add("Content-Type", "application/octet-stream")
	router.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	requestResult := new(RequestResult)

	_ = msgpack.Unmarshal(body, &requestResult)

	return requestResult, w
}

func TestEncryptionByJson(t *testing.T) {
	data := `{"key":"asdsdsdsdsdsdsdsdsdsdsds", "text": "aaaeee"}`
	result, _:= requestJson("/encryption", data)

	assert.Equal(t, "000000000000000020db5f18dcf8746e", result.Result)
}

func TestEncryptionByMsgPack(t *testing.T) {
	data := map[string]interface{}{"key":"asdsdsdsdsdsdsdsdsdsdsds", "text": "aaaeee"}
	result, _ := requestMsgPack("/encryption", data)

	assert.Equal(t, "000000000000000020db5f18dcf8746e", result.Result)
}


func TestEncryptionValidation(t *testing.T) {
	data := map[string]interface{}{"key":"asdsds", "text": "aaaeee"}
	result, _ := requestMsgPack("/encryption", data)

	expected :=  map[string]interface{}{"code":int8(0), "message": "Key length should be 24 letters."}
	assert.Equal(t, expected, result.Error)
}

func TestDecryption(t *testing.T) {
	data := map[string]interface{}{"key":"asdsdsdsdsdsdsdsdsdsdsds", "text": "000000000000000020db5f18dcf8746e"}
	result, _ := requestMsgPack("/decryption", data)

	assert.Equal(t, "aaaeee", result.Result)
}

func TestMultiEncryption(t *testing.T) {
	data :=map[string]interface{}{"key":"asdsdsdsdsdsdsdsdsdsdsds", "texts": []string{"aaaeee", "bbbeee", "ccceee"}}
	result, _ := requestMsgPack("/multi_encryption", data)

	var expected = []interface{}{"000000000000000020db5f18dcf8746e","00000000000000007e689dd9d9cf50a9","0000000000000000fb0f724999926ab4"}
	assert.Equal(t, expected, result.Result)
}