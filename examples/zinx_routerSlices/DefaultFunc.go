package main

import (
	"fmt"
	"time"

	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/znet"
)

func DefaultTest1(request ziface.IRequest) {
	fmt.Println("test1")
}
func DefaultTest2(request ziface.IRequest) {
	time.Sleep(1)
	arr := make([]int, 1)
	fmt.Println(arr[1])
}

func main() {
	s := znet.NewDefaultRouterSlicesServer()
	s.AddRouterSlices(1, DefaultTest1, DefaultTest2)
	s.Serve()
}
