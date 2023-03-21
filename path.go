package main

import (
	"os"
	"path/filepath"
)

func GetExcelInputDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(wd, buildConfig.Xlsx.Input)
}

func GetExcelOutputDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(wd, buildConfig.Xlsx.Output)
}
