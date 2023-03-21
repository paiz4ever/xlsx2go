package main

import (
	"log"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
)

func WatchFile() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln("初始化监听器异常: ", err)
	}
	defer watcher.Close()

	if err := watcher.Add(GetExcelInputDir()); err != nil {
		log.Fatalln("监听目录异常: ", err)
	}
	files := make(map[string]time.Time)
	tk := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-tk.C:
			for k := range files {
				Output(k)
				log.Println("文件重新编译完毕: ", k)
				delete(files, k)
			}
		case event := <-watcher.Events:
			// 只处理 Write
			// event.Op&fsnotify.Write == fsnotify.Write
			if path.Ext(event.Name) != ".xlsx" {
				continue
			}
			tk.Reset(time.Second * 2)
			files[event.Name] = time.Now()
			log.Println("文件变化: ", event.Name)
		case err := <-watcher.Errors:
			log.Println("文件监听错误: ", err)
		}
	}
}
