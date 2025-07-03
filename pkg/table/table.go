/*
 * @Author: nijineko
 * @Date: 2025-07-03 15:28:47
 * @LastEditTime: 2025-07-03 16:33:50
 * @LastEditors: nijineko
 * @Description: 表格数据处理包
 * @FilePath: \AutoTranslation\pkg\table\table.go
 */
package table

// 表格数据处理接口
type Table interface {
	Close() error
	Read() ([][]string, error)
	Write(Datas [][]string) error
	UpdateLine(Row int, Data []string) error
	UpdateCell(Row, Col int, Data string) error
	Append(Data []string) error
	Insert(Row int, Data []string) error
	Delete(Row int) error
}

/**
 * @description: 读取表格数据
 * @param {Table} TableInstance 表格实例
 * @return {[][]string} 返回表格数据
 * @return {error} 错误信息
 */
func ReadTable(TableInstance Table) ([][]string, error) {
	// 读取表格数据
	return TableInstance.Read()

}

/**
 * @description: 写入表格数据
 * @param {Table} TableInstance 表格实例
 * @param {[][]string} Datas 表格数据
 * @return {error} 错误信息
 */
func WriteTable(TableInstance Table, Datas [][]string) error {
	// 写入表格数据
	return TableInstance.Write(Datas)
}