package main

import "fmt"

type rect struct {
	width, height int
}

// 这个area方法有一个限定类型*rect，
// 表示这个函数是定义在rect结构体上的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 方法的定义限定类型可以为结构体类型
// 也可以是结构体指针类型
// 区别在于如果限定类型是结构体指针类型
// 那么在该方法内部可以修改结构体成员信息
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("person test")
	fmt.Println("Id:", this.Id, "Name", this.Name)
	fmt.Println("x=", x)
}
type Student struct {
	Person
	Score int
}

func (this Student) test() {
	fmt.Println("student test")
}

// 这个"intSeq"函数返回另外一个在intSeq内部定义的匿名函数，
// 这个返回的匿名函数包住了变量i，从而形成了一个闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
//传参数的闭包函数
func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
/**
一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法。
 */
func main() {
	r := rect{width: 10, height: 5}

	// 调用方法
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go语言会自动识别方法调用的参数是结构体变量还是
	// 结构体指针，如果你要修改结构体内部成员值，那么使用
	// 结构体指针作为函数限定类型，也就是说参数若是结构体
	//变量，仅仅会发生值拷贝。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())


	//语法糖（类似于闭包）instance.method(args)->(type).func(instance,args)
	p := Person{2, "张三"}

	p.test(1)
	var f1 func(int) = p.test
	f1(2)
	Person.test(p, 3)
	var f2 func(Person, int) = Person.test
	f2(p, 4)


	//使用匿名字段，实现模拟继承。即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。
	// 访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果
	s := Student{Person{2, "张三"}, 25}
	s.test()


	// 我们调用intSeq函数，并且把结果赋值给一个函数nextInt，
	// 这个nextInt函数拥有自己的i变量，这个变量每次调用都被更新。
	// 这里i的初始值是由intSeq调用的时候决定的。
	nextInt := intSeq()
	// 调用几次nextInt，看看闭包的效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 为了确认闭包的状态是独立于intSeq函数的，再创建一个。
	newInts := intSeq()
	fmt.Println(newInts())

	//传参数的闭包函数
	add10 := closure(10)//其实是构造了一个加10函数
	fmt.Println(add10(5))
	fmt.Println(add10(6))
	add20 := closure(20)
	fmt.Println(add20(5))




}