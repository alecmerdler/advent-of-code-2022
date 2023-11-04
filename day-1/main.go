package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Tried to use waitgroups and a goroutine per elf block to calculate the largest block, but
// naive double for-loop seems to be the fastest.
func main() {
	defer duration(track("foo"))

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

	top := 0
	elves := strings.Split(string(raw), "\n\n")
	for _, elf := range elves {
		total := 0
		for _, item := range strings.Split(elf, "\n") {
			if item == "" {
				continue
			}

			calories, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}

			total += calories
			if total > top {
				top = total
			}
		}
	}

	fmt.Printf("Day 1 answer: %d", top)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}
