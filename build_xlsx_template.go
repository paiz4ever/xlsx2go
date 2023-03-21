package main

import (
	"io"
	"log"
	"text/template"
)

func GenerateTemplate(outputFile io.Writer, templateData TemplateData) {
	outputTemplate := `package {{.PakageName}}

import (
	"log"
	"strconv"
	"github.com/xuri/excelize/v2"
)

func init() {
	f, err := excelize.OpenFile("xlsx/{{.FileName}}.xlsx")
	if err != nil {
		log.Println("xlsx/{{.FileName}}.xlsx异常: ", err)
		return
	}
	defer f.Close()
	
	{{range .StructInfos}}init{{.Name}}(f)
	{{end}}
}

{{range .StructInfos}}
{{$cName := getCfgName $.FileName .Name}}
var {{$cName}} = make(map[{{with index .Fields 0}}{{.Type}}{{end}}]*{{.Name}})

func init{{.Name}}(f *excelize.File) {
	rows, err := f.GetRows("{{.SheetName}}")
	if err != nil {
		log.Println("xlsx/{{$.FileName}}.xlsx获取sheet<{{.SheetName}}>异常: ", err)
		return
	}
	for _, row := range rows[2:] {
		data := &{{.Name}}{}
		{{range .Fields}}
		{{if eq .Type "string"}}data.{{.Name}} = row[{{.Index}}]
		{{else if eq .Type "int"}}
		v{{.Index}}, err := strconv.Atoi(row[{{.Index}}])
		if err != nil {
			log.Println("类型转换错误: ", err)
			return
		}
		data.{{.Name}} = v{{.Index}}{{else if eq .Type "bool"}}
		v{{.Index}}, err := strconv.ParseBool(row[{{.Index}}])
		if err != nil {
			log.Println("类型转换错误: ", err)
			return
		}
		data.{{.Name}} = v{{.Index}}{{end}}{{end}}
		{{with index .Fields 0}}{{$cName}}[data.{{.Name}}] = data{{end}}
	}
}

type {{.Name}} struct {
	{{range .Fields}}{{.Name}} {{.Type}}
	{{end}}
}

func (c *{{.Name}}) GetData(keys ...{{with index .Fields 0}}{{.Type}}{{end}}) []*{{.Name}} {
	datas := make([]*{{.Name}}, 0)
	for _, key := range keys {
		datas = append(datas, {{$cName}}[key])
	}
	return datas
}
{{end}}
	`

	tmpl, err := template.New("output").Funcs(template.FuncMap{
		"getCfgName": func(fname string, sname string) string {
			return fname + sname + "Cfgs"
		},
	}).Parse(outputTemplate)
	if err != nil {
		log.Println("模板解析失败: ", err)
		return
	}

	err = tmpl.Execute(outputFile, templateData)
	if err != nil {
		log.Println("模板导出失败: ", err)
		return
	}
}
