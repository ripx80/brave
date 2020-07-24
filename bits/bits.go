package bits

import "fmt"

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	return n &^ (1 << pos)
}

func hasBit(n int, pos uint) bool {
	return (n&(1<<pos) > 0)
}

func pb(code uint64, len int) {
	fmt.Printf(fmt.Sprintf("sending {%%d}: %%0%db \n", len),
		len,
		code,
	)
}
