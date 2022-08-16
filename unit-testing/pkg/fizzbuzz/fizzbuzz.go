package fizzbuzz

import (
	"log"
	"strconv"
)

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
func FizzBuzz(total, fizzAt, buzzAt int64) []string {
	defer recoverSliceOutfRange()
	if total < 0 || fizzAt < 0 || buzzAt < 0 {
		log.Default().Println("input can not be less than zero")
		return nil
	}

	result := make([]string, total)

	for i := int64(1); i <= total; i++ {
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) {
			result[i-1] = strconv.FormatInt(i, 10)
			continue
		}

		if i%fizzAt == 0 {
			result[i-1] = "Fizz"
		}

		if i%buzzAt == 0 {
			result[i-1] += "Buzz"
		}
	}
	return result
}

func recoverSliceOutfRange() {
	if r := recover(); r != nil {
		log.Default().Println("total too large: could not create slice ", r)
	}
}