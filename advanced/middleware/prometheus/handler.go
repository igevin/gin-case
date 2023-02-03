package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
)

func metrics(r gin.IRouter) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func promMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Counter
		AccessCounter.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Add(1)

		// Histogram
		HttpDurationsHistogram.With(prometheus.Labels{
			"path": c.Request.URL.Path,
		}).Observe(float64(rand.Intn(30)))

		// Summary
		HttpDurations.With(prometheus.Labels{
			"path": c.Request.URL.Path,
		}).Observe(float64(rand.Intn(30)))
	}
}
