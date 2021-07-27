package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
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

func GetFileNameWithoutExt(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
