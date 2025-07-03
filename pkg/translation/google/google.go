/*
 * @Author: nijineko
 * @Date: 2025-07-03 16:41:28
 * @LastEditTime: 2025-07-03 17:00:49
 * @LastEditors: nijineko
 * @Description: Google翻译实现
 * @FilePath: \AutoTranslation\pkg\translation\google\google.go
 */
package google

import (
	"errors"

	"github.com/HyacinthusAcademy/yuzuhttp"
)

const (
	APIURL = "https://translate.google.com/translate_a/single"
)

var (
	// 翻译数据响应为空
	ErrResponseEmpty = errors.New("ranslation failed: no response data")
	// 翻译数据格式错误
	ErrResponseFormat = errors.New("translation failed: unexpected response format")
)

// Goole翻译结构体
type GoogleTranslator struct{}

/**
 * @description: 创建一个新的Google翻译实例
 * @return {*GoogleTranslator} GoogleTranslator实例
 */
func New() *GoogleTranslator {
	return &GoogleTranslator{}
}

/**
 * @description: 翻译文本
 * @param {string} Text 要翻译的文本
 * @param {*string} SourceLanguage 源语言，为nil表示自动检测
 * @param {string} TargetLanguage 目标语言
 * @return {string} 返回翻译后的文本
 * @return {error} 错误信息
 */
func (g *GoogleTranslator) TranslateText(Text string, SourceLanguage *string, TargetLanguage string) (string, error) {
	Request := yuzuhttp.Get(APIURL).
		AddQuery("client", "gtx").
		AddQuery("tl", TargetLanguage).
		AddQuery("dt", "t").
		AddQuery("q", Text)
	if SourceLanguage != nil {
		Request.AddQuery("sl", *SourceLanguage)
	} else {
		Request.AddQuery("sl", "auto")
	}

	var ResponseData []any
	if err := Request.Do().BodyJSON(&ResponseData); err != nil {
		return "", err
	}

	for _, Item := range ResponseData {
		if Item == nil {
			continue
		}

		if Data, ok := Item.([]any); ok && len(Data) > 0 {
			for _, SubItem := range Data {
				if SubItem == nil {
					continue
				}

				if SubData, ok := SubItem.([]any); ok && len(SubData) > 0 {
					if TranslationText, ok := SubData[0].(string); ok {
						return TranslationText, nil
					}
				}
			}
		}
	}

	return "", ErrResponseFormat
}
