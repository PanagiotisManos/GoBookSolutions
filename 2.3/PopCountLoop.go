package popcount

var pcLoop [256]byte

func init() {
	for i := range pc {
		pcLoop[i] = pcLoop[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pcLoop[byte(x>>(i*8))])
	}
	return count
}
