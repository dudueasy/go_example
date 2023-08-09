package main

import "fmt"

func main() {
	// 默认nil值
	var m1 map[string]int
	fmt.Printf("value: %v\n", m1["aa"])
	fmt.Printf("m1: %v\n", m1)

	// nil 赋值会 panic
	// m1["aa"] = 1

	// 避免nil
	m2 := make(map[string]int)
	fmt.Printf("m2: %v\n", m2)

	m2["aa"] = 123
	fmt.Printf("m2: %v\n", m2)

	// 有值初始化
	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Printf("m3: %v\n", m3)

	// 赋值和读取
	newMap := map[string]int{}
	newMap["a"] = 1
	value := m3["a"]
	fmt.Printf("value: %v\n", value)

	delete(newMap, "a")

	// 判断某个键是否存在
	val, exisit := newMap["a"]
	fmt.Printf("val: %v\n", val)
	fmt.Printf("exisit: %v\n", exisit)

	// 覆盖值, 使用 _ 代替变量名, 不会提示 unused var
	_, exisit = m3["a"]
	fmt.Printf("exisit: %v\n", exisit)

	// 判断 map 的键值对数量
	fmt.Printf("len(m3): %v\n", len(m3))

	// 迭代map
	m4 := map[string]int{}
	for i := 0; i < 5; i++ {
		m4[fmt.Sprintf("key-%d", i)] = i
	}

	for key, val := range m4 {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("val: %v\n", val)
	}

}
