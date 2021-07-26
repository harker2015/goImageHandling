package utils

import (
	"io/ioutil"
)

func GetFileList(path string) []string {
	fs, err := ioutil.ReadDir(path)
	CheckErr(err)
	result := make([]string, 0)
	for _, file := range fs {
		if file.IsDir() {
			continue
		} else {
			result = append(result, path+"/"+file.Name())
		}
	}
	return result
}
