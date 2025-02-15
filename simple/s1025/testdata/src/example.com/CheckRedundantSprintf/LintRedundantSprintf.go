package pkg

import "fmt"

type T1 string
type T2 T1
type T3 int
type T4 int
type T5 int
type T6 string
type T7 []byte

type T9 string
type T10 []byte
type T11 int

type MyByte byte

type Alias = byte
type T13 = []byte
type T14 = []Alias

type T15 = string
type T16 = T9
type T17 = T6

func (T3) String() string        { return "" }
func (T6) String() string        { return "" }
func (T4) String(arg int) string { return "" }
func (T5) String()               {}

func (T9) Format(f fmt.State, c rune)  {}
func (T10) Format(f fmt.State, c rune) {}
func (T11) Format(f fmt.State, c rune) {}
func (T11) String() string             { return "" }

func fn() {
	var t1 T1
	var t2 T2
	var t3 T3
	var t4 T4
	var t5 T5
	var t6 T6
	var t7 T7
	var t9 T9
	var t10 T10
	var t11 T11
	var t12 []Alias
	var t13 T13
	var t14 T14
	var t15 T15
	var t16 T16
	var t17 T17

	_ = fmt.Sprintf("%s", "test")      //@ diag(`is already a string`)
	_ = fmt.Sprintf("%s", t1)          //@ diag(`is a string`)
	_ = fmt.Sprintf("%s", t2)          //@ diag(`is a string`)
	_ = fmt.Sprintf("%s", t3)          //@ diag(`should use String() instead of fmt.Sprintf`)
	_ = fmt.Sprintf("%s", t3.String()) //@ diag(`is already a string`)
	_ = fmt.Sprintf("%s", t4)
	_ = fmt.Sprintf("%s", t5)
	_ = fmt.Sprintf("%s %s", t1, t2)
	_ = fmt.Sprintf("%v", t1)
	_ = fmt.Sprintf("%s", t6)  //@ diag(`should use String() instead of fmt.Sprintf`)
	_ = fmt.Sprintf("%s", t7)  //@ diag(`underlying type is a slice of bytes`)
	_ = fmt.Sprintf("%s", t12) //@ diag(`underlying type is a slice of bytes`)
	_ = fmt.Sprintf("%s", t13) //@ diag(`underlying type is a slice of bytes`)
	_ = fmt.Sprintf("%s", t14) //@ diag(`underlying type is a slice of bytes`)
	_ = fmt.Sprintf("%s", t15) //@ diag(`is already a string`)
	_ = fmt.Sprintf("%s", t17) //@ diag(`should use String() instead of fmt.Sprintf`)

	// don't simplify types that implement fmt.Formatter
	_ = fmt.Sprintf("%s", t9)
	_ = fmt.Sprintf("%s", t10)
	_ = fmt.Sprintf("%s", t11)
	_ = fmt.Sprintf("%s", t16)
}
