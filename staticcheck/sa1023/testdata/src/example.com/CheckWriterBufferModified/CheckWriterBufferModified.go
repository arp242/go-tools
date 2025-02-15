package pkg

type T1 struct{}
type T2 struct{}
type T3 struct{}
type T4 struct{}
type T5 struct{}
type T6 struct{}
type T7 struct{}

type Bytes []byte

type AliasByte = byte
type AliasByteSlice = []byte
type AliasInt = int
type AliasError = error

func (T1) Write(b []byte) (int, error) {
	b = append(b, '\n') //@ diag(`io.Writer.Write must not modify the provided buffer`)
	_ = b
	return 0, nil
}

func (T2) Write(b []byte) (int, error) {
	b[0] = 0 //@ diag(`io.Writer.Write must not modify the provided buffer`)
	return 0, nil
}

func (T3) Write(b []byte) string {
	b[0] = 0
	return ""
}

func (T4) Write(b []byte, r byte) (int, error) {
	b[0] = r
	return 0, nil
}

func (T5) Write(b []AliasByte) (int, error) {
	b[0] = 0 //@ diag(`io.Writer.Write must not modify the provided buffer`)
	return 0, nil
}

func (T6) Write(b AliasByteSlice) (AliasInt, AliasError) {
	b[0] = 0 //@ diag(`io.Writer.Write must not modify the provided buffer`)
	return 0, nil
}

func (T7) Write(b Bytes) (int, error) {
	b[0] = 0
	return 0, nil
}
