package main

import (
	"image"
	"os"
	"path/filepath"

	"pptTools/libs"
	utils "pptTools/utils"

	log "github.com/sirupsen/logrus"
)

var (
	rootPath   = ""
	inputPath  = ""
	outputPath = ""
)

func init() {
	//设置日志格式
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	//初始化variable
	rootPath = utils.GetCurrentPath()
	inputPath = rootPath + "/input"
	outputPath = rootPath + "/output"
	utils.CheckErr(os.MkdirAll(outputPath, os.ModePerm))
}

func handleOneImage(fileFullName string) {
	_, fileName := filepath.Split(fileFullName)

	//image cutting
	rawImageFile, err := os.Open(fileFullName)
	utils.CheckErr(err)

	rawImage, _, err := image.Decode(rawImageFile)
	utils.CheckErr(err)

	//left
	subImage, err := libs.ImageCopy(rawImage, 0, 78, 97, 606)
	utils.CheckErr(err)

	subImagePath := outputPath + "/" + fileName + "_left.png"
	err = libs.SaveImage(subImagePath, subImage)
	utils.CheckErr(err)

	//right
	subImage, err = libs.ImageCopy(rawImage, 221, 86, 1029, 580)
	utils.CheckErr(err)

	subImagePath = outputPath + "/" + fileName + "_right.png"
	err = libs.SaveImage(subImagePath, subImage)
	utils.CheckErr(err)
}

func main() {
	inputFiles := utils.GetFileList(inputPath)
	for _, file := range inputFiles {
		log.Info("handling file {}", file)
		handleOneImage(file)
	}
	log.Info("completed")
}
