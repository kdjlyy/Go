/*
 * Filename: c:\Users\kdjlyy\Desktop\codes\go\Go\rpc\rectangle\server\server.go
 * Path: c:\Users\kdjlyy\Desktop\codes\go\Go\rpc\rectangle\server
 * Created Date: Wednesday, June 2nd 2021, 7:51:43 pm
 * Author: kdjlyy
 *
 * Copyright (c) 2021 -
 */

package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// golang实现RPC程序，实现求矩形面积和周长

type Params struct {
	Width, Height int
}

// 服务器端需要注册结构体对象，然后通过对象所属的方法暴露给调用者，从而提供服务，
// 该方法称之为输出方法，此输出方法可以被远程调用
type Rect struct{}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

// 主函数
func main() {
	// 1.注册服务
	rect := new(Rect)
	// 注册一个rect的服务
	rpc.Register(rect)

	//2.服务器处理绑定到http协议上
	rpc.HandleHTTP()

	//3.监听服务
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Panicln(err)
	}
}

