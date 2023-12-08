package day07

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

var faceCardMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
}

func GetHandType(hand string) int {
	cards := strings.Split(hand, "")
	cardCountsMap := map[string]int{}
	for iCard := 0; iCard < len(cards); iCard++ {
		card := cards[iCard]
		if count, exists := cardCountsMap[card]; exists {
			cardCountsMap[card] = count + 1
		} else {
			cardCountsMap[card] = 1
		}
	}
	numJokers := 0
	counts := []int{}
	for card, count := range cardCountsMap {
		counts = append(counts, count)
		if card == "J" {
			numJokers = count
		}
	}

	sort.Slice(counts, func(iA, iB int) bool {
		countA := counts[iA]
		countB := counts[iB]
		return countA > countB
	})

	countsLength := len(counts)

	countsLenNoJoke := countsLength - 1
	switch numJokers {
	case 4:
		// always 5kind
		return 0
	case 3:
		// 2 are eq => 5kind
		if countsLenNoJoke == 1 {
			return 0
		}
		// 2 are diff => 4kind
		return 1
	case 2:
		// 3 are eq => 5kind
		if countsLenNoJoke == 1 {
			return 0
		}
		// 2 are eq, 1 diff => 4kind
		if countsLenNoJoke == 2 {
			return 1
		}
		// 3 diff => 3kind
		if countsLenNoJoke == 3 {
			return 3
		}
	case 1:
		// 4 are eq => 5kind
		if countsLenNoJoke == 1 {
			return 0
		}
		if countsLenNoJoke == 2 {
			// 3 are eq, 1 diff => 4kind
			if counts[0] == 3 {
				return 1
			}
			// 2 eq, 2eq  => FH
			return 2
		}
		if countsLenNoJoke == 3 {
			// 2 eq, 2 diff => 3kind
			return 3
		}
		if countsLenNoJoke == 4 {
			// 4 diff => 1pair
			return 5
		}
	default:
		// do nothing
	}

	switch countsLength {
	case 1:
		// five of a kind
		return 0
	case 2:
		if counts[0] == 4 {
			// four of a kind
			return 1
		}
		// full house
		return 2
	case 3:
		if counts[0] == 3 {
			// three of a kind
			return 3
		}
		// two pair
		return 4
	case 4:
		// one pair
		return 5
	default:
		// high card
		return 6
	}
}

func GetCardStrength(card string) int {
	if val, ok := faceCardMap[card]; ok {
		return val
	}
	if val, err := utils.ParseInt(card); err == nil {
		return val
	}
	return 0
}

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day07/" + fileName + ".txt")
	rows := strings.Split(input, "\n")
	totalNumHands := len(rows)

	handTypes := [7][][2]string{}
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		// create tuples of each [hand, bid]
		handAndBidStrs := strings.Split(row, " ")
		hand := handAndBidStrs[0]
		bid := handAndBidStrs[1]

		handData := [2]string{
			hand,
			bid,
		}
		// identify hand type
		handTypeIndex := GetHandType(hand)
		handTypes[handTypeIndex] = append(handTypes[handTypeIndex], handData)
	}

	totalWinnings := 0
	// for each hand type (key of map)
	rank := totalNumHands
	for iHandType := 0; iHandType < len(handTypes); iHandType++ {
		hands := handTypes[iHandType]
		// sort by card strength
		sort.Slice(hands, func(iA, iB int) bool {
			// find which is greater
			handA := hands[iA][0]
			handB := hands[iB][0]
			splitA := strings.Split(handA, "")
			splitB := strings.Split(handB, "")
			for iCard := 0; iCard < len(splitA); iCard++ {
				cardA := splitA[iCard]
				cardB := splitB[iCard]
				strengthA := GetCardStrength(cardA)
				strengthB := GetCardStrength(cardB)
				if strengthA > strengthB {
					return true
				}
				if strengthA < strengthB {
					return false
				}
			}
			return false
		})
		// multiply hand bid by rank
		for iHand := 0; iHand < len(hands); iHand++ {
			handData := hands[iHand]
			bidStr := handData[1]
			if bid, err := utils.ParseInt(bidStr); err == nil {
				totalWinnings += rank * bid
			}
			rank--
		}
	}

	fmt.Println("Total winnings: ", totalWinnings)
}
