package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput            = errors.New("input is empty")
	errorNotTwoOperands        = errors.New("expecting two operands, but received more or less")
	errorUnsupportedCharacters = errors.New("some characters are not supported")
	errorInvalidInput          = errors.New("input is not valid")
)

func StringSum(input string) (output string, err error) {
	arr := make([]rune, 0)
	operands := 0
	for _, v := range input {
		if !strings.Contains("0123456789+-", string(v)) {
			return "", fmt.Errorf("something went wrong: %w", errorUnsupportedCharacters)
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
		return "", fmt.Errorf("something went wrong: %w", errorInvalidInput)
	}

	second, ok := strconv.ParseInt(string(arr[len(arr)-1]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("something went wrong: %w", errorInvalidInput)
	}

	if arr[len(arr)-2] == '-' {
		second *= -1
	}

	first, ok := strconv.ParseInt(string(arr[len(arr)-3]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("something went wrong: %w", errorInvalidInput)
	}

	if len(arr) == 4 && arr[0] == '-' {
		first *= -1
	}

	return strconv.FormatInt(first+second, 10), nil
}
