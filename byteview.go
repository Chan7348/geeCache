package geecache

// 缓存值的抽象和存放

type ByteView struct {
	bytes []byte
}

func (view ByteView) Len() int {
	return len(view.bytes)
}

// 获取一个拷贝
func (view ByteView) BytesSlice() []byte {
	return cloneBytes(view.bytes)
}

// 获取拷贝并转换成string
func (view ByteView) String() string {
	return string(view.bytes)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
