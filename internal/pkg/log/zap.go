package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zap struct {
	*zap.SugaredLogger
}

const filename = "app.log"

func newZap() (Logger, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(logFile),
		zap.InfoLevel,
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).Sugar()

	return &Zap{
		logger,
	}, nil
}

func (z *Zap) Debug(format string, args ...any) {
	z.Debugw(format, args...)
}

func (z *Zap) Info(format string, args ...any) {
	z.Infow(format, args...)
}

func (z *Zap) Error(format string, args ...any) {
	z.Errorw(format, args...)
}

func (z *Zap) Warn(formar string, args ...any) {
	z.Warnw(formar, args...)
}

func (z *Zap) Fatal(format string, args ...any) {
	z.Fatalw(format, args...)
}

func (z *Zap) Panic(format string, args ...any) {
	z.Panicw(format, args...)
}

func (z *Zap) Close() error {
	err := z.Sync()
	if err != nil {
		return err
	}
	return nil
}
