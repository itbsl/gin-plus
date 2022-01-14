package logger

import (
	"gin-plus/config"
	"gin-plus/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func Init(config *config.LogConfig) (err error) {
	logSavePath := config.SavePath + config.Filename + config.Ext
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logSavePath,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  true,
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(config.Level))
	if err != nil {
		return
	}

	var core zapcore.Core
	if global.Config.AppConfig.Mode == "debug" {
		//开发模式下，日志同时输出到控制台和写入文件
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
		//如果想开发模式下仅输出到控制台，就不必使用NewTee方法
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}
	logger := zap.New(core, zap.AddCaller())
	//替换zap库中全局的logger,因为默认的全局logger什么都不干
	//替换完成后，我们在其他地方就可以通过zap.L()或者zap.S()获取全局logger去记录日志了
	//zap.L()和zap.S()方法可以获取全局的Logger和SugaredLogger
	zap.ReplaceGlobals(logger)
	return
}
