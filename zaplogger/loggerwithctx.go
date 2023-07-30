package zaplogger

import (
	"context"
	"os"
	"strings"

	"github.com/kubescape/go-logger/helpers"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ helpers.ILogger = (*ZapLoggerWithCtx)(nil)

type ZapLoggerWithCtx struct {
	zapL *otelzap.LoggerWithCtx
	cfg  zap.Config
}

func (zl *ZapLoggerWithCtx) GetLevel() string                      { return zl.cfg.Level.Level().String() }
func (zl *ZapLoggerWithCtx) SetWriter(w *os.File)                  {}
func (zl *ZapLoggerWithCtx) GetWriter() *os.File                   { return nil }
func (zl *ZapLoggerWithCtx) Ctx(_ context.Context) helpers.ILogger { return zl }
func (zl *ZapLoggerWithCtx) LoggerName() string                    { return LoggerName }
func (zl *ZapLoggerWithCtx) SetLevel(level string) error {
	l := zapcore.Level(1)
	err := l.Set(level)
	if err == nil {
		zl.cfg.Level.SetLevel(l)
	}
	return err
}
func (zl *ZapLoggerWithCtx) Fatal(msg string, details ...helpers.IDetails) {
	zl.zapL.Fatal(strings.ToValidUTF8(msg, helpers.InvalidUtf8ReplacementString), detailsToZapFields(details)...)
}

func (zl *ZapLoggerWithCtx) Error(msg string, details ...helpers.IDetails) {
	zl.zapL.Error(strings.ToValidUTF8(msg, helpers.InvalidUtf8ReplacementString), detailsToZapFields(details)...)
}

func (zl *ZapLoggerWithCtx) Warning(msg string, details ...helpers.IDetails) {
	zl.zapL.Warn(strings.ToValidUTF8(msg, helpers.InvalidUtf8ReplacementString), detailsToZapFields(details)...)
}

func (zl *ZapLoggerWithCtx) Success(msg string, details ...helpers.IDetails) {
	// calling ZapLogger() to get the underlying logger and not attach the log to the span
	zl.zapL.ZapLogger().Info(strings.ToValidUTF8(msg, helpers.InvalidUtf8ReplacementString), detailsToZapFields(details)...)
}

func (zl *ZapLoggerWithCtx) Info(msg string, details ...helpers.IDetails) {
	// calling ZapLogger() to get the underlying logger and not attach the log to the span
	zl.zapL.ZapLogger().Info(strings.ToValidUTF8(msg, helpers.InvalidUtf8ReplacementString), detailsToZapFields(details)...)
}

func (zl *ZapLoggerWithCtx) Debug(msg string, details ...helpers.IDetails) {
	// calling ZapLogger() to get the underlying logger and not attach the log to the span
	zl.zapL.ZapLogger().Debug(strings.ToValidUTF8(msg, helpers.InvalidUtf8ReplacementString), detailsToZapFields(details)...)
}
