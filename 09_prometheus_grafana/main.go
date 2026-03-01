package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	AppServiceConnectionStatusMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "app",
		Subsystem: "test",
		Name:      "connection_status",
		Help:      " App service connection status (1 for connected, 0 for disconnected)",
	})
)

func main() {
	e := echo.New()

	// Standard Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from the Go Pod!")
	})

	// Metrics endpoint for Prometheus to scrape
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Custom Route
	e.GET("/status", func(c echo.Context) error {
		AppServiceConnectionStatusMetric.Set(1)
		return c.String(http.StatusOK, "App service is connected")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
