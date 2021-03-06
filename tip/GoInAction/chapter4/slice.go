package main

import "fmt"

// SliceCreate 切片创建及初始化
func SliceCreate() {
	// 1. 切片初始化
	// 创建一个空切片
	var s1 []string
	// 创建并初始化一个切片，注意它和数组初始化的区别
	s2 := []string{"hello", "world"}
	// 创建一个长度为3、容量为5的切片
	s3 := make([]string, 3, 5)

	fmt.Println("========SliceCreate========")
	fmt.Printf("s1(%d): %v\n", len(s1), s1)
	fmt.Printf("s2(%d): %v\n", len(s2), s2)
	fmt.Printf("s3(%d): %v\n", len(s3), s3)
}

// SliceClone 切片克隆
func SliceClone() {
	// 定义一个切片，长度和容量均为5
	a := []int{1, 2, 3, 4, 5}

	// 以a[2]为起始地址，长度为4-2=2，将a切片复制给b
	// 注意b的容量为3，因为a的实际容量是5，减去前面两个元素就为3
	b := a[2:4]

	// 注意个基本概念，尽管b的容量为3，也就是说一直延伸到a的末尾
	// a[4] = 111 ✅ 因为a的长度为5，所以访问第五元素没问题
	// b[2] = 111 ❌ 因为b的长度为2，所以最多访问到b[1]

	// 因为a和b指向同一个数组的不同地址 👈 务必理解
	b[1] = 666         // 相当于a[2] = 666
	b = append(b, 777) // 相当于追加b[2] = 777，也相当于 a[4] = 777
	b = append(b, 888) // 相当于追加b[3] = 888，已经超出a的边界了😏

	// 返回来修改b[0]，见证诡异的时刻！
	b[0] = 555 // 按照上边的惯例，应该也相当于a[2] = 555

	// 但实际情况...a[2] == 3，并没有被修改
	// 因为b的增长超出原始容量后，会分配新内存，彻底脱离原始切片
	// 此时再访问切片b时，已经和切片a没什么关系了
	fmt.Println("========SliceClone========")
	fmt.Printf("a(%d): %v\n", len(a), a) // a(5): [1 2 3 666 777]
	fmt.Printf("b(%d): %v\n", len(b), b) // b(5): [555 666 777 888]

	// 以a[2]为起始地址，长度为3-2=1，容量为3-2=1，将切片复制给c
	c := a[2:3:3]

	// 由于切片b的长度和容量一样，一旦append，就会分配新内存，脱离a
	c = append(c, 666)
	c[0] = 444 // a[2]不再受影响

	fmt.Printf("c(%d): %v\n", len(c), c) // c(2): [444 666]
}
