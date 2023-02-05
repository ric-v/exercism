package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {

	// initial list of pickList
	pickList := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	// setting fixed length to capture output
	n := len(pickList)

	// make a slice of len n to assign animal names to random position
	var animals = make([]string, n)

	// iterate for n times
	for i := 0; i < n; i++ {

		// on each iteration get a position to place the animal name in output slice
		pos := rand.Intn(len(pickList))
		animals[i] = pickList[pos]

		// remove the animal name from initial slice of animals to avoid duplicates and errors
		pickList = append(pickList[:pos], pickList[pos+1:]...)
	}

	// OR
	// animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	// rand.Shuffle(len(animals), func(i, j int) {
	// 	animals[i], animals[j] = animals[j], animals[i]
	// })

	return animals
}
