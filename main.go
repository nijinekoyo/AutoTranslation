/*
 * @Author: nijineko
 * @Date: 2025-07-03 15:26:27
 * @LastEditTime: 2025-07-03 20:53:02
 * @LastEditors: nijineko
 * @Description: main package
 * @FilePath: \AutoTranslation\main.go
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nijinekoyo/AutoTranslation/bootstrap"
	"github.com/nijinekoyo/AutoTranslation/internal/config"
	"github.com/nijinekoyo/AutoTranslation/internal/log"
	"github.com/nijinekoyo/AutoTranslation/pkg/table"
	"github.com/nijinekoyo/AutoTranslation/pkg/table/csv"
	"github.com/nijinekoyo/AutoTranslation/pkg/table/excel"
	"github.com/nijinekoyo/AutoTranslation/pkg/translation"
	"github.com/nijinekoyo/AutoTranslation/pkg/translation/google"
	"github.com/nijinekoyo/AutoTranslation/pkg/translation/openai"
	"github.com/nijinekoyo/AutoTranslation/tools/file"
	"github.com/noa-log/colorize"
)

func main() {
	// 初始化系统
	bootstrap.Init()

	// 获取第一个命令参数作为待翻译文件路径
	FilePathData := os.Args[1]

	// 检查传入路径是文件夹还是文件
	FileInfo, err := os.Stat(FilePathData)
	if err != nil {
		log.Print().Error("System", err)
		return
	}

	// 待翻译文件列表
	var FilePaths []string

	if FileInfo.IsDir() {
		// 如果是文件夹，则获取文件夹下所有文件
		FilePaths, err = file.GetDirectoryFilePaths(FilePathData)
		if err != nil {
			log.Print().Error("System", err)
			return
		}
	} else {
		// 如果是文件，则直接添加到列表
		FilePaths = append(FilePaths, FilePathData)
	}

	// 按照配置文件指定翻译器
	var TranslatorInstance translation.Translation
	switch config.Get().Translation.Service {
	case "google":
		TranslatorInstance = google.New()
	case "openai":
		TranslatorInstance = openai.New(config.Get().Translation.OpenAI.APIKey)
	default:
		log.Print().Error("System", "Unsupported translation service: "+config.Get().Translation.Service)
		return
	}

	// 遍历待翻译文件列表
	for _, FilePath := range FilePaths {
		var TableInstance table.Table

		// 通过扩展名需要使用的表格处理器
		switch filepath.Ext(FilePath) {
		case ".xlsx", ".xls":
			TableInstance, err = excel.New(FilePath)
			if err != nil {
				log.Print().Error("Translation", err)
				return
			}
		case ".csv":
			TableInstance, err = csv.New(FilePath)
			if err != nil {
				log.Print().Error("Translation", err)
				return
			}
		default:
			log.Print().Error("Translation", "Unsupported table format: "+filepath.Ext(FilePath))
		}
		defer TableInstance.Close()

		// 读取表格数据
		TableDatas, err := TableInstance.Read()
		if err != nil {
			log.Print().Error("Translation", "Failed to read table data from "+FilePath, err)
			return
		}

		// 遍历表格数据进行翻译
		for Index, Row := range TableDatas {
			if config.Get().SkipTableHeader && Index == 0 {
				// 如果跳过表头，则继续下一行
				continue
			}

			// 计算源列和目标列索引
			SourceColumn := config.Get().SourceColumn - 1 // 转换为0开始计数
			TargetColumn := config.Get().TargetColumn - 1 // 转换为0开始
			if len(Row) <= SourceColumn {
				log.Print().Error("Translation", fmt.Sprintf("Row %d: Source column index %d is out of range", Index+1, SourceColumn+1))
				continue
			}
			if len(TableDatas[Index]) <= TargetColumn {
				// 如果目标列索引超出范围，则扩展行数据
				for len(TableDatas[Index]) <= TargetColumn {
					TableDatas[Index] = append(TableDatas[Index], "")
				}
			}

			// 获取待翻译文本
			SourceText := Row[SourceColumn]

			if config.Get().SkipIfNotEmpty && TableDatas[Index][TargetColumn] != "" {
				// 如果待翻译单元格不为空且配置了跳过，则跳过翻译
				log.Print().Warning("Translation", fmt.Sprintf("Row %d: cell is not empty, skipping translation", Index+1))
				continue
			}

			// 翻译文本
			TranslatedText, err := TranslatorInstance.TranslateText(SourceText, config.Get().Translation.SourceLanguage, config.Get().Translation.TargetLanguage)
			if err != nil {
				log.Print().Error("Translation", err)
				continue
			}

			// 更新翻译结果到目标列
			TableDatas[Index][TargetColumn] = TranslatedText

			log.Print().Info("Translation", fmt.Sprintf("Row %d: %s -> %s", Index+1, colorize.YellowText(SourceText), colorize.GreenText(TranslatedText)))
		}

		// 保存翻译后的表格数据
		if err := TableInstance.Write(TableDatas); err != nil {
			log.Print().Error("System", err)
			return
		}
		log.Print().Info("Translation", fmt.Sprintf("Translation completed for file: %s", FilePath))
	}
}
