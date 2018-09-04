package main

import "fmt"
/**
Golang 的数组指针和指针数组
 */
func main(){
	x,y := 1, 2
	var arr3 = [3]int {1,2,3}
	fmt.Println(arr3)
	var arr =  [...]int{5:2}
	for i, v := range arr {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
	var c = [...]int{3: 4, 2: 3, 1: 2}
	for i, v := range c {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
	fmt.Println(arr)
	fmt.Println(c)
	//数组指针
	var pf *[6]int = &arr
	//指针数组
	pfArr := [...]*int{&x,&y}
	fmt.Println(pf)
	fmt.Println(pfArr)
}