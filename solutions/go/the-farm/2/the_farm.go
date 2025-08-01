package thefarm
import "fmt"
import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowsQty int) (float64, error) {
    fa, err := fc.FodderAmount(cowsQty)
    if err != nil {
        return 0, err
    }

    ff, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return fa * ff / float64(cowsQty), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowsQty int) (float64, error) {
    if cowsQty <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    
    return DivideFood(fc, cowsQty)
}

type InvalidCowsError struct {
    cowsQty int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cowsQty, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowsQty int) error {
    switch {
        case cowsQty < 0: return &InvalidCowsError{
            cowsQty: cowsQty,
            message: "there are no negative cows",
        }
        case cowsQty == 0: return &InvalidCowsError{
            cowsQty: cowsQty,
            message: "no cows don't need food",
        }
        default: return nil
    } 
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.