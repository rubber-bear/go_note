package main

import "fmt"

func StringLoop() {

	var name = "叶安世"

	// 索引获取字节
	v1 := name[0]
	fmt.Println(v1) // 229

	// 2.切片获取字节区间
	v2 := name[0:3]
	fmt.Println(v2) // 叶

	// 3. 循环获取所有字节
	/*
		0 229
		1 143
		2 182
		3 229
		4 174
		5 137
		6 228

	*/
	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])
	}

	// 4. for range 循环获取所有字符
	/*
		0 叶
		3 安
		6 世
	*/

	for index, item := range name {
		fmt.Println(index, string(item))
	}

	// 5.转换成rune集合
	dataList := []rune(name)                      // 将每一个字符放到集合里
	fmt.Println(dataList[0], string(dataList[0])) //21494 叶

}
