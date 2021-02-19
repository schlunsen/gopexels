package helpers

import (
	"fmt"
	"os"
)

func CreateDirIfNotExists(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		os.Mkdir(dirName, 0755)
	}
	fmt.Println("Create dir")
}
