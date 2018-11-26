package main


import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kylelemons/go-gypsy/yaml"
)

func GetFilePath(dirPath string, fileName string) (filePath string, err error) {
	//返回项目文件绝对路径
	//dirPath: 文件在项目里的绝对路径 比如 run.go 是 "/IntegralMall", 项目的根目录是 "/"
	//fileName: 文件全名 "config.json"
	if dirPath[len(dirPath)-1:] != "/" {
		dirPath += "/"
	}
	currentPath := GetCurrentPath()
	filePath = currentPath + dirPath + fileName
	exists, err := PathExists(filePath)
	if exists == true {
		return filePath, nil
	} else {
		return "", err
	}

}

func GetCurrentPath() string {
	//获取当前文件所在包的地址
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func PathExists(path string) (bool, error) {
	//如果返回的错误为nil,说明文件或文件夹存在
	//如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	//如果返回的错误为其它类型,则不确定是否在存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetConfig() *yaml.File {
	configFilePath, err := GetFilePath("/", "config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := yaml.ReadFile(configFilePath)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
