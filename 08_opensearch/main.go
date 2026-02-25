package main

import (
	"fmt"
	"time"

	logging "github.com/AVVKavvk/opensearch/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var log *logging.Logger

func main() {
	e := echo.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogMethod:   true,
		LogURI:      true,
		LogRemoteIP: true,
		LogLatency:  true,
		// This is mandatory for RequestLogger
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%v [INFO] [%s]: method=%s, status=%d, latency=%v\n",
				v.StartTime.Format("20060102150405.000"),
				v.URI,
				v.Method,
				v.Status,
				v.Latency,
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	count := 0
	e.GET("/", func(ctx echo.Context) error {
		log.Infoxf(&logging.XFields{"count": count, "time": time.Now().String()}, "Log message")
		count++
		return ctx.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))

}

func init() {
	// Initialize logger
	if err := logging.InitializeAllLoggers(); err != nil {
		panic(err)
	}
	logging.RegisterPackageLogger(func() {
		log = logging.DefaultV1Context.GetLogger(
			"opensearch.main",
			logging.LevelDebug,
		)
	})
}
