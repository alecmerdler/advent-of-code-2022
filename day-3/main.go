package main

import (
	"fmt"
	"os"
	"time"
)

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
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}
