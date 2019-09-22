package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

/**
beego logs模块的使用
*/
func main() {

	//配置log组件
	config := make(map[string]interface{})
	config["filename"] = "./logs/application.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	//初始化日志组件
	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is debug record")
	logs.Info("this is info record")
	logs.Warn("this is warn record")

}
