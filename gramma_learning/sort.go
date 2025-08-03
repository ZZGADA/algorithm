package gramma_learning

import (
	"fmt"
	"sort"
)

// Person 定义一个自定义结构体
type Person struct {
	Name string
	Age  int
}

// ByAge 定义一个Person切片类型
type ByAge []Person

// 实现sort.Interface接口

// Len 方法返回切片的长度
func (a ByAge) Len() int { return len(a) }

// Swap 交换
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less 方法用于排序  按年龄升序
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func SortSelfDefine() {
	people := []Person{
		{"Alice", 31},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// 按年龄排序
	// 定义一个接口类型
	sort.Sort(ByAge(people))
	fmt.Println("按年龄排序:")

	for _, p := range people {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
}
