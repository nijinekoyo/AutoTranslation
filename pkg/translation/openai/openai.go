/*
 * @Author: nijineko
 * @Date: 2025-07-03 17:09:58
 * @LastEditTime: 2025-07-03 17:45:16
 * @LastEditors: nijineko
 * @Description: OpenAI翻译实现
 * @FilePath: \AutoTranslation\pkg\translation\openai\openai.go
 */
package openai

import (
	"context"
	"errors"

	"github.com/nijinekoyo/AutoTranslation/internal/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAITranslator struct {
	APIKey string

	client openai.Client
}

var (
	ErrInvalidRole = errors.New("invalid role in OpenAI messages")
)

/**
 * @description: 创建一个新的OpenAI翻译实例
 * @param {string} APIKey OpenAI API密钥
 * @return {*OpenAITranslator} OpenAITranslator实例
 */
func New(APIKey string) *OpenAITranslator {
	return &OpenAITranslator{
		APIKey: APIKey,
		client: openai.NewClient(
			option.WithAPIKey(APIKey),
			option.WithBaseURL(config.Get().Translation.OpenAI.BaseURL),
		),
	}
}

/**
 * @description: 翻译文本
 * @param {string} Text 要翻译的文本
 * @param {*string} SourceLanguage 源语言，为nil表示自动检测
 * @param {string} TargetLanguage 目标语言
 * @return {string} 返回翻译后的文本
 * @return {error} 错误信息
 */
func (o *OpenAITranslator) TranslateText(Text string, SourceLanguage *string, TargetLanguage string) (string, error) {
	var Messages []openai.ChatCompletionMessageParamUnion

	// 添加前置消息
	for _, MessageStr := range config.Get().Translation.OpenAI.Messages {
		switch MessageStr.Role {
		case "user":
			Messages = append(Messages, openai.UserMessage(MessageStr.Content))
		case "system":
			Messages = append(Messages, openai.SystemMessage(MessageStr.Content))
		case "developer":
			Messages = append(Messages, openai.DeveloperMessage(MessageStr.Content))
		case "assistant":
			Messages = append(Messages, openai.AssistantMessage(MessageStr.Content))
		default:
			return "", ErrInvalidRole
		}
	}

	// 添加术语表
	if len(config.Get().Translation.LargeLanguageModel.Glossaries) > 0 {
		Messages = append(Messages, openai.AssistantMessage(config.Get().Translation.LargeLanguageModel.GlossaryPrompt))
	}
	for _, Glossary := range config.Get().Translation.LargeLanguageModel.Glossaries {
		GlossaryMessage := Glossary.Name + ": " + Glossary.Description + "\n"
		for _, Entry := range Glossary.Entries {
			GlossaryMessage += Entry.Source + " -> " + Entry.Target + "\n"
		}
		Messages = append(Messages, openai.AssistantMessage(GlossaryMessage))
	}

	// 添加文本
	Messages = append(Messages, openai.UserMessage(Text))

	ChatCompletion, err := o.client.Chat.Completions.New(context.Background(),
		openai.ChatCompletionNewParams{
			Messages: Messages,
			Model:    config.Get().Translation.OpenAI.Model,
		},
	)
	if err != nil {
		return "", err
	}

	return ChatCompletion.Choices[0].Message.Content, nil
}
