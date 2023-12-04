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
