package main

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"strconv"

	"pptTools/libs"
	"pptTools/model"
	utils "pptTools/utils"

	"github.com/nfnt/resize"
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

func handleOneImage(fileFullName string) {
	fileName := utils.GetFileNameWithoutExt(fileFullName)

	//image cutting
	rawImageFile, err := os.Open(fileFullName)
	utils.CheckErr(err)
	defer rawImageFile.Close()

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
	//cut
	inputFiles := utils.GetFileList(inputPath)
	for _, file := range inputFiles {
		log.Info("handling file {}", file)
		handleOneImage(file)
	}

	//merge
	for i := 1; i <= 8; i++ {
		left := outputPath + "/" + strconv.Itoa((i-1)*2+1) + "_left.png"
		rightTop := outputPath + "/" + strconv.Itoa((i-1)*2+1) + "_right.png"
		rightBottom := outputPath + "/" + strconv.Itoa((i-1)*2+2) + "_right.png"
		newImagePath := finalPath + "/" + strconv.Itoa(i) + "_final.png"
		mergeImage(left, rightTop, rightBottom, newImagePath)
	}

	log.Info("completed")
}

func mergeImage(left, rightTop, rightBottom string, newImagePath string) {
	newImage := image.NewRGBA(image.Rect(0, 0, 1200, 1200))
	blue := color.RGBA{255, 255, 255, 255}
	draw.Draw(newImage, newImage.Bounds(), &image.Uniform{blue}, image.Point{}, draw.Src)

	covers := []model.ImgParam{
		{
			FileName:  left,
			NewWidth:  144,
			NewHeight: 1164,
			Rectangle: image.Rectangle{
				Min: image.Point{18, 18},
				Max: image.Point{18 + 144, 18 + 1164},
			},
		},
		{
			FileName:  rightTop,
			NewWidth:  1006,
			NewHeight: 575,
			Rectangle: image.Rectangle{
				Min: image.Point{18 + 144 + 14, 18},
				Max: image.Point{18 + 144 + 14 + 1006, 18 + 575},
			},
		},
		{
			FileName:  rightBottom,
			NewWidth:  1006,
			NewHeight: 575,
			Rectangle: image.Rectangle{
				Min: image.Point{18 + 144 + 14, 18 + 575 + 14},
				Max: image.Point{18 + 144 + 14 + 1006, 18 + 575 + 14 + 575},
			},
		},
	}

	for _, cover := range covers {
		imgFileLeft, _ := os.Open(cover.FileName)
		defer imgFileLeft.Close()
		imgLeft, _, err := image.Decode(imgFileLeft)
		utils.CheckErr(err)
		newImgLeft := resize.Resize(uint(cover.NewWidth), uint(cover.NewHeight), imgLeft, resize.Lanczos3)
		draw.Draw(newImage, cover.Rectangle, newImgLeft, newImgLeft.Bounds().Min, draw.Src)
	}

	err := libs.SaveImage(newImagePath, newImage)
	utils.CheckErr(err)
}
