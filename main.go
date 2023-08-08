package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 申明变量
	// 变量名只能使用 字母, 数字, 下划线
	// 变量第一个字符必须是 字母或下划线, 不能是数字
	// 变量名区分大小写
	// 变量名不能是关键字.
	var name = "xiaoming"
	fmt.Println(name)
	name2 := "xiaming2"
	fmt.Println(name2)

	// 基本数据类型
	var b bool = true
	//fmt.Print("result is: %t\n", g)
	fmt.Printf("%v\n", b)

	i := 100
	fmt.Println(i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%x\n", i)

	f := 3.14
	fmt.Printf("%v\n", f)

	// 字节类型
	var g byte = 'A'
	fmt.Printf("%v\n", g)

	// 类型转换
	// int => float
	var A = 100
	var B float32
	B = float32(A)
	println(B)

	// float => int
	var C int = 10
	var D int32
	D = int32(C)
	println(D)

	// string => number
	s := "100"
	//s := "100AA" will throw error
	intValue, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T \n", intValue)
	fmt.Printf("%v \n", intValue)

	var bb bool
	var ii = 123
	if ii != 0 {
		bb = true
	}
	println(bb)

}
