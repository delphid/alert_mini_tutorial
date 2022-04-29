package main

import (
        "fmt"
        "strconv"

        "go.uber.org/zap"
        "go.uber.org/zap/zapcore"
        "github.com/gin-gonic/gin"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

type AlertManagerAlert struct {
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"` // must have description/runbook_url/summary
	EndsAt       string            `json:"endsAt"`
	Status       string            `json:"status"` // resolved|firing
	StartsAt     string            `json:"startsAt"`
	Fingerprint  string            `json:"fingerprint"`
	GeneratorURL string            `json:"generatorURL"`
}

type AlertManagerReceiverMsg struct {
	Alerts            []AlertManagerAlert `json:"alerts"` // alertmanager 中配置， 每次 5 个.
	Status            string              `json:"status"`
	Version           string              `json:"version"`
	GroupKey          string              `json:"groupKey"`
	Receiver          string              `json:"receiver"`
	ExternalURL       string              `json:"externalURL"`
	TruncatedAlerts   int                 `json:"truncatedAlerts"`
	GroupLabels       map[string]string   `json:"groupLabels"`
	CommonLabels      map[string]string   `json:"commonLabels"`
	CommonAnnotations map[string]string   `json:"commonAnnotations"`
}

func NewLogger() (*zap.Logger, error) {
  cfg := zap.NewProductionConfig()
  cfg.OutputPaths = []string{
    "./log/log",
  }
  encoderConfig := zap.NewProductionEncoderConfig()
  encoderConfig.TimeKey = "time"
  encoderConfig.LevelKey = ""
  encoderConfig.CallerKey = ""
  encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.000")
  cfg.EncoderConfig =encoderConfig
  return cfg.Build()
}

var (
    log, _ = NewLogger()
)

func alert(c *gin.Context) {
    var req AlertManagerReceiverMsg
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

func modify(c *gin.Context) {
    value, _ := strconv.ParseFloat(c.Param("value"), 64)
    valueGauge.Set(value)
    log.Info("change gauge", zap.Float64("value", value))
    fmt.Println("set value to: ", value)
}

var (
        valueGauge = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "my_gauge",
                Help: "this is my gauge",
        })
)

func main() {
    router := gin.Default()
    router.GET("/metrics", gin.WrapH(promhttp.Handler()))
    router.GET("/modify/:value", modify)
    router.POST("/alert", alert)
    router.Run("0.0.0.0:2112")
}
