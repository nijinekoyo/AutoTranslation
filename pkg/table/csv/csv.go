/*
 * @Author: nijineko
 * @Date: 2025-07-03 15:33:36
 * @LastEditTime: 2025-07-03 15:51:07
 * @LastEditors: nijineko
 * @Description: CSV表格数据处理实现
 * @FilePath: \AutoTranslation\pkg\table\csv\csv.go
 */
package csv

import (
	"encoding/csv"
	"os"
)

// CSV表格数据处理结构体
type CSVTable struct {
	fileHandle *os.File
	reader     *csv.Reader
	writer     *csv.Writer

	isClosed bool // 是否已关闭
}

/**
 * @description: 创建一个新的CSV表格处理实例
 * @param {string} FilePath CSV文件路径
 * @return {*CSVTable} 返回一个新的CSVTable实例
 * @return {error} 错误信息
 */
func New(FilePath string) (*CSVTable, error) {
	// 打开CSV文件，如果不存在则创建
	FileHandle, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &CSVTable{
		fileHandle: FileHandle,
		reader:     csv.NewReader(FileHandle),
		writer:     csv.NewWriter(FileHandle),
		isClosed:   false,
	}, nil
}

/**
 * @description: 关闭表格实例
 * @return {*}
 */
func (c *CSVTable) Close() error {
	// 刷新CSV写入器
	c.writer.Flush()

	// 关闭文件句柄
	if err := c.fileHandle.Close(); err != nil {
		return err
	}

	// 标记为已关闭
	c.isClosed = true

	return nil
}

/**
 * @description: 读取CSV表格数据
 * @return {[][]string} 表格数据
 * @return {error} 错误信息
 */
func (c *CSVTable) Read() ([][]string, error) {
	if c.isClosed {
		return nil, os.ErrClosed
	}

	return c.reader.ReadAll()
}

/**
 * @description: 写入CSV表格数据
 * @param {[][]string} Datas 要写入的数据
 * @return {error} 错误信息
 */
func (c *CSVTable) Write(Datas [][]string) error {
	if c.isClosed {
		return os.ErrClosed
	}

	for _, Data := range Datas {
		if err := c.writer.Write(Data); err != nil {
			return err
		}
	}

	return c.writer.Error()
}

/**
 * @description: 更新指定行的数据
 * @param {int} Row 行号，从1开始计数
 * @param {[]string} Data
 * @return {error} 错误信息
 */
func (c *CSVTable) UpdateLine(Row int, Data []string) error {
	if c.isClosed {
		return os.ErrClosed
	}

	// 转换为Row基索引
	Row--

	// 读取所有数据
	TableDatas, err := c.Read()
	if err != nil {
		return err
	}

	// 更新指定行的数据
	if Row < 0 || Row >= len(TableDatas) {
		return os.ErrInvalid
	}
	TableDatas[Row] = Data

	// 写回更新后的数据
	return c.Write(TableDatas)
}

/**
 * @description: 更新指定单元格的数据
 * @param {int} Row 行号，从1开始计数
 * @param {int} Col 列号，从1开始计数
 * @param {string} Data 数据内容
 * @return {error} 错误信息
 */
func (c *CSVTable) UpdateCell(Row, Col int, Data string) error {
	if c.isClosed {
		return os.ErrClosed
	}

	// 转换为Row基索引
	Row--
	// 转换为Col基索引
	Col--

	// 读取所有数据
	TableDatas, err := c.Read()
	if err != nil {
		return err
	}

	// 更新指定单元格的数据
	if Row < 0 || Row >= len(TableDatas) || Col < 0 || Col >= len(TableDatas[Row]) {
		return os.ErrInvalid
	}
	TableDatas[Row][Col] = Data

	// 写回更新后的数据
	return c.Write(TableDatas)
}

/**
 * @description: 插入新行数据
 * @param {[]string} Data 新行数据
 * @return {error} 错误信息
 */
func (c *CSVTable) Append(Data []string) error {
	if c.isClosed {
		return os.ErrClosed
	}

	// 读取当前数据
	TableDatas, err := c.Read()
	if err != nil {
		return err
	}

	// 添加新行数据
	TableDatas = append(TableDatas, Data)

	// 写回更新后的数据
	return c.Write(TableDatas)
}

/**
 * @description: 插入新行数据到指定位置
 * @param {int} Row 行号，从1开始计数
 * @param {[]string} Data 新行数据
 * @return {error} 错误信息
 */
func (c *CSVTable) Insert(Row int, Data []string) error {
	if c.isClosed {
		return os.ErrClosed
	}

	// 转换为Row基索引
	Row--

	// 读取当前数据
	TableDatas, err := c.Read()
	if err != nil {
		return err
	}

	// 检查行号是否有效
	if Row < 0 || Row > len(TableDatas) {
		return os.ErrInvalid
	}

	// 插入新行数据
	TableDatas = append(TableDatas[:Row], append([][]string{Data}, TableDatas[Row:]...)...)

	// 写回更新后的数据
	return c.Write(TableDatas)
}

/**
 * @description: 删除指定行数据
 * @param {int} Row 行号，从1开始计数
 * @return {error} 错误信息
 */
func (c *CSVTable) Delete(Row int) error {
	if c.isClosed {
		return os.ErrClosed
	}

	// 转换为Row基索引
	Row--

	// 读取当前数据
	TableDatas, err := c.Read()
	if err != nil {
		return err
	}

	// 检查行号是否有效
	if Row < 0 || Row >= len(TableDatas) {
		return os.ErrInvalid
	}

	// 删除指定行数据
	TableDatas = append(TableDatas[:Row], TableDatas[Row+1:]...)

	// 写回更新后的数据
	return c.Write(TableDatas)
}