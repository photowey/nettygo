/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	"os"
	"path/filepath"

	"github.com/photowey/nettygo/interal/utilz/expression"
	"github.com/photowey/nettygo/interal/utilz/filepathz"
	"github.com/photowey/perrors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultPath     = "logs"
	defaultFileName = "nettygo.log"
	defaultLevel    = "debug"
	emptyString     = ""
)

var (
	_       Logger = (*loggerx)(nil)
	_logger Logger
)

type Logger interface {
	Debug(args ...any)
	Debugf(template string, args ...any)
	Info(args ...any)
	Infof(template string, args ...any)
	Infow(template string, args ...any)
	Warn(args ...any)
	Warnf(template string, args ...any)
	Warnw(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
	Errorw(template string, args ...any)
	Fatal(args ...any)
	Fatalf(template string, args ...any)
	Fatalw(template string, args ...any)
	Panic(args ...any)
	Panicf(template string, args ...any)
	Panicw(template string, args ...any)
	Sync()
}

type loggerx struct {
	logger *zap.SugaredLogger
}

func (log loggerx) Debug(args ...any) {
	log.logger.Debug(args...)
}

func (log loggerx) Debugf(template string, args ...any) {
	log.logger.Debugf(template, args...)
}

func (log loggerx) Info(args ...any) {
	log.logger.Info(args...)
}

func (log loggerx) Infof(template string, args ...any) {
	log.logger.Infof(template, args...)
}

func (log loggerx) Infow(template string, keysAndValues ...any) {
	log.logger.Infow(template, keysAndValues...)
}

func (log loggerx) Warn(args ...any) {
	log.logger.Warn(args...)
}

func (log loggerx) Warnf(template string, args ...any) {
	log.logger.Warnf(template, args...)
}

func (log loggerx) Warnw(template string, keysAndValues ...any) {
	log.logger.Warnw(template, keysAndValues...)
}

func (log loggerx) Error(args ...any) {
	log.logger.Error(args...)
}

func (log loggerx) Errorf(template string, args ...any) {
	log.logger.Errorf(template, args...)
}

func (log loggerx) Errorw(template string, keysAndValues ...any) {
	log.logger.Errorw(template, keysAndValues...)
}

func (log loggerx) Fatal(args ...any) {
	log.logger.Fatal(args...)
}

func (log loggerx) Fatalf(template string, args ...any) {
	log.logger.Fatalf(template, args...)
}

func (log loggerx) Fatalw(template string, args ...any) {
	log.logger.Fatalw(template, args...)
}

func (log loggerx) Panic(args ...any) {
	log.logger.Panic(args...)
}

func (log loggerx) Panicf(template string, args ...any) {
	log.logger.Panicf(template, args...)
}

func (log loggerx) Panicw(template string, args ...any) {
	log.logger.Panicw(template, args...)
}

func (log loggerx) Sync() {
	_ = log.logger.Sync()
}

// ----------------------------------------------------------------

func loggerLevelMappings() map[string]zapcore.Level {
	return map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
		"fatal": zapcore.FatalLevel,
		"panic": zapcore.PanicLevel,
	}
}

func populateLoggerEncoder(config Config) zapcore.Encoder {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.EncodeLevel = zapcore.CapitalLevelEncoder

	if config.JsonEnabled {
		return zapcore.NewJSONEncoder(conf)
	}

	return zapcore.NewConsoleEncoder(conf)
}

func populateLoggerWriter(conf Config) (zapcore.WriteSyncer, error) {
	conf.Path = expression.TrinaryOperationString(emptyString == conf.Path, conf.Path, defaultPath)
	conf.FileName = expression.TrinaryOperationString(emptyString == conf.FileName, conf.FileName, defaultFileName)

	if exist := filepathz.Exists(conf.Path); !exist {
		if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	_core := populateLumberjackLogger(conf)

	syncer, err, done := handleMultiWriteSyncer(conf, _core)
	if done {
		return syncer, err
	}

	return zapcore.AddSync(_core), nil
}

func populateLumberjackLogger(conf Config) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filepath.Join(conf.Path, conf.FileName),
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAgeDay,
		Compress:   conf.CompressEnabled,
	}
}

func handleMultiWriteSyncer(conf Config, _logger *lumberjack.Logger) (zapcore.WriteSyncer, error, bool) {
	if conf.StdoutEnabled {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(_logger), zapcore.AddSync(os.Stdout)), nil, true
	}

	return nil, nil, false
}

// ----------------------------------------------------------------

// New create a Logger instance
func New(conf Config) (Logger, error) {
	mappings := loggerLevelMappings()

	writeSyncer, err := populateLoggerWriter(conf)
	if err != nil {
		return nil, perrors.Errorf("populate logger error:%v", err)
	}
	encoder := populateLoggerEncoder(conf)
	level, ok := mappings[conf.Level]
	if !ok {
		level = mappings[defaultLevel]
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)
	_core := zap.New(core)
	_loggerx := &loggerx{
		logger: _core.Sugar(),
	}

	return _loggerx, nil
}

// ----------------------------------------------------------------

func Instance() Logger {
	return _logger
}

// ----------------------------------------------------------------

func Debug(args ...any) {
	_logger.Debug(args...)
}

func Debugf(template string, args ...any) {
	_logger.Debugf(template, args...)
}

func Info(args ...any) {
	_logger.Info(args...)
}

func Infof(template string, args ...any) {
	_logger.Infof(template, args...)
}

func Infow(template string, keysAndValues ...any) {
	_logger.Infow(template, keysAndValues...)
}

func Warn(args ...any) {
	_logger.Warn(args...)
}

func Warnf(template string, args ...any) {
	_logger.Warnf(template, args...)
}

func Warnw(template string, keysAndValues ...any) {
	_logger.Warnw(template, keysAndValues...)
}

func Error(args ...any) {
	_logger.Error(args...)
}

func Errorf(template string, args ...any) {
	_logger.Errorf(template, args...)
}

func Errorw(template string, keysAndValues ...any) {
	_logger.Errorw(template, keysAndValues...)
}

func Fatal(args ...any) {
	_logger.Fatal(args...)
}

func Fatalf(template string, args ...any) {
	_logger.Fatalf(template, args...)
}

func Fatalw(template string, args ...any) {
	_logger.Fatalw(template, args...)
}

func Panic(args ...any) {
	_logger.Panic(args...)
}

func Panicf(template string, args ...any) {
	_logger.Panicf(template, args...)
}

func Panicw(template string, args ...any) {
	_logger.Panicw(template, args...)
}

func Sync() {
	_logger.Sync()
}
