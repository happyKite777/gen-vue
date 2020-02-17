package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonConf struct {
	NAME string `json:"NAME"`
	FILTER string `json:"FILTER"`
	PARAM []string `json:"PARAM"`
	COLUMNS []Column `json:"COLUMNS"`
}

type Column struct {
	Label string `json:"label"`
	Prop string `json:"prop""`
}

func ParseJson()  {
	fh, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}

	var post JsonConf
	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}

