package main

import (
	"fmt"
	"os"
)

func test(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case map[string]any:
		fmt.Println("Object")
	case []any:
		fmt.Println("Array")
	default:
		fmt.Println("Unknown")
	}
}

func create() map[string]any {
	obj := make(map[string]any)
	obj["string"] = "hello"
	return obj
}

func solve(content string) {
	obj := make(map[string]any)
	obj["int"] = 32
	array := make([]any, 0)
	array = append(array, 3)
	obj["array"] = array
	obj["string"] = "test"
	obj["obj"] = create()

	for _, value := range obj {
		test(value)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	solve(content)
}
