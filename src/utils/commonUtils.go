package utils

import (
	"fmt"
	"os"
)

func TestPrint() {
	fmt.Printf("test print\n")
}

func Version() string {
	return "0.0.1"
}

func IsFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ParasInt(val interface{}) int {
	return int(val.(float64))
}
