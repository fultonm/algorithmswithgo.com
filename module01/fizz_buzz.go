package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	i := 1
	s := ""
	for i <= n {
		if i == 1 {
			s += "1"
		} else if i%15 == 0 {
			s += ", Fizz Buzz"
		} else if i%3 == 0 {
			s += ", Fizz"
		} else if i%5 == 0 {
			s += ", Buzz"
		} else {
			s += ", " + fmt.Sprint(i)
		}
		i++
	}
	fmt.Println(s)
}
