/*
 * @Author: nijineko
 * @Date: 2025-07-03 16:34:04
 * @LastEditTime: 2025-07-03 16:34:12
 * @LastEditors: nijineko
 * @Description: 翻译包
 * @FilePath: \AutoTranslation\pkg\translation\translation.go
 */
package translation

type Translation interface {
	TranslateText(Text string, SourceLanguage *string, TargetLang string) (string, error)
}

/**
 * @description: 翻译文本
 * @param {Translation} TranslatorInstance 翻译器实例
 * @param {string} Text 要翻译的文本
 * @param {*string} SourceLanguage 源语言，为nil表示自动检测
 * @param {string} TargetLanguage 目标语言
 * @return {string} 返回翻译后的文本
 * @return {error} 错误信息
 */
func TranslateText(TranslatorInstance Translation, Text string, SourceLanguage *string, TargetLanguage string) (string, error) {
	// 调用翻译器进行翻译
	return TranslatorInstance.TranslateText(Text, SourceLanguage, TargetLanguage)
}