package goutils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/xuri/excelize/v2"
)

// escapeCSVCell 处理单元格内容，如果包含逗号，则用双引号包裹
func escapeCSVCell(cell string) string {
	if strings.Contains(cell, ",") {
		return fmt.Sprintf("\"%s\"", cell)
	}
	return cell
}

// ReadExcelToCSV 读取 XLSX 文件并返回 CSV 格式的字符串
func ReadExcelToCSV(filePath string) (string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 获取第一个工作表的名称
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return "", fmt.Errorf("no sheets found in the file")
	}
	sheetName := sheets[0]

	// 读取工作表中的所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return "", err
	}

	// 将每一行转换为 CSV 格式的字符串
	var csvBuilder strings.Builder
	for _, row := range rows {
		for i, cell := range row {
			row[i] = escapeCSVCell(cell)
		}
		csvRow := strings.Join(row, ",")
		csvBuilder.WriteString(fmt.Sprintf("%s\n", csvRow))
	}

	return csvBuilder.String(), nil
}

// 读取XLSX文件到Gota DataFrame中
func ReadExcelToGotaDF(filePath string) (dataframe.DataFrame, error) {
	csvString, err := ReadExcelToCSV(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(csvString)
	df := dataframe.ReadCSV(strings.NewReader(csvString))
	return df, nil
}

// WriteDataFrameToFile 将 Gota DataFrame 写入 CSV 或 XLSX 文件
// outputPath: 输出文件的路径
// delimiter: CSV 文件的分隔符（默认为逗号）
func WriteDataFrameToFile(df dataframe.DataFrame, outputPath string, delimiter ...rune) error {
	ext := strings.ToLower(filepath.Ext(outputPath))
	sep := ','
	if len(delimiter) > 0 {
		sep = delimiter[0]
	}

	switch ext {
	case ".csv":
		return WriteToCSV(df, outputPath, sep)
	case ".xlsx":
		return WriteToExcel(df, outputPath)
	case ".xls":
		panic("不支持保存为xls格式文件")
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
}

// WriteToCSV 将Gota DataFrame 写入 CSV 文件
func WriteToCSV(df dataframe.DataFrame, outputPath string, sep rune) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = sep
	defer w.Flush()

	// 写入 DataFrame
	if err := w.WriteAll(df.Records()); err != nil {
		return err
	}
	return nil
}

// WriteToExcel 将Gota DataFrame 写入 Excel XLSX文件,暂不支持xls文件写入
func WriteToExcel(df dataframe.DataFrame, outputPath string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// 写入 DataFrame
	for i, record := range df.Records() {
		for j, cell := range record {
			cellName, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue(sheetName, cellName, cell)
		}
	}

	if err := f.SaveAs(outputPath); err != nil {
		return err
	}
	return nil
}

// ReadExcelCol 根据文件路径、Sheet名称、列索引或标题读取指定列的数据
// 如果 sheetName 为空，则默认使用第一个 Sheet
// colIndex 是列的数字索引（从1开始），title 是列的标题
// 两者只能传一个，如果两者都传或都不传，将返回错误
func ReadExcelCol(filePath, title string, colIndex int, sheetName string) ([]string, error) {
	// 检查参数
	if (colIndex != 0 && title != "") || (colIndex == 0 && title == "") {
		return nil, fmt.Errorf("必须且只能指定 colIndex 或 title 其中一个")
	}

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 如果没有指定 sheetName，则使用第一个 Sheet
	if sheetName == "" {
		sheets := f.GetSheetList()
		if len(sheets) == 0 {
			return nil, fmt.Errorf("Excel File Specifal had no Sheet")
		}
		sheetName = sheets[0]
	}

	var colNum int
	if colIndex != 0 {
		colNum = colIndex
	} else {
		colNum, err = FindColumnIndexByTitle(f, sheetName, title)
		if err != nil {
			return nil, err
		}
	}

	cols, err := f.GetCols(sheetName)
	if err != nil {
		return nil, err
	}
	if colNum <= 0 || colNum > len(cols) {
		return nil, fmt.Errorf("找不到指定的列")
	}

	return cols[colNum-1], nil
}

// 按照列名查找列
func FindColumnIndexByTitle(f *excelize.File, sheetName, title string) (int, error) {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return 0, err
	}

	for c, cell := range rows[0] { // 假设标题在第一行
		if cell == title {
			return c + 1, nil
		}
	}
	return 0, fmt.Errorf("找不到标题为 %s 的列", title)
}
