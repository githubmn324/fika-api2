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
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message":      "hello I'm Cloud Run API 2!",
			"requstHeader": string(requestDump),
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
