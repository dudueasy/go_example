package main

import (
	"fmt"
)

func main() {
	// 数组值默认会初始化为0
	var arr [10]int
	fmt.Printf("%v\n", arr)

	// 部分初始化
	var arr2 [10]int = [10]int{1, 3, 4}
	fmt.Printf("%v\n", arr2)

	// 省略写法 - 快速初始化数组
	arr3 := [3]int{1, 3, 5}
	fmt.Printf("%v\n", arr3)
	// 访问数组成员
	fmt.Printf("%v\n", arr3[1])
	// 获取数组长度
	fmt.Printf("%v\n", len(arr3))
	// 遍历数组
	fmt.Printf("%v\n", "for loop")
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("%v\n", arr3[i])
	}

	// 遍历数组简便写法
	fmt.Printf("%v\n", "for loop shortcuts")
	for i, v := range arr3 {
		fmt.Printf("%v\n", i)
		fmt.Printf("%v\n", v)
	}

	// 修改数组元素的值.
	fmt.Printf("%v\n", arr3)
	changeArrItem(&arr3)
	fmt.Printf("%v\n", arr3)

	// 声明一个多元数组
	var arr4 = [4][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Printf("%v\n", arr4)

	// 数组常见问题:
	//1. 数组值不能越界
	//var arr5 = [3]int{1, 2, 3, 4, 5}
	//fmt.Printf("%v\n", arr5)

	//2. 数组长度是类型的一部分
	arr6 := [3]int{}
	arr7 := [4]int{1, 2, 3, 4}

	arr6 = arr7

}

func changeArrItem(arr *[3]int) {
	// 此处的 arr 自动解引用了
	for i := range arr {
		// (*arr)[i] += 1
		arr[i] += 1
	}
}
