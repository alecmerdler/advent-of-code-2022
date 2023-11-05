package main

import (
	"fmt"
	"os"
	"strings"
)

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
			if pos, ok := foundAt[item]; ok && sameCompartment(index, pos, len(line)) {
				// Flip the character from upper->lower or vice-versa so we can use the ASCII value
				var flipped string
				if lowered := strings.ToLower(string(item)); lowered != string(item) {
					fmt.Printf("flipping %s -> %s\n", string(item), lowered)
					flipped = lowered
				} else {
					fmt.Printf("flipping %s -> %s\n", string(item), strings.ToUpper(string(item)))
					flipped = strings.ToUpper(string(item))
				}

				asciiCode := int(flipped[0])
				priority := int(asciiCode) - 64
				total += priority

				// FIXME(alecmerdler): Seeing what the int value of each rune is...
				fmt.Printf("%s - %s, %d, %d\n", line, string(item), priority, total)

				break
			}

			foundAt[item] = index
		}
	}

	fmt.Println(total)
}

func sameCompartment(posA, posB, totalSize int) bool {
	return (posA < totalSize/2 && posB < totalSize/2) || (posA >= totalSize/2 && posB >= totalSize/2)
}
