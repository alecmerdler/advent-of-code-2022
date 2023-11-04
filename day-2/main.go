package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Outcome int

const (
	Loss Outcome = 0
	Draw Outcome = 3
	Win  Outcome = 6
)

type Hand int

func (h Hand) String() string {
	switch h {
	case RockHand:
		return "Rock"
	case PaperHand:
		return "Paper"
	case ScissorsHand:
		return "Scissors"
	default:
		return "Unknown"
	}
}

const (
	RockHand     Hand = 1
	PaperHand    Hand = 2
	ScissorsHand Hand = 3
)

// Only defining this function because we have to wrap int enums with `int()` in order to add them together.
func score(hand Hand, outcome Outcome) int {
	return int(hand) + int(outcome)
}

type Hands struct {
	opponent Hand
	player   Hand
}

type Score struct {
	opponent int
	player   int
}

// Pre-computed possibleHands so we can compute in O(1) time (only because 3x3 permutations)
var possibleHands = map[Hands]Score{
	{RockHand, RockHand}:         {score(RockHand, Draw), score(RockHand, Draw)},
	{RockHand, PaperHand}:        {score(RockHand, Loss), score(PaperHand, Win)},
	{RockHand, ScissorsHand}:     {score(RockHand, Win), score(ScissorsHand, Loss)},
	{PaperHand, RockHand}:        {score(PaperHand, Win), score(RockHand, Loss)},
	{PaperHand, PaperHand}:       {score(PaperHand, Draw), score(PaperHand, Draw)},
	{PaperHand, ScissorsHand}:    {score(PaperHand, Loss), score(ScissorsHand, Win)},
	{ScissorsHand, RockHand}:     {score(ScissorsHand, Loss), score(RockHand, Win)},
	{ScissorsHand, PaperHand}:    {score(ScissorsHand, Win), score(PaperHand, Loss)},
	{ScissorsHand, ScissorsHand}: {score(ScissorsHand, Draw), score(ScissorsHand, Draw)},
}

var inputs = map[string]Hand{
	"A": RockHand,
	"B": PaperHand,
	"C": ScissorsHand,
	"X": RockHand,
	"Y": PaperHand,
	"Z": ScissorsHand,
}

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

	total := 0
	lines := strings.Split(string(raw), "\n")
	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}

		parts := strings.Split(line, " ")
		opponent := Hand(inputs[parts[0]])
		player := Hand(inputs[parts[1]])
		score := possibleHands[Hands{opponent, player}]

		total += score.player
	}

	fmt.Printf("\nscore: %d", total)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("\n%v: %v\n", msg, time.Since(start))
}
