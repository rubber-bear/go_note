package main

import "fmt"

type Man struct {
	name string
	age  int
}

func main() {

	/*
		 			-------------
			m1 -- > | wang |  19 |          结构体类型
				     -------------
					内存地址为 0xc000a60200

					-------------
			m2 -- > | wang |  19 |          结构体类型
					 -------------
				   内存地址为 xc000a6040

	*/
	m1 := Man{"wang", 19}
	m2 := m1            // 内部将m1重新拷贝一份
	fmt.Println(m1, m2) // {wang 19} {wang 19}

	m1.name = "jerry"
	fmt.Println(m1, m2) // {jerry 19} {wang 19}

	/*
		                 -------------
			             | han | 20 |
						 -------------
					内存地址：xc0000a60200
								^
								|
						 -------------
			m3 ----> 	|oxc0000a60200|   结构体指针
						 -------------
						内存地址： xc000a640


		                 -------------
			m4 ----> 	|oxc0000a60200|   结构体指针
						 -------------
						内存地址： xc0076700788
	*/

	m3 := &Man{"han", 20}
	m4 := m3

	fmt.Println(m3, m4)

	m3.name = "ww"
	fmt.Println(m3, m4)

}
