package hamming

import "errors"

// Distance calculates the hamming distance between 2 strands of DNA
//	input : a (DNA strand 1), b (DNA strand 2)
//	output: hammingDistance, error
func Distance(a, b string) (int, error) {

	// if the strands are of unequal lengths return error
	if len(a) != len(b) {
		return 0, errors.New("Mismatch")
	}

	var hammingDistance int
	// iterate on each DNA from strand 1
	for index, dna := range a {

		// check the DNA strand value against the strand value corresponding to the same
		// position in strand 2, if they donot match, update hamming distance counter
		if byte(dna) != b[index] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
