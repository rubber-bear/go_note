package main

import (
	"encoding/json"
	"fmt"
)

// json包主要用于json数据的序列化和反序列化
//	Marshal， 序列化， Go -> json 格式
//  Unmarshal， 反序列化

// 结构体的字段【首字母必须大写】, 不然序列化时候读取不到
// 想让序列化后的字段为小写，可以在结构体中写tag实现

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	v1 := []interface{}{
		"hany",
		123,
		true,
		3.14,
		map[string]interface{}{"name": "xll", "age": 19},
		Info{"hanL", 555},
	}

	res, _ := json.Marshal(v1)
	data := string(res)
	fmt.Println(data) //["hany",123,true,3.14,{"age":19,"name":"xll"}]

	content := `["hany",123,true,3.14,{"age":19,"name":"xll"},{"name":"hanL","age":555}]`
	var value []interface{}
	json.Unmarshal([]byte(content), &value)
	fmt.Println(value)

}
