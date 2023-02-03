package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

// Error method implementaion for SillyNephewError
func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	// get the fodder amount and check for errors
	fodderAmnt, err := weightFodder.FodderAmount()

	// handle different cases
	switch {

	// check for errors but ignore ErrScaleMalfunction which is handled later, for all generic errors return 0 and the error
	case err != nil && err != ErrScaleMalfunction:
		return 0, err

	// if there are no cows, return an error for division by zero
	case cows == 0:
		return 0, errors.New("division by zero")

	// if there are negative fodder OR negative fodder and ErrScaleMalfunction return an error for negative fodder
	case fodderAmnt < 0 || (fodderAmnt < 0 && err == ErrScaleMalfunction):
		return 0, errors.New("negative fodder")

	// if there are negative cows, return an error for silly nephew
	case cows < 0:
		return 0, &SillyNephewError{cows}

	// if fodder amount is positive but with ErrScaleMalfunction, return (double the fodder amount / number of cows)
	case fodderAmnt > 0 && err == ErrScaleMalfunction:
		return fodderAmnt * 2 / float64(cows), nil
	}

	// if fodder amount is positive and no errors, return (fodder amount / number of cows)
	return fodderAmnt / float64(cows), nil
}
