package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func generateMD5(secretKey string, number int) [16]byte {
	input := fmt.Sprintf("%s%d", secretKey, number)
	bytes := []byte(input)
	return md5.Sum(bytes)
}

func convertMD5toHexadecimal(hash [16]byte) string {
	return hex.EncodeToString(hash[:])
}

func isHexadecimalMD5Valid(hash string, leadingZeroesNeeded int) bool {
	for i := 0; i < leadingZeroesNeeded; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

func solve(secretKey string) {
	resultPart1 := 0
	resultPart2 := 0
	for number := 1; resultPart1 == 0 || resultPart2 == 0; number++ {
		hash := generateMD5(secretKey, number)
		hexa := convertMD5toHexadecimal(hash)
		if resultPart2 == 0 && isHexadecimalMD5Valid(hexa, 6) {
			resultPart2 = number
			if resultPart1 == 0 {
				resultPart1 = number
			}
		}
		if resultPart1 == 0 && isHexadecimalMD5Valid(hexa, 5) {
			resultPart1 = number
		}
	}
	fmt.Println("Part 1 result is:", resultPart1)
	fmt.Println("Part 2 result is:", resultPart2)

}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	solve(content)
}
