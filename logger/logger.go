package logger

import (
	"example/fundemo01/web-app/settings"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)



func Init(cfg *settings.LogConfig) (err error){
	//filename := viper.GetString("log.filename")
	//maxSize := viper.GetInt("log.max_size")
	//maxBackup := viper.GetInt("log.maxBackup")
	//maxAge := viper.GetInt("log.maxAge")
	//level := viper.GetString("log.Level")
	logmode := viper.GetString("log.mode")
	var level string
	filename := cfg.Filename
	maxSize := cfg.MaxSize
	maxBackup := cfg.MaxBackups
	maxAge := cfg.MaxAge
	level = cfg.Level
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge)
	var lev = new(zapcore.Level)
	err = lev.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	var core zapcore.Core
	switch logmode {
	case "dev":
		devEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(devEncoder,writeSyncer,lev),
			zapcore.NewCore(devEncoder,zapcore.Lock(os.Stdout),zap.DebugLevel),
		)
	case "product":
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.TimeKey = "time"
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		prodEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewTee(
			zapcore.NewCore(prodEncoder,writeSyncer,lev),
		)
	}
	log := zap.New(core,zap.AddCaller())
	zap.ReplaceGlobals(log)
	return
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer{
	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,
		MaxSize: maxSize,
		MaxBackups: maxBackup,
		MaxAge: maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func GinLogger() gin.HandlerFunc{
	return func(c *gin.Context){
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
