// Copyright 2014 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package log is a wrapper of logs for short calling name.
package log

import (
	"os"

	"github.com/gogits/logs"
)

var (
	loggers []*logs.BeeLogger
)

func init() {
	NewLogger(0, "console", `{"level": 0}`)
}

func NewLogger(bufLen int64, mode, config string) {
	logger := logs.NewLogger(bufLen)

	isExist := false
	for _, l := range loggers {
		if l.Adapter == mode {
			isExist = true
			l = logger
		}
	}
	if !isExist {
		loggers = append(loggers, logger)
	}
	logger.SetLogFuncCallDepth(3)
	logger.SetLogger(mode, config)
}

func Trace(format string, v ...interface{}) {
	for _, logger := range loggers {
		logger.Trace(format, v...)
	}
}

func Debug(format string, v ...interface{}) {
	for _, logger := range loggers {
		logger.Debug(format, v...)
	}
}

func Info(format string, v ...interface{}) {
	for _, logger := range loggers {
		logger.Info(format, v...)
	}
}

func Error(format string, v ...interface{}) {
	for _, logger := range loggers {
		logger.Error(format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	for _, logger := range loggers {
		logger.Warn(format, v...)
	}
}

func Critical(format string, v ...interface{}) {
	for _, logger := range loggers {
		logger.Critical(format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	Error(format, v...)
	for _, l := range loggers {
		l.Close()
	}
	os.Exit(2)
}
