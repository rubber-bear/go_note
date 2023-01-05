package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
字符串
	字符串的本质
	Go语言中的字符串是utf-8 编码的序列
	unicode字符集， 文字 - > 码点 （ucs4, 4个字符表示）
	utf-8 编码，堆unicode字符集的码点进行编码

*/

func StringTest() {
	var name string = "王畅雄"
	// 王 =》11100111 10001110 10001011
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	// 获取字符串的长度，获取的是utf - 8的字节长度
	fmt.Println(len(name))

	// 字符串转换为一个字节集合
	byteSet := []byte(name)
	fmt.Println(byteSet) // [231 142 139 231 149 133 233 155 132]

	// 字节的集合转换为字符串
	byteList := []byte{231, 142, 139, 231, 149, 133, 233, 155, 132}
	targetString := string(byteList)
	fmt.Println(targetString)

	// rune 集合 将字符串转换为 unicode字符集码点的集合 （表现形式是十进制）
	tmpSet := []rune(name)
	fmt.Println(tmpSet) // [29579 30021 38596]     // 738b   7545  96c4

	fmt.Println(tmpSet[0], strconv.FormatInt(int64(tmpSet[0]), 16))
	fmt.Println(tmpSet[1], strconv.FormatInt(int64(tmpSet[1]), 16))
	fmt.Println(tmpSet[2], strconv.FormatInt(int64(tmpSet[2]), 16))

	// rune集合转换为字符串
	runeList := []rune{29579, 30021, 38596}
	targetString1 := string(runeList)
	fmt.Println(targetString1)

	// 长度的处理
	runeLength := utf8.RuneCountInString(name)
	fmt.Println(runeLength) //3

}
