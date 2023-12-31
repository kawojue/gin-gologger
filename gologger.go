package gologger

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Writer.Status(), ctx.Request.URL.Path, ctx.Request.Method)
		ctx.Next()
	}
}
