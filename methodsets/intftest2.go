package methodsets

import (
	"fmt"
	"reflect"
)

func init() {

}

type InrfTestVal interface {
	IntfValTest(int)
}

type InrfTestPtr interface {
	IntfPtrTest(int)
}

type IntfTestType struct {
	C1 int
	C2 int
}

func (it IntfTestType) IntfValTest(int) { fmt.Printf("IntfValTest\n") }

func (it *IntfTestType) IntfPtrTest(int) { fmt.Printf("IntfPtrTest\n") }

func IntfTest2() {

	var val IntfTestType
	ptr := new(IntfTestType)

	var n1 InrfTestVal
	n1 = val
	fmt.Printf("Type of InrfTestVal is %s", reflect.TypeOf(n1))
	var n2 InrfTestVal
	n2 = ptr
	fmt.Printf("Type of InrfTestVal is %s", reflect.TypeOf(n2))

	// var n3 InrfTestPtr
	// n3 = val
	// fmt.Printf("Type of InrfTestVal is %s", reflect.TypeOf(n3))
	var n4 InrfTestPtr
	n4 = ptr
	fmt.Printf("Type of InrfTestVal is %s", reflect.TypeOf(n4))

}
