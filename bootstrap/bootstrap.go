/*
 * @Author: nijineko
 * @Date: 2025-07-03 18:14:35
 * @LastEditTime: 2025-07-03 18:15:12
 * @LastEditors: nijineko
 * @Description: 启动初始化
 * @FilePath: \AutoTranslation\bootstrap\bootstrap.go
 */
package bootstrap

import (
	"path/filepath"

	"github.com/nijinekoyo/AutoTranslation/internal/config"
	"github.com/nijinekoyo/AutoTranslation/internal/log"
)

const (
	CONFIG_PATH = "config.toml" // 配置文件路径
)

func Init() {
	// 读取配置文件
	var ConfigInstance config.ConfigInstance

	switch filepath.Ext(CONFIG_PATH) {
	case ".toml":
		ConfigInstance = config.NewTomlConfig()
	default:
		panic("Unsupported config file format: " + filepath.Ext(CONFIG_PATH))
	}

	// 获取配置
	ConfigData, err := ConfigInstance.Get(CONFIG_PATH)
	if err != nil {
		log.Print().Error("System", err)
	}
	// 赋值到全局配置
	config.Data = ConfigData
}
