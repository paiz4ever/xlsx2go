package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

type BuildConfig struct {
	Xlsx BuildXlsx `json:"xlsx"`
}

type BuildXlsx struct {
	Input   string `json:"input"`
	Output  string `json:"output"`
	Package string `json:"package"`
}

var buildConfig BuildConfig

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	data, err := ioutil.ReadFile("build.json")
	if err != nil {
		log.Fatalln("读取配置异常: ", err)
	}
	err = json.Unmarshal(data, &buildConfig)
	if err != nil {
		log.Fatalln("解析配置异常: ", err)
	}
	log.Println("开始构建")
	dirpath := GetExcelInputDir()
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatalln("xlsx配置路径错误: ", err)
	}
	for _, file := range dir {
		// 当xlsx文件开启时会存在一个~的同名临时文件
		if strings.HasPrefix(file.Name(), "~") {
			continue
		}
		Output(filepath.Join(dirpath, file.Name()))
	}
	log.Println("构建完毕")
	log.Println("开始监听文件")
	go WatchFile()
	<-c
}
