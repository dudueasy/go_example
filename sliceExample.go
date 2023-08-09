package main

import "fmt"

func main() {
	// 初始化 slice
	var s []int
	fmt.Printf("%v\n", s)

	// 给 slice 扩容
	//for i := 0; i < 10; i++ {
	//	s = append(s, i)
	//}
	appendChild(&s)
	fmt.Printf("s: %v\n", s)

	// slice 的长度和容量
	sliceLen := len(s)
	sliceCap := cap(s)
	fmt.Printf("%v\n", sliceLen)
	fmt.Printf("%v\n", sliceCap)

	// 定义一个确定容量的slice
	s2 := make([]int, 0, 100)
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
	}
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", len(s2))
	fmt.Printf("%v\n", cap(s2))

	// 取数组元素
	var firstItem = s2[0]
	fmt.Printf("firstItem : %v\n", firstItem)

	var secondToLast = s2[1:]
	fmt.Printf("secondToLast : %v\n", secondToLast)

	var firstThreeItem = s2[:3]
	fmt.Printf("firstThreeItem: %v\n", firstThreeItem)

	var secondAndThree = s2[1:3]
	fmt.Printf("secondAndThree: %v\n", secondAndThree)

	// 修改切片的item值
	changeSliceValue(s2)
	fmt.Printf("s2: %v\n", s2)

	// 通过函数给切片扩容
	appendChild(&s2)
	fmt.Printf("s2: %v\n", s2)

	// 切片删除元素.
	// 删除第0项
	s2 = append(s2[:1], s[2:]...)
	fmt.Printf("s2: %v\n", s2)

	// 数组和 slice 的转化
	// 创建数组
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("arr: %v\n", arr)
	// 创建 slice
	s3 := arr[:]
	fmt.Printf("s3 : %v\n", s3)

}

func changeSliceValue(arr []int) {
	for i := range arr {
		arr[i] *= 2
	}
}

// 如果要 "通过函数" 给切片扩容, 需要传指针
func appendChild(arr *[]int) {
	for i := 0; i < 10; i++ {
		*arr = append(*arr, i)
	}
}
