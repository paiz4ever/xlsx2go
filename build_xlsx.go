package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Field struct {
	Index int
	Name  string
	Type  string
}

type StructInfo struct {
	SheetName string
	Name      string
	Fields    []Field
}

type TemplateData struct {
	StructInfos []StructInfo
	FileName    string
	PakageName  string
}

func Output(filePath string) {
	// NOTICE 这里path.Base不支持windows路径 先转成linux路径形式
	fileNameWithSuffix := path.Base(filepath.ToSlash(filePath))
	fileType := path.Ext(fileNameWithSuffix)
	fileName := strings.TrimSuffix(fileNameWithSuffix, fileType)
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println("xlsx文件读取异常: ", err)
		return
	}
	defer f.Close()

	sheets := f.GetSheetList()
	var structInfos []StructInfo = make([]StructInfo, 0)
	for _, sheet := range sheets {
		structInfo := StructInfo{
			SheetName: sheet,
			Name:      strings.Title(sheet),
			Fields:    []Field{},
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Println("读取rows异常: ", err)
			return
		}
		for i, row := range rows {
			if i == 0 {
				for idx, cell := range row {
					field := Field{
						Index: idx,
						Name:  strings.Title(cell),
					}
					structInfo.Fields = append(structInfo.Fields, field)
				}
			} else if i == 1 {
				for j, cell := range row {
					if j < len(structInfo.Fields) {
						field := &structInfo.Fields[j]
						field.Type = cell
					}
				}
			}
		}
		structInfos = append(structInfos, structInfo)
	}
	outdir := GetExcelOutputDir()
	outpath := filepath.Join(outdir, fileName+".cfg.go")
	err = os.MkdirAll(outdir, os.ModePerm)
	if err != nil {
		log.Fatalln("创建目录失败: ", err)
	}
	outputFile, err := os.Create(outpath)
	if err != nil {
		log.Println("创建编译文件异常: ", err)
		return
	}
	defer outputFile.Close()
	GenerateTemplate(outputFile, TemplateData{
		StructInfos: structInfos,
		FileName:    fileName,
		PakageName:  buildConfig.Xlsx.Package,
	})
}
