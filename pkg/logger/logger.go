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
	`os`
	`path/filepath`

	`github.com/photowey/nettygo/interal/helper`
	`github.com/photowey/perrors`
	`go.uber.org/zap`
	`go.uber.org/zap/zapcore`
	`gopkg.in/natefinch/lumberjack.v2`
)

const (
	defaultPath     = "logs"
	defaultFileName = "nettygo.log"
	defaultLevel    = "debug"
	emptyString     = ""
)

var logger *zap.SugaredLogger

// Init populate logger instance
func Init(conf Config) error {
	mappings := loggerLevelMappings()

	writeSyncer, err := populateLoggerWriter(conf)
	if err != nil {
		return perrors.Errorf("populate logger error:%v", err)
	}
	encoder := populateLoggerEncoder(conf)
	level, ok := mappings[conf.Level]
	if !ok {
		level = mappings[defaultLevel]
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)
	_logger := zap.New(core)
	logger = _logger.Sugar()

	return nil
}

func loggerLevelMappings() map[string]zapcore.Level {
	return map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
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
	conf.Path, _ = helper.Match(emptyString == conf.Path, conf.Path, defaultPath).(string)
	conf.FileName, _ = helper.Match(emptyString == conf.FileName, conf.FileName, defaultFileName).(string)

	if exist := helper.IsExist(conf.Path); !exist {
		if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	_logger := populateLumberjackLogger(conf)

	syncer, err, done := handleMultiWriteSyncer(conf, _logger)
	if done {
		return syncer, err
	}

	return zapcore.AddSync(_logger), nil
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

func Logger() *zap.SugaredLogger {
	return logger
}

// ----------------------------------------------------------------

func Debug(template string, args ...any) {
	logger.Debugf(template, args)
}

func Info(template string, args ...any) {
	logger.Infof(template, args)
}

func Infow(template string, keysAndValues ...any) {
	logger.Infow(template, keysAndValues)
}

func Warn(template string, args ...any) {
	logger.Warnf(template, args)
}

func Warnw(template string, keysAndValues ...any) {
	logger.Warnw(template, keysAndValues)
}

func Error(template string, args ...any) {
	logger.Errorf(template, args)
}

func Errorw(template string, keysAndValues ...any) {
	logger.Errorw(template, keysAndValues)
}

func Fatal(template string, args ...any) {
	logger.Fatalf(template, args)
}

func Panic(template string, args ...any) {
	logger.Panicf(template, args)
}

func Sync() {
	_ = logger.Sync()
}
