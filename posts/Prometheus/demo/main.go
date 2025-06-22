package main

import (
	"math/rand/v2"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 自定义业务指标：业务状态码 Counter
var statusCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_response_status_count",
	},
	[]string{"method", "path", "status"},
)

func initRegistry() *prometheus.Registry {
	// New Registry
	reg := prometheus.NewRegistry()

	// Add Go 编译信息
	reg.MustRegister(collectors.NewBuildInfoCollector())

	// Go runtime metrics
	reg.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")})))

	// 注册自定义的业务指标
	reg.MustRegister(statusCounter)

	return reg
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		// Mock 业务逻辑
		status := 0
		if rand.IntN(10)%3 == 0 {
			status = 1
		}

		// 记录业务指标
		statusCounter.WithLabelValues(
			ctx.Request.Method,
			ctx.Request.URL.Path,
			strconv.Itoa(status),
		).Inc()

		ctx.JSON(200, gin.H{
			"status": status,
			"msg":    "pong",
		})
	})

	// 对外提供 /metrics 接口，支持 prometheus 采集
	reg := initRegistry()
	r.GET("/metrics", gin.WrapH(
		promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				Registry: reg,
			},
		)))

	_ = r.Run("127.0.0.1:5567")
}
