package main

func reverseBits(num uint32) uint32 {
	var bits [32]uint32
	for i := 0; i < 32; i++ {
		bits[i] = (num >> i) & 1
	}
	var result uint32
	for i := 0; i < 32; i++ {
		result |= bits[i] << (31 - i)
	}
	return result
}
