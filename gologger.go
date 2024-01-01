package gologger

import (
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func init() {
	color.NoColor = false
}

// It takes "dev" on development mode.
// It takes "release" on release mode.
func Logger(mode string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		remoteAddr := ctx.ClientIP()

		var statusColor *color.Color

		switch {
		case ctx.Writer.Status() < 200:
			statusColor = color.New(color.FgWhite)
		case ctx.Writer.Status() < 300:
			statusColor = color.New(color.FgGreen)
		case ctx.Writer.Status() < 400:
			statusColor = color.New(color.FgYellow)
		case ctx.Writer.Status() < 500:
			statusColor = color.New(color.FgRed)
		default:
			statusColor = color.New(color.FgRed)
		}

		defer func() {
			elapsedTime := time.Since(startTime)

			if mode == "release" {
				statusColor.Printf(
					"\n[%s]\t|\t%d\t|\t%s\t|\t%v\t|\t%s\n",
					method, ctx.Writer.Status(), path, elapsedTime, remoteAddr,
				)
			} else {
				statusColor.Printf(
					"\n[%s]\t|\t%d\t|\t%s\n",
					method, ctx.Writer.Status(), path,
				)
			}
		}()

		ctx.Next()
	}
}
