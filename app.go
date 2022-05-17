package main

import (
	"fmt"
	"strconv"
  "time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
  alertmanagerModel "github.com/prometheus/alertmanager/template"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./log/log",
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = ""
	encoderConfig.CallerKey = ""
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	cfg.EncoderConfig = encoderConfig
	return cfg.Build()
}

var (
	log, _ = NewLogger()
)

var (
	gaugeA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gauge_a",
		Help: "this is gauge_a",
	}, []string{"proc"})
	gaugeB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "gauge_b",
		Help: "this is gauge_b",
	})
)

func alert(c *gin.Context) {
	var req alertmanagerModel.Data
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Printf("req: %+v\n", req)
	alertsLen := len(req.Alerts)
	alertsStatus := make([]string, alertsLen)
	for i, _ := range alertsStatus {
		alertsStatus[i] = req.Alerts[i].Status
	}
	log.Info("receive alert", zap.String("status", req.Status), zap.Int("len", alertsLen), zap.Any("alerts_status", alertsStatus), zap.Any("full_massage", req))
}

func setGaugeA(c *gin.Context) {
	proc := c.Param("proc")
	value, _ := strconv.ParseFloat(c.Param("value"), 64)
	gaugeA.WithLabelValues(proc).Set(value)
	log.Info("set gauge a", zap.String("proc", proc), zap.Float64("value", value))
}

func setGaugeB(c *gin.Context) {
	value, _ := strconv.ParseFloat(c.Param("value"), 64)
	gaugeB.Set(value)
	log.Info("set gauge b", zap.Float64("value", value))
}

func main() {
	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/set_gauge_a/proc/:proc/to_value/:value", setGaugeA)
	router.GET("/set_gauge_b/to_value/:value", setGaugeB)
	router.POST("/alert", alert)
	router.Run("0.0.0.0:80")
}
