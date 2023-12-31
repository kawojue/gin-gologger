package gologger

import (
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// It takes "dev" on development mode
// It takes "release" on release mode
func Logger(mode string) gin.HandlerFunc {
	color.NoColor = false

	return func(ctx *gin.Context) {
		var statusColor *color.Color

		startTime := time.Now()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		remoteAddr := ctx.ClientIP()
		statusCode := ctx.Writer.Status()

		switch {
		case statusCode < 200:
			statusColor = color.New(color.FgWhite)
		case statusCode < 300:
			statusColor = color.New(color.FgGreen)
		case statusCode < 400:
			statusColor = color.New(color.FgYellow)
		case statusCode < 500:
			statusColor = color.New(color.FgMagenta)
		default:
			statusColor = color.New(color.FgRed)
		}

		ctx.Next()

		elapsedTime := time.Since(startTime)

		if mode == "release" {
			statusColor.Printf(
				"\n[%s]\t|\t%d\t|\t%s\t|\t%v\t|\t%s\n",
				method, statusCode, path, elapsedTime, remoteAddr,
			)
		} else {
			statusColor.Printf(
				"\n[%s]\t|\t%d\t|\t%s\n",
				method, statusCode, path,
			)
		}
	}
}
