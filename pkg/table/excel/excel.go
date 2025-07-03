/*
 * @Author: nijineko
 * @Date: 2025-07-03 15:54:30
 * @LastEditTime: 2025-07-03 16:22:34
 * @LastEditors: nijineko
 * @Description: Excel表格数据处理实现
 * @FilePath: \AutoTranslation\pkg\table\excel\excel.go
 */
package excel

import (
	"os"

	"github.com/xuri/excelize/v2"
)

type ExcelTable struct {
	filePath       string
	excelizeHandle *excelize.File

	defaultSheetName string // 默认工作表名称

	isClosed bool // 是否已关闭
}

/**
 * @description: 创建一个新的Excel表格处理实例
 * @param {string} FilePath Excel文件路径
 * @return {*ExcelTable} 返回一个新的ExcelTable实例
 * @return {error} 错误信息
 */
func New(FilePath string) (*ExcelTable, error) {
	// 创建一个新的Excel文件处理实例
	ExcelizeHandle, err := excelize.OpenFile(FilePath)
	if err != nil {
		// 如果文件不存在，则创建一个新的Excel文件
		ExcelizeHandle = excelize.NewFile()
		if err := ExcelizeHandle.SaveAs(FilePath); err != nil {
			return nil, err
		}
	}

	// 获取默认工作表名称
	DefaultSheetName := ExcelizeHandle.GetSheetName(0)

	return &ExcelTable{
		filePath:         FilePath,
		excelizeHandle:   ExcelizeHandle,
		defaultSheetName: DefaultSheetName,
		isClosed:         false,
	}, nil
}

/**
 * @description: 关闭Excel表格处理实例
 * @return {error} 错误信息
 */
func (e *ExcelTable) Cose() error {
	// 保存Excel文件
	if err := e.excelizeHandle.SaveAs(e.filePath); err != nil {
		return err
	}
	// 关闭Excelize文件句柄
	if err := e.excelizeHandle.Close(); err != nil {
		return err
	}

	// 标记为已关闭
	e.isClosed = true

	return nil
}

/**
 * @description: 读取Excel表格数据
 * @return {[][]string} 表格数据
 * @return {error} 错误信息
 */
func (e *ExcelTable) Read() ([][]string, error) {
	// 读取Excel文件中的数据
	rows, err := e.excelizeHandle.GetRows(e.defaultSheetName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

/**
 * @description: 写入Excel表格数据
 * @param {[][]string} Datas 要写入的数据
 * @return {error} 错误信息
 */
func (e *ExcelTable) Write(Datas [][]string) error {
	if e.isClosed {
		return os.ErrClosed
	}

	// 写入新数据到工作表
	for RowIndex, RowData := range Datas {
		for ColIndex, CellData := range RowData {
			Cell, err := excelize.CoordinatesToCellName(ColIndex+1, RowIndex+1)
			if err != nil {
				return err
			}
			if err := e.excelizeHandle.SetCellValue(e.defaultSheetName, Cell, CellData); err != nil {
				return err
			}
		}
	}

	return nil
}

/**
 * @description: 更新指定行的数据
 * @param {int} Row 行号，从1开始计数
 * @param {[]string} Data 要更新的数据
 * @return {error} 错误信息
 */
func (e *ExcelTable) UpdateLine(Row int, Data []string) error {
	if e.isClosed {
		return os.ErrClosed
	}

	// 更新指定行的数据
	for ColIndex, CellData := range Data {
		Cell, err := excelize.CoordinatesToCellName(ColIndex+1, Row)
		if err != nil {
			return err
		}
		if err := e.excelizeHandle.SetCellValue(e.defaultSheetName, Cell, CellData); err != nil {
			return err
		}
	}

	return nil
}

/**
 * @description: 更新指定单元格的数据
 * @param {int} Row 行号，从1开始计数
 * @param {int} Col 列号，从1开始计数
 * @param {string} Data 数据内容
 * @return {error} 错误信息
 */
func (e *ExcelTable) UpdateCell(Row, Col int, Data string) error {
	if e.isClosed {
		return os.ErrClosed
	}

	// 更新指定单元格的数据
	Cell, err := excelize.CoordinatesToCellName(Col, Row)
	if err != nil {
		return err
	}
	if err := e.excelizeHandle.SetCellValue(e.defaultSheetName, Cell, Data); err != nil {
		return err
	}

	return nil
}

/**
 * @description: 在Excel表格中追加一行数据
 * @param {[]string} Data 要追加的数据
 * @return {error} 错误信息
 */
func (e *ExcelTable) Append(Data []string) error {
	if e.isClosed {
		return os.ErrClosed
	}

	// 获取当前行数
	Rows, err := e.excelizeHandle.GetRows(e.defaultSheetName)
	if err != nil {
		return err
	}
	CurrentRow := len(Rows) + 1 // 新数据将添加到最后一行

	// 写入新数据到工作表
	for ColIndex, CellData := range Data {
		Cell, err := excelize.CoordinatesToCellName(ColIndex+1, CurrentRow)
		if err != nil {
			return err
		}
		if err := e.excelizeHandle.SetCellValue(e.defaultSheetName, Cell, CellData); err != nil {
			return err
		}
	}

	return nil
}

/**
 * @description: 在Excel表格中插入一行数据
 * @param {int} Row 行号，从1开始计数
 * @param {[]string} Data 要插入的数据
 * @return {error} 错误信息
 */
func (e *ExcelTable) Insert(Row int, Data []string) error {
	if e.isClosed {
		return os.ErrClosed
	}

	// 获取当前行数
	Rows, err := e.excelizeHandle.GetRows(e.defaultSheetName)
	if err != nil {
		return err
	}

	// 检查插入行号是否有效
	if Row < 1 || Row > len(Rows)+1 {
		return os.ErrInvalid
	}

	// 插入空白行
	if err := e.excelizeHandle.InsertRows(e.defaultSheetName, Row, 1); err != nil {
		return err
	}

	// 写入新数据到指定行
	for ColIndex, CellData := range Data {
		Cell, err := excelize.CoordinatesToCellName(ColIndex+1, Row)
		if err != nil {
			return err
		}
		if err := e.excelizeHandle.SetCellValue(e.defaultSheetName, Cell, CellData); err != nil {
			return err
		}
	}

	return nil
}

/**
 * @description: 删除指定行的数据
 * @param {int} Row 行号，从1开始计数
 * @return {error} 错误信息
 */
func (e *ExcelTable) Delete(Row int) error {
	if e.isClosed {
		return os.ErrClosed
	}

	// 删除指定行
	return e.excelizeHandle.RemoveRow(e.defaultSheetName, Row)
}
