package foo

import "fmt"

func Foo() int {
	return 0
}

func init() {
	fmt.Printf("Foo init func.\n")
}
