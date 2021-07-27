package main

import (
	"os"
	"strconv"

	"pptTools/service"
	utils "pptTools/utils"

	log "github.com/sirupsen/logrus"
)

var (
	rootPath   = ""
	inputPath  = ""
	outputPath = ""
	finalPath  = ""
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
	finalPath = rootPath + "/output_final"
	utils.CheckErr(os.MkdirAll(outputPath, os.ModePerm))
	utils.CheckErr(os.MkdirAll(finalPath, os.ModePerm))
}

func main() {
	//cut
	inputFiles := utils.GetFileList(inputPath)
	for _, file := range inputFiles {
		log.Info("cutting file :", file)
		service.CutOneImage(file, outputPath)
	}

	//merge
	for i := 1; i <= 20; i++ {
		left := outputPath + "/" + strconv.Itoa((i-1)*2+1) + "_left.png"
		rightTop := outputPath + "/" + strconv.Itoa((i-1)*2+1) + "_right.png"
		rightBottom := outputPath + "/" + strconv.Itoa((i-1)*2+2) + "_right.png"
		newImagePath := finalPath + "/" + strconv.Itoa(i) + "_final.png"
		//check file exist
		if !utils.PathExists(left) || !utils.PathExists(rightTop) || !utils.PathExists(rightBottom) {
			break
		}
		log.Info("merging file :", left)
		service.MergeImage(left, rightTop, rightBottom, newImagePath)
	}

	log.Info("completed")
}
