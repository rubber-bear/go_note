package main

/*
	接口是一种数据类型，Go的空接口可以代指任意类型，从而实现参数、"容器"中可以处理多种数据类型

	type eface struct {
		_type * _type
		data unsafe.Pointer
	}

	如果在代码中出现其他对象赋值给空接口，其实就是将其他对象相关的值放到eface的_type 和data中

	_type 是一个结构体，内部存储挺多信息，这里统称为类型相关的信息

	nums:= 666                                       eface
	或：                          ---------->         类型相关     type
	info =Person({name:"h", age: 19})                数据        data




	todo： 对于非空接口情况比较复杂
	type iface struct {
		tab * itab           // 类型和方法相关
		data unsafe.Pointer  // 数据
	}

	type itab struct {
		inter *interfacetype  // 接口信息，如：接口中定义的方法
		_type * _type         // 类型
		hash  uint32
		_   [4]byte
		fun [1]uintptr
	}

	type interfacetype struct {
		typ  _type
		pkgpath name
		mhdr   []imethod   // 接口方法
	}


	nums:= 666                         iface
		或：           ---------->     类型和方法  tab  ---> itab
	info =Person({name:"h"})           数据      data      接口信息  inter   --->   interfacetype
															类型   _type           方法集   mhdr
                                                            -                      -


*/

func main() {
	num := 666
	var _ interface{}
	// 将num的类型int存储到 _type中， 值存储到data中
	_ = num
}
