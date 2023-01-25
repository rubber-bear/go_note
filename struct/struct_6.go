package main

import "fmt"

// 在存在结构体嵌套时，赋值会拷贝一份所有的数据
//todo: 对于那些默认拷贝的情况，可以改变为指针类型，让数据实现同步修改

type Address4 struct {
	city, state string
}

type Person4 struct {
	name string
	age  int
	Address4
}

type Person5 struct {
	name   string
	age    int
	hobby  [2]string // 拷贝
	hobby2 *[2]string
	num    []int             // 未拷贝
	parent map[string]string // 未拷贝
	Address4
}

func main() {
	a1 := Person4{name: "xx", age: 100, Address4: Address4{"tj", "zg"}}
	a2 := a1 // 嵌套的结构体会拷贝一份
	fmt.Println(a1.Address4, a2.Address4)

	a1.city = "nj"
	fmt.Println(a1.Address4, a2.Address4)

	// 本质上赋值都会拷贝一份，由于切片和map 内部维护的是指针，指向数据存储的位置

	c1 := Person5{"mm", 15, [2]string{"hj", "cy"}, &[2]string{"dj", "dlp"},
		[]int{11, 22, 33, 44}, map[string]string{"f": "a", "m": "b"}, Address4{"bj", "zg"}}
	c2 := c1
	fmt.Println(c1, c2)

	c1.hobby[0] = "dbj"
	c1.hobby2[0] = "pc"
	c1.num[0] = 999
	c1.parent["f"] = "lw"
	fmt.Println(c1, c2)
	fmt.Println(c1.hobby2, c2.hobby2)

}
