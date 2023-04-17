package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Printf("starting HTTP")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// context から取得した authToken を受け取る
		auth0Token, ok := c.Value("auth0-token").(string)
		if !ok {
			// エラー処理
		}
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message":      "hello I'm Cloud Run API 2!",
			"requstHeader": string(requestDump),
			"auth0-token":  auth0Token,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
