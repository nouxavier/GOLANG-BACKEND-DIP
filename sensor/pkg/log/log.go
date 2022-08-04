//package log - we chose based on this article https://blog.logrocket.com/5-structured-logging-packages-for-go/
package log

import (
	"log"
	"time"

	"sensor/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLog(config *config.ConfigLog) *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger.WithOptions()
}
