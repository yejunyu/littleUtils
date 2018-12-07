package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func changeAllName(pathRoot string, regString string) error {
	rd, err := ioutil.ReadDir(pathRoot)
	for _, fi := range rd {
		if fi.IsDir() {
			path := pathRoot + string(filepath.Separator) + fi.Name() + string(filepath.Separator)
			fmt.Println(path)
			changeAllName(path, regString)
		} else {
			filePath := pathRoot + string(filepath.Separator) + fi.Name()
			newPath := repalce(filePath, regString)
			os.Rename(filePath, newPath)
		}
	}
	return err
}

func repalce(filePath string, regString string) string {
	re := strings.Replace(strings.Replace(regString, "[", "\\[", 1), "]", "\\]", 1)
	re1, _ := regexp.Compile(re)
	rep := re1.ReplaceAllString(filePath, "")
	return rep
}

func main() {
	pathRoot := flag.String("path", "E:\\BaiduNetdiskDownload\\深入理解Java虚拟机", "文件夹目录的绝对路径")
	regString := flag.String("str", "[.*?]", "要去掉的字符串")
	// 需要 parse 来使生效
	flag.Parse()
	//遍历打印所有的文件名
	changeAllName(*pathRoot, *regString)
}
