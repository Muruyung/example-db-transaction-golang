package logger

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type key string

const (
	// ActivityID generate uuid
	ActivityID key = "activityID"

	// IsUseES check log using elastic
	IsUseES key = "isUseES"
)

// APILogger api logger
func APILogger(ctx context.Context, method, path string, body interface{}) {
	var (
		data = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   path,
			"data":       body,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"method": method,
			"path":   path,
			"body":   data,
		}
		Logger.WithFields(logFields).Info("HTTP")

		logFields["date"] = time.Now().UnixMilli()
		Logger.Elastic.Index(Logger.Elastic.GetIndex()).Request(logFields).Do(ctx)
	}
}

// DetailLoggerInfo detail logger info
func DetailLoggerInfo(ctx context.Context, command, details string, logData interface{}) {
	var (
		isUseES = ctx.Value(IsUseES) == true
		data    = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   details,
			"data":       logData,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"command": command,
			"details": data,
		}
		Logger.WithFields(logFields).Info("Function")

		if isUseES {
			logFields["date"] = time.Now().UnixMilli()
			Logger.Elastic.Index(Logger.Elastic.GetIndex()).Request(logFields).Do(ctx)
		}
	}
}

// DetailLoggerError detail logger error
func DetailLoggerError(ctx context.Context, command, details string, err ...interface{}) {
	var (
		isUseES     = ctx.Value(IsUseES) == true
		reflectData = reflect.ValueOf(err)
		dataString  = fmt.Sprintf("%v", reflectData)
		data        = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   details,
			"data":       dataString,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"command": command,
			"details": data,
		}
		Logger.WithFields(logFields).Error("Function")

		if isUseES {
			logFields["date"] = time.Now().UnixMilli()
			Logger.Elastic.Index(Logger.Elastic.GetIndex()).Request(logFields).Do(ctx)
		}
	}
}

// DetailLoggerWarn detail logger warning
func DetailLoggerWarn(ctx context.Context, command, details string, warn ...interface{}) {
	var (
		isUseES     = ctx.Value(IsUseES) == true
		reflectData = reflect.ValueOf(warn)
		dataString  = fmt.Sprintf("%v", reflectData.Interface())
		data        = map[string]interface{}{
			"activityID": ctx.Value(ActivityID),
			"activity":   details,
			"data":       dataString,
		}
	)
	if Logger.Name != "test" {
		logFields := map[string]interface{}{
			"command": command,
			"details": data,
		}
		Logger.WithFields(logFields).Warn("Function")

		if isUseES {
			logFields["date"] = time.Now().UnixMilli()
			Logger.Elastic.Index(Logger.Elastic.GetIndex()).Request(logFields).Do(ctx)
		}
	}
}
