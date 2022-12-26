package main

import (
	"fmt"
	"math"
)

func MathOperation() {
	fmt.Println(math.Abs(-19))                // 取绝对值
	fmt.Println(math.Floor(3.145))            // 向下取整
	fmt.Println(math.Ceil(3.123))             // 向上取整
	fmt.Println(math.Round(3.3463))           // 就近取整
	fmt.Println(math.Round(3.5478*100) / 100) // 保留两位小数
	fmt.Println(math.Mod(11, 3))              // 取余数 同 11 %3
	fmt.Println(math.Pow(2, 5))               // 计算次方 2的5次方
	fmt.Println(math.Pow10(2))                // 计算10次方， 10的2次方
	fmt.Println(math.Max(2, 5))               // 取最大值
	fmt.Println(math.Min(2, 5))               // 取最小值
}
