package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strings"
)

/**
初始化日志配置
*/
func initLogger() (err error) {

	config := make(map[string]interface{})
	config["filename"] = appConfig.logPath
	config["level"] = convertLogLevel(appConfig.logLevel)

	//序列化为JSON
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

/**
文字转换日志消息级别
*/
func convertLogLevel(logLevel string) int {
	lowerLogLevel := strings.ToLower(logLevel)
	switch lowerLogLevel {
	case "trace":
		return logs.LevelTrace
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "error":
		return logs.LevelError
	}
	return logs.LevelDebug
}
