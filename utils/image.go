package utils

import "math"

func CalculateRatioFit(srcWidth, srcHeight int, defaultWidth, defaultHeight float64) (int, int) {
	ratio := math.Min(defaultWidth/float64(srcWidth), defaultHeight/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}
