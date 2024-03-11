package math

var seed int64 = 255

func Seed() int64 {
	return seed
}

func lcg() int {
	// LCG parameters (use appropriate values for your needs)
	a, c, m := int64(1664525), int64(1013904223), int64(1<<31)

	// Linear Congruential Generator formula
	seed = (a*seed + c) % m

	// Convert the result to an int
	return int(seed)
}

func RandomInRange(min, max int) int {
	// Ensure the range is valid
	if min >= max {
		panic("Invalid range")
	}

	// Use lcg() to generate a random number within the range [min, max)
	return min + lcg()%(max-min)
}
