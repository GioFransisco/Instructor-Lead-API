package middleware

import (
	"log"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
	"github.com/gin-gonic/gin"
)

type LogMiddleware interface {
	LogRequest() gin.HandlerFunc
}

type logMiddleware struct {
	logService common.MyLogger
}

func (l *logMiddleware) LogRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := l.logService.InitializeLogger()
		if err != nil {
			log.Fatal("err:", err.Error())
		}
		t := time.Now()

		ctx.Next()

		logString := utilsmodel.RequestLog{
			AcessTime: t,
			Latency:   time.Since(t),
			ClientIP:  ctx.ClientIP(),
			Method:    ctx.Request.Method,
			Code:      ctx.Writer.Status(),
			Path:      ctx.Request.URL.Path,
			UserAgent: ctx.Request.UserAgent(),
		}

		switch {
		case ctx.Writer.Status() >= 500:
			l.logService.LogWarn(logString)
		case ctx.Writer.Status() >= 400:
			l.logService.LogWarn(logString)
		default:
			l.logService.LogInfo(logString)
		}
	}
}

func NewLogMiddleware(logService common.MyLogger) LogMiddleware {
	return &logMiddleware{logService: logService}
}
