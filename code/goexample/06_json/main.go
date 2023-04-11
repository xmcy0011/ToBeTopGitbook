package main

import (
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"
)

// 嵌套一个对象
type Info struct {
	Age int `json:"age"`
}

var info = Info{Age: 12}

// 嵌套一个对象数组
type Extra struct {
	Address string `json:"address"`
}

// 定义需要反序列化的结构体
type UserRequest struct {
	Name     string  `json:"userName"`  // 通过tag里面的json，来指定json字符串中该字段的值从那里解析，不需要和字段名一样
	NickName string  `json:"nick_name"` // 如果没对应上，解析不了
	info     Info    `json:"info"`      // 小写私有的，故反序列化失效，该字段永远为空
	Extra    []Extra `json:"extra"`
}

func main() {
	jsonStr := `
	{
		"userName":"admin",
		"nick_name":"管理员",
		"info":{
		   "age":18
		},
		"extra":[
		   {
			  "address":"上海市"
		   },
		   {
			  "address":"北京市"
		   }
		]
	 }`

	// 方式一：序列化成map（不常用）
	anyMap := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(jsonStr), &anyMap); err != nil {
		panic(err)
	}
	log.Println("Unmarshal to map result:", anyMap)

	// 方式二：序列化成对象，经常使用
	req := UserRequest{}
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		panic(err)
	}
	log.Println("Unmarshal to struct:", req)

	// 方式三：不反序列化，只读取单个key，经常使用。适合特别复杂的json字符串，或者有多种if else结构的场景
	userName := gjson.Get(jsonStr, "userName")
	nickName := gjson.Get(jsonStr, "nick_name")
	age := gjson.Get(jsonStr, "info.age").Int()

	// 取得extra数组0位置的对象
	address1 := gjson.Get(jsonStr, "extra").Array()[1]
	log.Println("get raw value by key:", userName, nickName, age, address1.Get("address"))
}
