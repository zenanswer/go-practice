package methodsets

import (
	"fmt"
	"reflect"
)

func init() {

}

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func IntfTest() {
	// A bare value
	var lst List
	// CountInto(lst, 1, 10) // INVALID: Append has a pointer receiver
	CountInto(&lst, 1, 10)
	if LongEnough(lst) { // VALID: Identical receiver type
		fmt.Printf(" - lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {  // VALID: a *List can be dereferenced for the receiver
		fmt.Printf(" - plst is long enough\n")
	}

	var n1 Lener
	n1 = lst
	fmt.Printf("Type of Appender is %s", reflect.TypeOf(n1))
	var n2 Appender
	n2 = plst
	fmt.Printf("Type of Appender is %s", reflect.TypeOf(n2))

}
