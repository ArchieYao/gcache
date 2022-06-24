package gcache

type ByteView struct {
	b []byte
}

func (bv ByteView) Len() int {
	return len(bv.b)
}

func (bv ByteView) ByteSlice() []byte {
	return cloneBytes(bv.b)
}

func (bv ByteView) String() string {
	return string(bv.b)
}

func cloneBytes(b []byte) []byte {
	b1 := make([]byte, len(b))
	copy(b1, b)
	return b1
}
