package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("请提供要遍历的根目录和要读取的文件后缀")
		return
	}

	rootDir := os.Args[1]
	fileExt := os.Args[2]

	saveFile, err := os.Create("save.txt")
	if err != nil {
		fmt.Println("无法创建save.txt文件：", err)
		return
	}
	defer saveFile.Close()

	fileList := []string{}
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("遍历目录时发生错误：", err)
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "."+fileExt) {
			fileList = append(fileList, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("遍历目录时发生错误：", err)
		return
	}
	
	sort.Slice(fileList, func(i, j int) bool {
		count1 := strings.Count(fileList[i], ".") + strings.Count(fileList[i], "\\") + strings.Count(fileList[i], "/")
		count2 := strings.Count(fileList[j], ".") + strings.Count(fileList[j], "\\") + strings.Count(fileList[j], "/")
		if count1 == count2 {
			return fileList[i] < fileList[j]
		}
		return count1 < count2
	})

	for i := 0; i < len(fileList); i++ {
		filename := fileList[i]
		fmt.Fprintln(saveFile, "文件路径：", filename)
		fileContent, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("读取文件内容时发生错误：", err)
			return
		}
		fmt.Fprintln(saveFile, string(fileContent))
		fmt.Fprintln(saveFile)
	}

	fmt.Println("保存文件内容到save.txt完成")
}
