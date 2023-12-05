package gin_metric

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRegisterMetricMonitor(t *testing.T) {
	g := gin.New()
	RegisterMetricMonitor(false, g)
	RegisterApiStatistics(true, g)
	g.GET("/hello", func(c *gin.Context) {
		c.String(0, "hello")
	})
	g.Run(":8090")
	select {}
}
