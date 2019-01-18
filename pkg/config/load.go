package config

/**
加载配置文件
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("load读取配置文件出错--->", filename)
		panic(err)
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("load序列化配置文件出错--->", filename)
		panic(err)
	}
}
