package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileIntoStringSlice(path string) []string {
	file, _ := os.Open(path)
	bytes, _ := ioutil.ReadAll(file)
	str := string(bytes)
	slice := strings.Split(str, "\n")
	return slice
}
