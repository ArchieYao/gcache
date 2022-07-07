package gcache

type ByteView struct {
	B []byte
}

func (bv ByteView) Len() int {
	return len(bv.B)
}

func (bv ByteView) ByteSlice() []byte {
	return CloneBytes(bv.B)
}

func (bv ByteView) String() string {
	return string(bv.B)
}

func CloneBytes(b []byte) []byte {
	b1 := make([]byte, len(b))
	copy(b1, b)
	return b1
}
