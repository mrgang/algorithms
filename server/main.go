package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	route.NoRoute(func(context *gin.Context) {
		context.String(http.StatusOK, ""+context.Request.RequestURI)
	})
	fmt.Println(route.Run(":80"))
}
