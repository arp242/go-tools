package pkg

import (
	"bytes"
)

type AliasByte = byte
type AliasSlice = []byte
type AliasSlice2 = []AliasByte

func fn() {
	buf := bytes.NewBufferString("str")
	_ = string(buf.Bytes())  //@ diag(`should use buf.String() instead of string(buf.Bytes())`)
	_ = []byte(buf.String()) //@ diag(`should use buf.Bytes() instead of []byte(buf.String())`)

	m := map[string]*bytes.Buffer{"key": buf}
	_ = string(m["key"].Bytes())  //@ diag(`should use m["key"].String() instead of string(m["key"].Bytes())`)
	_ = []byte(m["key"].String()) //@ diag(`should use m["key"].Bytes() instead of []byte(m["key"].String())`)

	_ = []AliasByte(m["key"].String()) //@ diag(`should use m["key"].Bytes() instead of []AliasByte(m["key"].String())`)
	_ = AliasSlice(m["key"].String())  //@ diag(`should use m["key"].Bytes() instead of AliasSlice(m["key"].String())`)
	_ = AliasSlice2(m["key"].String()) //@ diag(`should use m["key"].Bytes() instead of AliasSlice2(m["key"].String())`)

	var m2 map[string]int
	_ = m2[string(buf.Bytes())] // no warning, this is more efficient than buf.String()

	string := func(_ interface{}) interface{} {
		return nil
	}
	_ = string(m["key"].Bytes())
}
