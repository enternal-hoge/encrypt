package main

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"bytes"
)


func Liveness(c *gin.Context)  {
	c.String(200, "OK!")
}


func Encryption(c *gin.Context) {
	var (
		data SoloData
		result []byte
	)

	BuildRequest(c, &data)
	data.validate()

	result, err = encrypt([]byte(data.Text), []byte(data.Key))
	if err != nil {
		panic(err.Error())
	}

	Response(c, hex.EncodeToString(result), 200, false)
}

func Decryption(c *gin.Context) {
	var (
		data SoloData
		result []byte
	)

	BuildRequest(c, &data)
	data.validate()

	var ciphertext, _ = hex.DecodeString(data.Text)

	result, err = decrypt(ciphertext, []byte(data.Key))
	if err != nil {
		panic(err.Error())
	}

	result = bytes.Trim(result, "\x00") // delete nulls
	Response(c, string(result), 200, false)
}

func MultiEncryption(c *gin.Context) {
	var (
		data MultiData
		result []string
	)

	BuildRequest(c, &data)
	data.validate()

	for _,v := range data.Texts{
		res, err := encrypt([]byte(v), []byte(data.Key))
		if err == nil {
			result = append(result, hex.EncodeToString(res))
		} else {
			panic(err.Error())
		}
	}

	Response(c, result, 200, false)
}
