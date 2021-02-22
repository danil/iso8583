package bitmap64

// Bitmap holds bits with indexes within range from 1 to 64
type Bitmap [8]byte

func New(p []byte) Bitmap {
	var b Bitmap
	copy(b[:], p[:8])
	return b
}

// Get returns true or false for the indexes of the bits in range from 1 to 64
func (b *Bitmap) Get(i int) bool {
	if i < 1 || i > 64 {
		panic("index out of range from 1 to 64")
	}
	i -= 1
	return b[i/8]&(1<<uint(7-i%8)) != 0
}

// Set sets true value for the indexes of the bits in range from 1 to 64
func (b *Bitmap) Set(i int) {
	if i < 1 || i > 64 {
		panic("index out of range from 1 to 64")
	}
	i -= 1
	b[i/8] |= 1 << uint(7-i%8)
}
