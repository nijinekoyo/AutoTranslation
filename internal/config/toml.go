/*
 * @Author: nijineko
 * @Date: 2025-07-03 17:29:00
 * @LastEditTime: 2025-07-03 17:30:15
 * @LastEditors: nijineko
 * @Description: TOML配置文件解析
 * @FilePath: \AutoTranslation\internal\config\toml.go
 */
package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

// TOML配置文件结构
type TomlConfig struct{}

/**
 * @description: 创建一个新的TOML配置实例
 * @return {*TomlConfig} TomlConfig实例
 */
func NewTomlConfig() *TomlConfig {
	return &TomlConfig{}
}

/**
 * @description: 创建空白配置文件
 * @param {string} FilePath 配置文件路径
 * @return {error} 错误
 */
func (t *TomlConfig) Create(FilePath string) error {
	// 创建配置文件
	ConfigBytes, err := toml.Marshal(&Config{})
	if err != nil {
		return err
	}

	// 写入配置文件
	err = os.WriteFile(FilePath, ConfigBytes, 0664)
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 初始化配置文件
 * @param {string} FilePath 配置文件路径
 * @return {Config} 配置数据
 * @return {error} 错误
 */
func (t *TomlConfig) Get(FilePath string) (Config, error) {
	// 读取配置文件
	ConfigBytes, err := os.ReadFile(FilePath)
	if err != nil {
		return Data, err
	}

	// 反序列化，写入全局变量
	err = toml.Unmarshal(ConfigBytes, &Data)
	if err != nil {
		panic(err)
	}

	return Data, nil
}
