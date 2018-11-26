package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kylelemons/go-gypsy/yaml"
)

var Config *yaml.File

func ConfigInit() {
	Config = GetConfig()
}

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

func ConfigStr(arg string) string {
	env, _ := Config.Get("env")
	path := "base." + env + "." + arg
	configStr, err := Config.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	return configStr
}

func ConfigInt(arg string) int64 {
	env, _ := Config.Get("env")
	path := "base." + env + "." + arg
	configInt, err := Config.GetInt(path)
	if err != nil {
		log.Fatal(err)
	}
	return configInt
}

func ConfigBool(arg string) bool {
	env, _ := Config.Get("env")
	path := "base." + env + "." + arg
	configBool, err := Config.GetBool(path)
	if err != nil {
		log.Fatal(err)
	}
	return configBool
}

func ReadFile(filePath string, reType string) (interface{}, error) {
	/*
		返回 读取文件的字节流或者字符串,注意 数据是 interface{}类型
		filePath: 文件的绝对路径 比如 "/Users/wy/GoProjects/go/src/IntegralMall/config.json"
		reType: "str" 返回字符串, 其他 返回字节流 []byte
	*/
	data, err := ioutil.ReadFile(filePath)
	if reType != "str" {
		return data, err
	}
	return string(data), err
}

func ReadFileBytes(filePath string) ([]byte, error) {
	/*
		返回 读取文件的字节流
		filePath: 文件的绝对路径 比如 "/Users/wy/GoProjects/go/src/IntegralMall/config.json"
	*/
	data, err := ioutil.ReadFile(filePath)
	return data, err
}

func ReadFileStr(filePath string) (string, error) {
	/*
		返回 读取文件的字符串
		filePath: 文件的绝对路径 比如 "/Users/wy/GoProjects/go/src/IntegralMall/config.json"
	*/
	data, err := ioutil.ReadFile(filePath)
	return string(data), err
}

func testRead(filePath string) []byte {
	/*
		测试文件可读,打印读出来的内容
	*/
	fp, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1000)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:n]))
	return data[:n]
}
