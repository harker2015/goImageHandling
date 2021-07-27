package service

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"pptTools/libs"
	"pptTools/model"
	"pptTools/utils"

	"github.com/nfnt/resize"
)

// cut one image into two parts
func CutOneImage(fileFullName string, outputPath string) {
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

// merge three images
func MergeImage(left, rightTop, rightBottom string, newImagePath string) {
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
