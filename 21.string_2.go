package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func StringFunc() {
	var place = "H合HHH生汇hH"
	// 1.获取长度
	fmt.Println(len(place))
	fmt.Println(utf8.RuneCountInString(place))

	// 2.是否以xx开头
	res1 := strings.HasPrefix(place, "和")
	fmt.Println(res1)

	// 3.是否以xx结尾
	res2 := strings.HasSuffix(place, "汇")
	fmt.Println(res2)

	// 4.是否包含
	res3 := strings.Contains(place, "生")
	fmt.Println(res3)

	// 5.大写
	res5 := strings.ToUpper(place)
	// 注意res5是变大写，place依然是小写
	fmt.Println(res5)

	// 6.小写
	res6 := strings.ToLower(place)
	fmt.Println(res6)

	// 7. 去两边
	res71 := strings.TrimRight(place, "h")
	res72 := strings.TrimLeft(place, "H")
	res73 := strings.Trim(place, "H")
	fmt.Println(res71, res72, res73)

	// 8. 替换
	res81 := strings.Replace(place, "H", "X", 1)
	res82 := strings.Replace(place, "H", "X", -1)

	fmt.Println(res81, res82)

	// 9, 分割
	name := "纵有疾风起，人生不言弃"
	result9 := strings.Split(name, ",")
	fmt.Println(result9)

	// 10.拼接， 可以使用 + 对两个字符串进行拼接，但是效率比较低，不建议使用

	dataList := []string{"我爱", "北京天安门"}
	result10 := strings.Join(dataList, "-")
	fmt.Println(result10)

	//  拼接 - 效率更高的方式
	var buffer bytes.Buffer
	buffer.WriteString("你是")
	buffer.WriteString("个")
	buffer.WriteString("什么人")
	data := buffer.String()
	fmt.Println(data)

	//更高效的方式
	var builder strings.Builder
	builder.WriteString("哈哈哈")
	builder.WriteString("你是谁")
	value := builder.String()
	fmt.Println(value)

	// 11. string 转化为int
	num := "666"
	data11, _ := strconv.Atoi(num)
	fmt.Println(data11, reflect.TypeOf(data11))

	// 12. int 转换为string
	res := strconv.Itoa(999)
	fmt.Println(res, reflect.TypeOf(res))

	// 15.string 和字符
	v1 := string(65)
	fmt.Println(v1)
	v2, size := utf8.DecodeRuneInString("A")
	fmt.Println(v2, size)

}
