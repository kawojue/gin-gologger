package gologger

import (
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func init() {
	color.NoColor = false
}

// Logger returns a Gin middleware that logs information about each incoming request.
//
// It takes a mode parameter, which can be "dev" for development mode or "release" for release mode.
func Logger(mode string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		remoteAddr := ctx.ClientIP()

		var statusColor *color.Color

		defer func() {
			elapsedTime := time.Since(startTime)
			statusCode := ctx.Writer.Status()

			switch {
			case statusCode < 200:
				statusColor = color.New(color.FgWhite)
			case statusCode < 300:
				statusColor = color.New(color.FgGreen)
			case statusCode < 400:
				statusColor = color.New(color.FgYellow)
			case statusCode < 500:
				statusColor = color.New(color.FgRed)
			default:
				statusColor = color.New(color.FgRed)
			}

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
		}()

		ctx.Next()
	}
}
