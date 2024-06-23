package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	fmt.Printf("\nstring to parse:\n\n%v\n\n", s)

	var builder strings.Builder

	var digAfter, digBefore int
	var err1, err2 error
	var news string

	builder.Grow(len(s))

	for i := range len(s) {

		d, _ := strconv.Atoi(string(s[i]))
		// if err != nil {
		// 	continue
		// }

		if i < len(s)-1 {
			digAfter, err1 = strconv.Atoi(string(s[i+1]))
		} else {
			digAfter, err1 = strconv.Atoi(string(s[i]))
		}

		if i > 0 {
			digBefore, err2 = strconv.Atoi(string(s[i-1]))
		}

		if IsDigitAfterBefore(err1, d, digAfter) || IsDigitAfterBefore(err2, d, digBefore) {
			continue
		}

		news += strings.Repeat(string(s[i-1]), d)
		// if !(err1 != nil) {
		// 	builder.Write([]byte(strings.Repeat(string(s[i-1]), d)))
		// }
		builder.Write([]byte(strings.Repeat(string(s[i-1]), d)))

		if err1 != nil && err2 != nil && i < len(s)-1 {
			news += string(s[i+1])
			builder.Write([]byte(string(s[i+1])))
		}

		fmt.Println(i, "\t", string(s[i]), d, "\tnew str:", news, "\tdigAfter:", digAfter, "digBefore:", digBefore)
	}

	// fmt.Printf("final result: %v", builder.String())
	fmt.Printf("final result: %v", news)

	return builder.String(), nil // news, nil
}

func IsDigitAfterBefore(e error, d1, d2 int) bool {
	return (e == nil && d1 >= 0 && d2 >= 0)
}
