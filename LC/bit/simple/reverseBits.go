package simple

// reverseBits 190. 颠倒二进制位
func reverseBits(num uint32) (rev uint32) {
	// 逐位颠倒
	for i := 0; i < 32; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}
	return

}
