package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func get_PWD_version1() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	//fmt.Println(exPath)
	return exPath
}

func get_PWD_version__ABS_PATH() string {

	//==this way works best for giving abs directory of this files location :-)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(pwd)

	return pwd
}

func main() {

	fmt.Println(get_PWD_version1())

	//==and then version 2

	fmt.Println(get_PWD_version__ABS_PATH())

}
