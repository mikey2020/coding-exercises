package chance

import (
    "fmt"
    "math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    num := rand.Intn(20)
    fmt.Println(num)
    return num
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
