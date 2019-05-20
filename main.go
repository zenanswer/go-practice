package main

import (
	"fmt"

	"github.com/zenanswer/go-practice/foo"
	"github.com/zenanswer/go-practice/methodsets"
	"github.com/zenanswer/go-practice/version"
	"github.com/zenanswer/go-practice/zaptest"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Printf("hello, world\n")
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Printf("foo: %d\n", foo.Foo())
	fmt.Printf("version: %s\n", version.Version)
	methodsets.Test()
	methodsets.IntfTest()
	zaptest.Zaptest()
	zaptest.Zaptest2()
	handlePanic()
}

func genPanic() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}

func handlePanic() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	genPanic()
	fmt.Println("e")
}
