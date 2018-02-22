package ifile

import (
	"fmt"
	"os"
)

//CheckFileIsExist is 检查目录是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(filename + " not exist")
		exist = false
	}
	return exist
}

// Mkdir is
func Mkdir(name string) {
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	fmt.Println(path)
	dir, _ := os.Getwd()                        //当前的目录
	err := os.Mkdir(dir+path+name, os.ModePerm) //在当前目录下生成目录
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("mkdir " + dir + path + name)
	}
}
