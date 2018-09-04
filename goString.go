package main

import s "strings"
import "fmt"

// 这里给fmt.Println起个别名，因为下面我们会多处使用。
var p = fmt.Println
//https://www.jianshu.com/p/c62ebccca79d
func main() {

	// 下面是strings包里面提供的一些函数实例。注意这里的函数并不是
	// string对象所拥有的方法，这就是说使用这些字符串操作函数的时候
	// 你必须将字符串对象作为第一个参数传递进去。
	p("Contains:  ", s.Contains("test", "es"))//判断是否包含，返回true false 如果 substr 为空，则返回 true

	//这个是查询字符串中是否包含多个字符.
	p("Contains:  ", s.ContainsAny("test", "as"))//ContainsAny 判断字符串 test 中是否包含 as 中的任何一个字符，如果 chars 为空，则返回 false


	//rune中只能是单引号并且只有一个string
	p("Contains:  ", s.ContainsRune("test", rune('我')))//这里边当然是字符串中是否包含rune类型，其中rune类型是utf8.RUneCountString可以完整表示全部Unicode字符的类型
	p("Contains:  ", s.ContainsRune("我中国", 99))//这里边当然是字符串中是否包含rune类型，其中rune类型是utf8.RUneCountString可以完整表示全部Unicode字符的类型


	p("Count:     ", s.Count("test", "t"))//获取存在的个数，返回int
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))//这个函数是查找字符串，然后返回当前的位置，输入的都是string类型，然后int的位置信息。
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// 你可以在strings包里面找到更多的函数

	// 这里还有两个字符串操作方法，它们虽然不是strings包里面的函数，
	// 但是还是值得提一下。一个是获取字符串长度，另外一个是从字符串中
	// 获取指定索引的字符
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}