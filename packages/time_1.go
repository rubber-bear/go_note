package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	t1 := time.Now()

	// 格式化处理,将time类型转换为字符串 %Y %m %d %H %M %S
	dataString := t1.Format("2006_01_02_15_04_05")
	fmt.Println(dataString)

	// 创建文件夹
	// os.Mkdir(dataString, 0755)

	// 2- 获取当前UTC时间 t2.Local() 得到一个cst时间
	t2 := time.Now().UTC()
	fmt.Println("当前UTC时间（Time类型）：", t2)

	// 3: 创建一个时间，字符串类型 -> Time类型
	t3, _ := time.Parse("2006-01-02", "2023-01-24")
	fmt.Println("根据字符串转换为时间（Time类型）", t3)

	// 创建一个时间
	t4 := time.Date(2023, 1, 24, 17, 28, 19, 18, time.Local)
	fmt.Println("根据字符串转换为时间（Time类型）", t4)

	t5 := time.Date(2023, 1, 24, 17, 28, 19, 18, time.UTC)
	fmt.Println("根据字符串转换为时间（Time类型）", t5)

	// 时间格式化  Time类型 - > 字符串类型
	fmt.Println("格式化之后的时间（string类型）", t1.Format("2006-01-02"))

	// 时间增加
	t6 := t1.Add(time.Hour * 1)
	fmt.Println("当前时间+1小时（Time类型）", t6)

	// 时间减小
	t7 := t1.Add(-time.Hour * 1)
	fmt.Println("当前时间+1小时（Time类型）", t7)
	
	// 时间间隔
	t8 := t1.Sub(t4)
	fmt.Println("时间间隔为（Duration类型）", t8)

}
