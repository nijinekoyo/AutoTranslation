/*
 * @Author: nijineko
 * @Date: 2025-07-03 17:26:25
 * @LastEditTime: 2025-07-03 17:42:09
 * @LastEditors: nijineko
 * @Description: 配置文件模块
 * @FilePath: \AutoTranslation\internal\config\config.go
 */
package config

// 配置文件结构
type Config struct {
	SourceColumn int `toml:"source_column"` // 待翻译列，从1开始计数
	TargetColumn int `toml:"target_column"` // 翻译目标列，从1开始计数

	SkipTableHeader bool `toml:"skip_table_header"` // 翻译时是否跳过表头
	SkipIfNotEmpty  bool `toml:"skip_if_not_empty"` // 如果待翻译单元格不为空，则跳过翻译

	Translation struct {
		Service string `toml:"service"` // 需要使用的翻译服务 (google, openai, etc.)

		// 语言配置
		SourceLanguage *string `toml:"source_language"` // 源语言，为nil表示自动检测
		TargetLanguage string  `toml:"target_language"` // 目标语言

		LargeLanguageModel struct {
			GlossaryPrompt string `toml:"glossary_prompt"` // 术语表提示，提示大语言模型使用术语表进行翻译
			Glossaries     []struct {
				Name        string `toml:"name"`        // 术语表名称
				Description string `toml:"description"` // 术语表描述
				Entries     []struct {
					Source string `toml:"source"` // 源语言术语
					Target string `toml:"target"` // 目标语言术语
				} `toml:"entries"` // 术语表条目
			} `toml:"glossaries"` // 翻译术语表
		} // 大语言模型翻译配置

		OpenAI struct {
			BaseURL  string `toml:"base_url"` // OpenAI API URL
			APIKey   string `toml:"api_key"`  // API密钥
			Model    string `toml:"model"`    // 模型名称
			Messages []struct {
				Role    string `toml:"role"`    // 消息角色 (user, system, developer, assistant)
				Content string `toml:"content"` // 消息内容 (消息需要约束返回必须是翻译后的文本，而且必须是纯文本，不受语言配置影响)
			} `toml:"messages"` // 请求前置消息
		} `toml:"openai"` // OpenAI翻译配置
	} `toml:"translation"` // 翻译配置
}

// 全局参数
var Data Config

/**
 * @description: 获取配置
 * @return {ConfigData} 配置
 */
func Get() Config {
	return Data
}

type ConfigInstance interface {
	Create(FilePath string) error        // 创建配置文件
	Get(FilePath string) (Config, error) // 获取配置
}
