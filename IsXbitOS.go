package main

func IsXbitSystem(bits uint) bool {
	// Performs a right shift into a 111...111 literal bits - 1 times
	// If only the value "1" remains, that means that a right shift bits times had the potential to result in exactly 0
	// If that's the case, then one must assume that the variable bits has the exact value of the number of bits in ^uint(0)
	return ^uint(0)>>(bits-1) == 1
}

func main() {
	for i := uint(0); i <= 1000; i++ {
		println("System has", i, "bits word:", IsXbitSystem(i))
	}
}
