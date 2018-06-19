package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"github.com/gin-gonic/gin/render"
	"log"
)

var (
	err error
	resultType string
)


func recoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Response(c, map[string]interface{}{"code": 0, "message": err}, 200, true)
			}
		}()
		c.Next() // execute all the handlers
	}
}

func BuildRequest(c *gin.Context, obj interface{})  {
	if c.ContentType() == "application/json" {
		if err = c.ShouldBindJSON(&obj); err == nil {
			return
		} else {
			panic(err.Error())
		}
	} else { // c.ContentType() == "application/octet-stream"
		body, _ := ioutil.ReadAll(c.Request.Body)
		if err = msgpack.Unmarshal(body, &obj); err == nil {
			return
		} else {
			panic(err.Error())
		}
	}
}


func Response(c *gin.Context, result interface{}, status int, error bool)  {
	if error == true {
		resultType = "error"
	} else {
		resultType = "result"
	}
	log.Println(resultType, result)

	if c.ContentType() == "application/json" {
		c.JSON(status, map[string]interface{}{resultType: result})

	} else { // "application/octet-stream"
		c.Header("Content-Type", "application/octet-stream")
		c.Render(status, render.MsgPack{map[string]interface{}{resultType: result}})
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(recoveryHandler())

	r.POST("/encryption", Encryption)
	r.POST("/decryption", Decryption)
	r.POST("/multi_encryption", MultiEncryption)

	r.GET("/_liveness", Liveness)
	r.GET("/_readiness", Liveness)

	return r
}


func main(){
	r := setupRouter()

	r.Run(":8090")
}

