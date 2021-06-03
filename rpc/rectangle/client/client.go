/*
 * Filename: c:\Users\kdjlyy\Desktop\codes\go\Go\rpc\rectangle\client\client.go
 * Path: c:\Users\kdjlyy\Desktop\codes\go\Go\rpc\rectangle\client
 * Created Date: Thursday, June 3rd 2021, 10:41:33 am
 * Author: kdjlyy
 *
 * Copyright (c) 2021 -
 */

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 传的参数
type Params struct {
	Width, Height int
}

func main() {
	// 1.连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	// 2.调用方法
	// 面积
	retArea := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &retArea)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Area:", retArea)

	// 周长
	retPerimeter := 0
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &retPerimeter)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("Perimeter:", retPerimeter)
}
