package log

import "github.com/noa-log/noa"

var (
	LogConfig *noa.LogConfig // 默志实例
)

/**
 * @description: 初始化日志实例
 * @param {*}
 */
func init() {
	if LogConfig == nil {
		LogConfig = noa.NewLog()
	}
}

/**
 * @description: 获取日志实例
 * @return {*noa.LogConfig} 返回日志配置实例
 */
func Print() *noa.LogConfig {
	return LogConfig
}