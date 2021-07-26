package utils

import "os"

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	CheckErr(err)
	return dir
}
