package main

import (
	"fmt"
	"os"
	"strings"
)

type Outcome int

const (
	Loss Outcome = 0
	Draw Outcome = 3
	Win  Outcome = 6
)

type Hand int

const (
	RockHand     Hand = 1
	PaperHand    Hand = 2
	ScissorsHand Hand = 3
)

type PlayerScore int

// NOTE: Only defining this function because we have to wrap int enums with `int()` in order to add them together.
func score(hand Hand, outcome Outcome) PlayerScore {
	return PlayerScore(int(hand) + int(outcome))
}

type Hands struct {
	opponent Hand
	player   Hand
}

// Pre-computed hands so we can compute in O(1) time (only because 3x3 permutations)
var possibleHands = map[Hands]PlayerScore{
	{RockHand, RockHand}:         score(RockHand, Draw),
	{RockHand, PaperHand}:        score(PaperHand, Win),
	{RockHand, ScissorsHand}:     score(ScissorsHand, Loss),
	{PaperHand, RockHand}:        score(RockHand, Loss),
	{PaperHand, PaperHand}:       score(PaperHand, Draw),
	{PaperHand, ScissorsHand}:    score(ScissorsHand, Win),
	{ScissorsHand, RockHand}:     score(RockHand, Win),
	{ScissorsHand, PaperHand}:    score(PaperHand, Loss),
	{ScissorsHand, ScissorsHand}: score(ScissorsHand, Draw),
}

// inputsAsHands is used when parsing the input file as opponent and player hands.
var inputsAsHands = map[string]Hand{
	"A": RockHand,
	"B": PaperHand,
	"C": ScissorsHand,
	"X": RockHand,
	"Y": PaperHand,
	"Z": ScissorsHand,
}

func parseHands(line string) Hands {
	parts := strings.Split(line, " ")
	opponent := Hand(inputsAsHands[parts[0]])
	player := Hand(inputsAsHands[parts[1]])

	return Hands{opponent, player}
}

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

	lines := strings.Split(string(raw), "\n")

	fmt.Printf("Part 1 answer: %d\n", partOne(lines))
	fmt.Printf("Part 2 answer: %d\n", partTwo(lines))
}

func partOne(lines []string) PlayerScore {
	total := PlayerScore(0)
	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}

		hands := parseHands(line)
		score := possibleHands[hands]
		total += score
	}

	return total
}

type PartialRound struct {
	opponent Hand
	outcome  Outcome
}

var responses = map[PartialRound]PlayerScore{
	{RockHand, Loss}:     score(ScissorsHand, Loss),
	{RockHand, Draw}:     score(RockHand, Draw),
	{RockHand, Win}:      score(PaperHand, Win),
	{PaperHand, Loss}:    score(RockHand, Loss),
	{PaperHand, Draw}:    score(PaperHand, Draw),
	{PaperHand, Win}:     score(ScissorsHand, Win),
	{ScissorsHand, Loss}: score(PaperHand, Loss),
	{ScissorsHand, Draw}: score(ScissorsHand, Draw),
	{ScissorsHand, Win}:  score(RockHand, Win),
}

var inputAsOutcome = map[string]Outcome{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

func parsePartialRound(line string) PartialRound {
	parts := strings.Split(line, " ")
	opponent := Hand(inputsAsHands[parts[0]])
	outcome := Outcome(inputAsOutcome[parts[1]])

	return PartialRound{opponent, outcome}
}

func partTwo(lines []string) PlayerScore {
	score := PlayerScore(0)
	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}

		partialRound := parsePartialRound(line)
		roundScore := responses[PartialRound{partialRound.opponent, partialRound.outcome}]
		score += roundScore
	}

	return score
}
