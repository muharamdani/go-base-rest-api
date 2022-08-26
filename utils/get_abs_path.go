package utils

import (
	"os"
	"regexp"
)

var projectDirName = "go-base-rest-api"

func GetRootPath() string {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return string(rootPath) + "/"
}
