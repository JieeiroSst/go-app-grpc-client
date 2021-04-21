package log 

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
    config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    config.EncoderConfig.TimeKey = "timestamp"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    logger, _ := config.Build()
	writeFile(logger)
    return logger
}

func writeFile(byteSlice interface {}){
	file,_:=os.OpenFile("./data.log",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
	b, err := json.Marshal(byteSlice)
    if err != nil {
        panic(err)
    }
    _, _ = file.Write(b)
}

