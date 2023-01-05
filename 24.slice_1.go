package main

import "fmt"

/*
	切片是Go中重要的数据类型，每个切片对象内部都维护者：数组指针、切片长度、切片容量三个数据
	type slice struct {
		array unsafe.Pointer
		len int
		cap int
	}
                           数组（存数据）
								｜
	切片 -------->   array = 指针 len=长度 cap=容量

	在向切片中追加的数据个数大于容量时候，内部会自动扩容且每次扩容都是当前容量的两倍
    （当容量超过1024时候每次扩容只增加 1/4 容量）

*/

func SliceFunc() {
	// 创建切片的三种方式
	var nums []int
	var data = []int{11, 22, 33}

	// make 只能用于切片，字典，channel
	var users = make([]int, 2, 5) //默认初始化两个数据len=2，容量cap =5
	fmt.Println(nums, data, users)

	// 切片的指针类型 v1指向的是一块内存，指向一个切片
	var v1 = new([]int)
	// 指针类型  v2指向的是一个nil
	var v2 *[]int
	fmt.Println(v1, v2)

	// -----------------------------------------------------------------------------

	// 自动扩容
	v3 := make([]int, 1, 3)
	fmt.Println(v3, len(v3), cap(v3))

	v4 := make([]int, 3) // 容量是3，长度也是3
	fmt.Println(v4, len(v4), cap(v4))

	// 返回一个新的切片
	/*
					cap = 3 len = 1
			v3 =>   array = 0x00020303    \
		                                    ->    | 0 | 99 |  | (内存地址0x00020303)
					cap = 3 len = 2      /
			v5 =>   array = 0x00020303
	*/
	v5 := append(v3, 99)
	fmt.Println(v3, v5)

	// 需求:又一个切片，情网一个切片中追加一个数据
	v3 = append(v3, 999)
	fmt.Println(v3)

	// -----------------------------------------------------------------------------
	/*
				cap = 3 len = 3
		v6 =>   array = 0x000234323    ->    | 11 | 22 | 33 |  |   (内存地址0x000234323)

				cap = 6 len = 4
		v7 =>   array = 0x000223124    ->    | 11 | 22 | 33 | 44 |  |  |  (内存地址0x000223124)
	*/

	v6 := []int{11, 22, 33}
	v7 := append(v6, 44)

	v6[0] = 77
	fmt.Println(v6)
	fmt.Println(v7)

}
