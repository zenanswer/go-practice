package methodsets

import "fmt"

func init() {

}

type TestType struct {
	C1 int
	C2 int
}

func (t TestType) ValTest()  { fmt.Printf("ValTest\n") }
func (t *TestType) PtrTest() { fmt.Printf("PtrTest\n") }

func ValTest2(t TestType)  { fmt.Printf("ValTest\n") }
func PtrTest2(t *TestType) { fmt.Printf("PtrTest\n") }

func Test() {
	// A bare value
	var lst List
	lst.Append(1)
	(&lst).Append(2)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len())

	// A pointer value
	plst := new(List)
	plst.Append(1)
	fmt.Printf("%v (len: %d)\n", plst, plst.Len())
	fmt.Printf("%v (len: %d)\n", plst, (*plst).Len())

	val := TestType{1, 2}
	ptr := &TestType{3, 4}

	val.PtrTest()
	val.ValTest()

	ptr.PtrTest()
	ptr.ValTest()

	ValTest2(val)
	PtrTest2(&val)

	ValTest2(*ptr)
	PtrTest2(ptr)

}
