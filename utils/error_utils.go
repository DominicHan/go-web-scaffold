package utils

//package main

/*
配置文件形如：
{
  "1":{
    "code": "1",
    "message": "fail"
  },
  "2":{
    "code": "2",
    "message": "success"
  }
}
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var ErrJson = make(map[string]map[string]string)

func readFile(filename string) (map[string]map[string]string, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &ErrJson); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return ErrJson, nil
}

func GetErrorMap(errorCode string) (oneMap map[string]string, oneErr error) {
	var resultMap = map[string]string{}
	errMap, err := readFile("/Users/dominichan/go/src/go-web-scaffold/tools/xxx.json")
	if err != nil {
		return resultMap, err
	}

	value, isOk := errMap[errorCode]

	if isOk {
		resultMap["message"] = value["message"]
		resultMap["code"] = value["code"]
	} else {
		resultMap["message"] = "error unknow"
		resultMap["code"] = "9999"
	}
	return resultMap, nil
}

//func main(){
//	aa, err := GetErrorMap("1")
//	if err == nil{
//		fmt.Println(aa)
//	}
//
//	bb, err := GetErrorMap("6")
//	if err == nil{
//		fmt.Println(bb)
//	}
//}
