package wordy
import "fmt"
import "slices"
import "regexp"
import "strconv"
import "errors"

type Operation func(int, int) int
func Add(a int, b int) int { return a + b }
func Sub(a int, b int) int { return a - b }
func Mult(a int, b int) int { return a * b }
func Div(a int, b int) int { return a / b }

var OperandsMap = map[string]Operation{
    "plus": Add,
    "minus": Sub,
    "multiplied by": Mult,
    "divided by": Div,
}

type LoopChecker func(int, []string) error
type Checker func([]string, string) error
func CheckMissingOperand(index int, extracted []string) error {
    if index % 2 == 0 || index == len(extracted) - 1 {
        if OperandsMap[extracted[index]] != nil { return errors.New("missing operand")}
    }
    return nil
}
func CheckNumberPlacement(index int, extracted []string) error {
    if index % 2 != 0 {
        _, err := strconv.Atoi(extracted[index])
        if err == nil {
            return errors.New("Number not well placed")
        }
    } 
    return nil
}
func CheckDuplicates(index int, extracted []string) error {
    if index == 0 {
        return nil
    }
    if extracted[index-1] == extracted[index] {
        return errors.New("duplicate")
    }  
    return nil
}
func CheckMathQuestion(extracted []string, question string) error {
    if len(extracted) == 0 {
        return errors.New(fmt.Sprintf("non math question: %s", question))
    }
    return nil
}
func CheckInvalidOperation(extracted []string, question string) error {
    if len(extracted) != 1 {
        return nil
    }
	reg := regexp.MustCompile(`([a-zA-Z]+\s*\?$)`)
    ending := reg.FindString(question)
    if ending != "" {
        return errors.New("invalid operation")
    }
    return nil
}
var CheckerSlice = []Checker {
    CheckMathQuestion,
    CheckInvalidOperation,
}
var LoopCheckerSlice = []LoopChecker {
    CheckMissingOperand,
    CheckNumberPlacement,
    CheckDuplicates,
}

func Answer(question string) (int, bool) {
    reg, err := regexp.Compile(`(-?\d+|plus|minus|multiplied\sby|divided\sby)`)
    
    if err != nil {
        return 0, false
    }
   
    extracted := reg.FindAllString(question, -1)
    
    if Check(extracted, question) != nil {
        return 0, false
    }
    
    result, err := strconv.Atoi(Calculate(extracted)[0])
    
    if err != nil {
        return 0, false
    }
    
    return result, true
}

func Check(extracted []string, question string) error {
    var err error
    
    for _, c := range CheckerSlice {
        err = c(extracted, question)
        if err != nil {
            return err
        }
    }
 
    for i, _ := range extracted {
		for _, c := range LoopCheckerSlice {
            err = c(i, extracted)
            if err != nil {
                return err
            }
        }	
    }

    return nil
}

func Calculate(extracted []string) ([]string) {
    if (len(extracted) == 1) {
        return extracted
    }
    var val1, _ = strconv.Atoi(extracted[0])
    var operand = extracted[1]
    var val2, _ = strconv.Atoi(extracted[2])
    var result int
    var outExtracted []string

    result = OperandsMap[operand](val1, val2)

    outExtracted = slices.Concat(extracted[:0], []string{strconv.Itoa(result)}, extracted[3:])
    return Calculate(outExtracted)
}