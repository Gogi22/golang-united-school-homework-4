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
	generalErrorText           = "something went wrong"
)

func StringSum(input string) (output string, err error) {
	arr := make([]rune, 0)
	operands := 0
	for _, v := range input {
		if !strings.Contains("0123456789+- ", string(v)) {
			return "", fmt.Errorf("%s: %w", generalErrorText, errorUnsupportedCharacters)
		}
		if v != ' ' {
			arr = append(arr, v)
		}
		if v == '+' || v == '-' {
			operands++
		}
	}

	if arr == nil || len(arr) == 0 {
		return "", fmt.Errorf("%s: %w", generalErrorText, errorEmptyInput)
	}

	i := 0
	if arr[0] == '+' || arr[0] == '-' {
		i = 1
	}

	for i < len(arr) && !(arr[i] == '+' || arr[i] == '-') {
		i++
	}

	first, ok := strconv.ParseInt(string(arr[:i]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("%s: %w", generalErrorText, errorInvalidInput)
	}

	second, ok := strconv.ParseInt(string(arr[i:]), 10, 0)

	if ok != nil {
		return "", fmt.Errorf("%s: %w", generalErrorText, errorInvalidInput)
	}

	return strconv.FormatInt(first+second, 10), nil
}
