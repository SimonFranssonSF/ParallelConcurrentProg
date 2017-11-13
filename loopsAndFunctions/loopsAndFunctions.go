// loopsAndFunctions uses Newton's Method to calculate the square root of a positive number.
package main

// Multiple import paths to packages fmt and math.
import (
	"errors"
	"fmt"
	"math"
)

// IN: x float64 with a square root to be calculated.
// OUT: z float64 which is an approximation of the square root of x and error.
// Will iterate the Newton method 10 times.
func Sqrt(x float64) (float64, int, error) {
	z := float64(1)
	nbrOfItr := 0
	// Returns -1 if the input is negative, 0 iterations, communicates it via an error
	if x < 0 {
		return -1, 0, errors.New("Function Sqrt only accepts numbers greater than or equal to zero.")
	}

	// Newton's Method for approximating a square root to the parameter x.
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2)-x)/(2*z)
		nbrOfItr++
	}

	// Returns the approximated square root of parameter x and no error aka nil.
	return z, nbrOfItr, nil
}

// Same as Sqrt but stops when delta < math.Abs(math.Sqrt(x) - z)
func SqrtDelta(x float64, delta float64) (float64, int, error) {
	z := float64(1)
	nbrOfItr := 0

	// Returns -1 if the input is negative, 0 iterations, communicates it via an error
	if x < 0 {
		return -1, 0, errors.New("Function Sqrt only accepts numbers greater than or equal to zero.")
	}

	// Newton's Method for approximating a square root to the parameter x.
	for delta < math.Abs((math.Sqrt(x) - z)) {
		z = z - (math.Pow(z, 2)-x)/(2*z)
		nbrOfItr++
	}

	// Returns the approximated square root of parameter x and no error aka nil.
	return z, nbrOfItr, nil
}

func main() {
	input := 2.0
	delta := 0.0

	root, nbrOfItr, err := Sqrt(input)
	fmt.Println("In:", input, "Square root:", root, "Iterations:", nbrOfItr, "Error:", err)

	root, nbrOfItr, err = SqrtDelta(input, delta)
	fmt.Println("In:", input, "Square root:", root, "Iterations:", nbrOfItr, "Error:", err)

	// Square root of input using math.Sqrt
	fmt.Println("Calculating square root using math.Sqrt:", math.Sqrt(2))

	// Testing float32, presenting square root as a float32
	f32, _, _ := Sqrt(input)
	fmt.Println("Square root as float32 instead of float64", float32(f32))
}
