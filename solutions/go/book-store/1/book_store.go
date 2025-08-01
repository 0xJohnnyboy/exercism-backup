package bookstore
import "fmt"

const bookPriceInCents int = 800
const maxGroupSize int = 5
var discountsFactorsByNbOfBooks map[int]float64 = map[int]float64{
    2: 0.95,
    3: 0.90,
    4: 0.80,
    5: 0.75,
}

func Cost(books []int) int {
	counts := GetOccurences(books)
    return MinCost(counts, map[string]int{})
}

func MinCost(counts []int, memo map[string]int) int {
    key := Encode(counts)
    if val, ok := memo[key]; ok {
        return val
    }

    if AllZero(counts) {
        return 0
    }

    minTotal := int(^uint(0) >> 1)
    available := []int{}
    for i, c := range counts {
        if c > 0 {
            available = append(available, i)
        }
    }

    for size := 5; size >= 2; size-- {
        if len(available) < size {
            continue
        }

        for _, group := range Combos(available, size) {
            // decrementing counts
            for _, i := range group {
                counts[i]--
            }

            costGroup := int(float64(size * bookPriceInCents) * Discount(size))
            total := costGroup + MinCost(counts, memo)

            if total < minTotal {
                minTotal = total
            }

            // backtrack
            for _, i := range group {
                counts[i]++
            }
        }
    }

    // fallback : no groups if < 2 books
    priceWithoutGroup := 0
    for _, c := range counts {
        priceWithoutGroup += c * bookPriceInCents
    }
    if priceWithoutGroup < minTotal {
        minTotal = priceWithoutGroup
    }

    memo[key] = minTotal
    return minTotal
}

func GetOccurences(books []int) []int {
    occurences := make([]int, 5)
    for _, v := range books {
        occurences[v - 1]++
    }
    return occurences
}

func Encode(counts []int) string {
    return fmt.Sprintf("%v", counts)
}

func AllZero(counts []int) bool {
    for _, c := range counts {
        if c > 0 {
            return false
        }
    }
    return true
}

func Discount(size int) float64 {
    if f, ok := discountsFactorsByNbOfBooks[size]; ok {
        return f
    }
    return 1.0
}

func Combos(available []int, k int) [][]int {
    var result [][]int
    var backtrack func(start int, path []int)
    backtrack = func(start int, path []int) {
        if len(path) == k {
            combo := make([]int, k)
            copy(combo, path)
            result = append(result, combo)
            return
        }
        for i := start; i < len(available); i++ {
            backtrack(i+1, append(path, available[i]))
        }
    }
    backtrack(0, nil)
    
    return result
}
