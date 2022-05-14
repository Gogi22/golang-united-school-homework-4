package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	arr := make([]rune, 0)
	operands := 0
	for _, v := range input {
		if !strings.Contains("0123456789+-", string(v)) {
			return "", fmt.Errorf("use of unsupported character")
		}
		if v != ' ' {
			arr = append(arr, v)
		}
		if v == '+' || v == '-' {
			operands++
		}
	}

	if arr == nil || len(arr) == 0 {
		return "", fmt.Errorf("something went wrong: %w", errorEmptyInput)
	}

	if !(len(arr) == 4 && operands == 2 || len(arr) == 3 && operands == 1) {
		return "", fmt.Errorf("something went wrong: %w", errorNotTwoOperands)
	}

	second, ok := strconv.ParseInt(string(arr[len(arr)-1]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("invalid input")
	}

	if arr[len(arr)-2] == '-' {
		second *= -1
	}

	first, ok := strconv.ParseInt(string(arr[len(arr)-3]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("invalid input")
	}

	if len(arr) == 4 && arr[0] == '-' {
		first *= -1
	}

	return strconv.FormatInt(first+second, 10), nil
}
