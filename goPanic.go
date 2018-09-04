/**
Panic表示的意思就是有些意想不到的错误发生了。通常我们用来表示程序正常运行过程中不应该出现的，或者我们没有处理好的错误。
 */
package main

import "os"

func main() {

	// 我们使用panic来检查预期不到的错误
	panic("a problem")

	// Panic的通常使用方法就是如果一个函数
	// 返回一个我们不知道怎么处理的错误的
	// 时候，直接终止执行。
	_, err := os.Create("/Users/zhiweizhang/Desktop/defer.txt")
	if err != nil {
		panic(err)
	}
}
//和其他的编程语言不同的是，Go并不使用exception来处理错误，而是通过函数返回值返回错误代码。