# Golang函数、方法、接口

---

## 函数

---

函数的使用方法如下：

```go
//模板
func 函数名(参数1，……) (返回值1，……) {
	doSomething
}
//例子
func PrintInt(a int) int {
 	fmt.Println(a)
 	return 0
}
```

## 方法

---

**方法和函数的最大区别是方法有接收者(从属)**，即方法都是有主人的。方法的使用方法如下：

```go
//模板
func (主人名 类型) 方法名(参数列表) (返回值列表) {
 	doSomething
}

//构造Animal结构体，即主人类型
type Animal struct{} 
var an Animal //声明主人
an.Run(1)     //主人调用方法

//Animal类型的主人an，有一个Run方法，
//这里是值接收，也可以使用指针接收
func (an Animal) Run(a int) int {
    fmt.Println("Run ", a)
    return 0
}
```

## 接口

Go语言中虽然没有传统面向对象语言中类、集成的概念，不过提供了接口的支持，可以使用接口来使用一些面向对象的特性。

接口也是一种类型，和struct类似；区别是struct中存放各种属性，而**接口中存放各种方法的声明**。另外有几点需要注意：

1. 接口内的方法数可以为0，即空接口；默认所有对象都实现了空接口；
2. 同一个接口内所有的方法都被实现后，该接口才能被正常使用；例如下方的run和eat方法都需要被实现；
3. 建议将相同的行为放在同一个接口内

### 接口的定义

```go
type Namer interface {
    method1(param_list) return_list
    method2(param_list) return_list
    ...
}
```

这里的Namer就是接口的名字，只要符合标识符的规则即可。不过，通常约定的接口的名字最好以er, r, able结尾(视情况而定)，这样一眼就知道它是接口。

### 接口的实现

#### 🌰 1

```go
// 接口是一种类型，和struct类似，包含n个方法,其中n>=0
type Action interface{ 
    // 模板 方法名() 返回类型  
    run() int
    eat() int
}
type NullInterface interface{} // 空接口

type Human struct{
    weight float32 // 属性
    name   string
}

func (h Human) run() int { // 实现方法
    fmt.Println("run run run")
    return 0
}
func (h Human) eat() int { // 实现方法
    fmt.Println("eat eat eat")
    return 0
}

// 使用方法
// 空接口可以接受任意类型，故可以赋值string，struct等其他类型
var n nullInterface = 1 
fmt.Println("null interface", n)

/* 理解接口：
 * 接口类似工厂，我们希望生成一个工具箱（Human），它的相关参数是120.0和Bob；
 * 我们将生产需求下单给工厂Action，工厂Action生产了一个名字为act的工具箱，里
 * 面包含run和eat两种工具
*/
var act Action = Human{120.0, "Bob"}
fmt.Println("hello", act)
act.run()
act.eat()
```

**运行结果：**

```
1.5
{120 Bob}  
run run run
eat eat eat
```
