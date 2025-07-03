/*
 * @Author: nijineko
 * @Date: 2025-07-03 18:21:28
 * @LastEditTime: 2025-07-03 18:23:05
 * @LastEditors: nijineko
 * @Description: 目录操作工具
 * @FilePath: \AutoTranslation\tools\file\directory.go
 */
package file

import (
	"os"
	"path/filepath"
)

/**
 * @description: 获取指定目录下的所有文件路径
 * @param {string} Path 目录路径
 * @return {[]string} 返回文件路径列表
 * @return {error} 错误信息
 */
func GetDirectoryFilePaths(Path string) ([]string, error) {
	var FileList []string
	Files, err := os.ReadDir(Path)
	if err != nil {
		return FileList, err
	}
	for _, File := range Files {
		if !File.IsDir() {
			FileList = append(FileList, filepath.Join(Path, File.Name()))
		} else {
			// 如果是文件夹，则递归获取文件夹内的文件
			SubFiles, subErr := GetDirectoryFilePaths(filepath.Join(Path, File.Name()))
			if subErr != nil {
				return FileList, subErr
			}
			FileList = append(FileList, SubFiles...)
		}
	}
	return FileList, err
}
