package main

import (
	"fmt"
	"os"
	"strings"
)

// asciiBetweenCapsAndLowercase is the number of ASCII symols that separate the lowercase alphabet from the uppercase alphabet.
const asciiBetweenCapsAndLowercase = 6

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	raw := make([]byte, int(info.Size()))
	if _, err := file.Read(raw); err != nil {
		panic(err)
	}

	total := 0
	lines := strings.Split(string(raw), "\n")
	for _, line := range lines {
		foundAt := map[rune]int{}
		for index, item := range line {
			if pos, ok := foundAt[item]; ok && !sameCompartment(index, pos, len(line)) {
				foundAt[item] = index

				// Flip the character from upper->lower or vice-versa so we can use the ASCII value
				var priority int
				if lowered := strings.ToLower(string(item)); lowered != string(item) {
					flipped := lowered
					asciiCode := int(flipped[0])
					priority = int(asciiCode) - 64 - asciiBetweenCapsAndLowercase
				} else {
					flipped := strings.ToUpper(string(item))
					asciiCode := int(flipped[0])
					priority = int(asciiCode) - 64
				}

				total += priority
				break
			} else {
				foundAt[item] = index
			}
		}
	}

	fmt.Println(total)
}

func sameCompartment(posA, posB, totalSize int) bool {
	return (posA < totalSize/2 && posB < totalSize/2) || (posA >= totalSize/2 && posB >= totalSize/2)
}
