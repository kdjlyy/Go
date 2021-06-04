package main

import "fmt"

// 首字母大写，代表对外部可见，首字母小写代表对外部不可见，适用于所有对象，包括函数、方法

// 接口是一种类型，类似struct，里面存放方法的声明
type Action interface {
	// 方法名() 返回类型
	run() int
	eat() int
}

// 空接口，可以接收任意类型，故可以赋值给string, struct等类型
type NullInterface interface{}

type Human struct {
	weight float64
	name   string
}

func (h Human) run() int {
	fmt.Println("run run run")
	return 0
}
func (h Human) eat() int {
	fmt.Println("eat eat eat")
	return 0
}

func main() {
	// 使用接口
	var n NullInterface = 1.50
	fmt.Println(n)

	var act Action = Human{120.0, "Bob"}
	fmt.Println(act)
	act.run()
	act.eat()
}
