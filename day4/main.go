package day4

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/strangelytyped/aoc2021/utils"
)


var winningStates = []int32{
	0x0000001F,
	0x000003E0,
	0x00007C00,
	0x000F8000,
	0x01F00000,

	0x00108421,
	0x00210842,
	0x00421084,
	0x00842108,
	0x01084210,
}

type card struct {
	valueMap map[string]int32
	state int32
}

func (c *card) addNumber(num string) {
	idx, ok := c.valueMap[num]
	if ok {
		c.state |= idx
	}
}

func (c *card) isDone() bool {
	for _, state := range winningStates {
		if (c.state & state) == state {
			return true
		}
	}
	return false
}

func (c *card) sumUnset() int {
	result := 0
	for num, idx := range c.valueMap {
		if (c.state & idx) == 0 {
			result += utils.ParseIntOrPanic(num)
		}
	}
	return result
}

func makeCards(rawCards [][]string) []*card {
	var cards []*card
	for _, rawCard := range rawCards {
		valueMap := make(map[string]int32)
		for i, val := range rawCard {
			valueMap[val] = (1 << i)
		}
		cards = append(cards, &card{valueMap, 0})
	}
	return cards
}

func callNumber(cards []*card, n string) []*card {
	var winners []*card
	for _, card := range cards {
		if card.isDone() {
			continue
		}
		card.addNumber(n)
		if card.isDone() {
			winners = append(winners, card)
		}
	}
	return winners
}

func runForWinnerCount(numbers []string, rawCards [][]string, desiredWinners int) int {
	cards := makeCards(rawCards)
	var lastCard *card
	lastIdx := 0
	finishedCardCount := 0
	for idx, n := range numbers {
		finishedCards := callNumber(cards, n)
		finishedCardCount += len(finishedCards)

		if finishedCardCount >= desiredWinners {
			lastCard = finishedCards[0]
			lastIdx = idx
			break;
		}
	}

	if lastCard == nil {
		panic("No winning card found")
	}

	winningNum := utils.ParseIntOrPanic(numbers[lastIdx])
	cardSum := lastCard.sumUnset()
	return cardSum * winningNum
}

func partOne(numbers []string, rawCards [][]string) int {
	return runForWinnerCount(numbers, rawCards, 1)
}

func partTwo(numbers []string, rawCards [][]string) int {
	return runForWinnerCount(numbers, rawCards, len(rawCards))
}

func readInput(input io.Reader) ([]string, [][]string) {
	byteData, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}
	items := strings.Split(string(byteData), "\n\n")
	numberList := strings.Split(items[0], ",")

	var cards [][]string
	for _, cardStr := range items[1:] {
		cardNumbers := strings.Fields(cardStr)
		cards = append(cards, cardNumbers)
	}

	return numberList, cards
}

func Main(input io.Reader) {
	numbers, cards := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers, cards))
	fmt.Printf("Part 2: %d\n", partTwo(numbers, cards))
}
