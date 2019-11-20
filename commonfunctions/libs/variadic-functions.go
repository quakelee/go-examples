package libs

import "fmt"

// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// [_可变参数函数_](https://zh.wikipedia.org/wiki/%E5%8F%AF%E8%AE%8A%E5%8F%83%E6%95%B8%E5%87%BD%E6%95%B8)
// 可变参数函数为参数数量为不定数个，例如fmt.Println函数即为最常用的可变参数函数
// 本例演示如何将一个slice作为可变参数函数的参数送入
// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func Sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
