package main

import (
	"fmt"
	"reflect"
)

type Person6 struct {
	name  string   "名字"
	age   int      "年龄"
	hobby []string "爱好"
}

func StructLabel() {
	// 方式一: 先后顺序
	var p1 = Person6{"hany", 30, []string{"chi", "lc"}}
	fmt.Println(p1, p1.name, p1.age, p1.hobby)

	p1Type := reflect.TypeOf(p1)
	filed1 := p1Type.Field(0)
	fmt.Println(filed1.Tag)

	filed2, _ := p1Type.FieldByName("name")
	fmt.Println(filed2.Tag)

	filedNum := p1Type.NumField()
	for index := 0; index < filedNum; index++ {
		field := p1Type.Field(index)
		fmt.Println(field.Name, field.Tag)
	}

}
