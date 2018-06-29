package bitmap

// Bitmap bitmap struct
type Bitmap struct {
	data   []byte
	length uint64
}

// New create bitmap instance
func New() *Bitmap {
	return &Bitmap{}
}

// Set set the offset
func (bitmap *Bitmap) Set(offset uint64) {
	idx, pos := int(offset>>3), offset&0x07

	if len(bitmap.data) <= idx {
		size := idx - len(bitmap.data) + 1
		bitmap.data = append(bitmap.data, make([]byte, size)...)
	}

	if bitmap.Has(offset) {
		return
	}

	bitmap.length++
	bitmap.data[idx] |= (0x01 << pos)
}

// Has check the offset existence
func (bitmap *Bitmap) Has(offset uint64) bool {
	idx, pos := int(offset>>3), offset&0x07

	if len(bitmap.data) <= idx {
		return false
	}

	return (bitmap.data[idx] & (0x01 << pos)) != 0
}

// Lenght get set count
func (bitmap *Bitmap) Lenght() uint64 {
	return bitmap.length
}
