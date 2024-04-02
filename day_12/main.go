package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// receives a input string starting with "
// terminating with "
// parse the string and return it
// returns the remainder of the string
func parseString(input string) (string, string) {
	result, follow, _ := strings.Cut(input[1:], "\"")
	return result, follow
}

// receives a string with a continuation of number
// followed by the rest of the thing to parse
// parse the string, convert it to int and return it
// along with the remainder of the string
func parseInt(input string) (int, string) {
	var i int
	for i = 0; input[i] == '-' || input[i] >= '0' && input[i] <= '9'; i++ {
	}
	number, _ := strconv.ParseInt(input[:i], 10, 32)
	return int(number), input[i:]
}

func parseArray(input string) ([]any, string) {
	next := input
	array := make([]any, 0)
	for next[0] != ']' {
		value, continuation := parseUnknown(next[1:])
		array = append(array, value)
		next = continuation
	}
	return array, next[1:]
}

// This function receive an input string that starts with {
// and which is a correct object until the terminating }
// This function will parse the entire object, and return it
// With the remainder of the string
func parseObject(input string) (map[string]any, string) {
	next := input
	object := make(map[string]any)
	for next[0] != '}' {
		key, continuation := parseString(next[1:])
		value, continuation := parseUnknown(continuation[1:])
		object[key] = value
		next = continuation
	}
	return object, next[1:]
}

func parseUnknown(input string) (any, string) {
	switch input[0] {
	case '"':
		return parseString(input)
	case '[':
		return parseArray(input)
	case '{':
		return parseObject(input)
	default:
		return parseInt(input)
	}
}

func parseContent(content string) map[string]any {
	object, _ := parseObject(content)
	return object
}

func sumIntInArray(array []any, redMatters bool) int {
	result := 0
	for _, value := range array {
		number, _ := getSumIntOfAnyAndIfRed(value, redMatters)
		result += number
	}
	return result
}

func sumIntInObject(object map[string]any, redMatters bool) int {
	result := 0
	for _, value := range object {
		number, isRedString := getSumIntOfAnyAndIfRed(value, redMatters)
		if redMatters && isRedString {
			return 0
		}
		result += number
	}
	return result
}

func getSumIntOfAnyAndIfRed(value any, redMatters bool) (int, bool) {
	switch value.(type) {
	case int:
		return value.(int), false
	case map[string]any:
		return sumIntInObject(value.(map[string]any), redMatters), false
	case []any:
		return sumIntInArray(value.([]any), redMatters), false
	default:
		return 0, value.(string) == "red"
	}
}

func solve(object map[string]any) {
	fmt.Println("Result of part 1:", sumIntInObject(object, false))
	fmt.Println("Result of part 2:", sumIntInObject(object, true))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	object := parseContent(content)
	solve(object)
}
