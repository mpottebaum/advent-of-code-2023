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
	"J": 11,
	"T": 10,
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

		// identify hand type
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
		countsSlice := []int{}
		for _, count := range cardCountsMap {
			countsSlice = append(countsSlice, count)
		}

		// create a map k: hand type v: [][hand, bid] slice of hand tuples
		handData := [2]string{
			hand,
			bid,
		}
		if len(countsSlice) == 1 {
			// five of a kind
			handTypes[0] = append(handTypes[0], handData)
		} else if len(countsSlice) == 2 {
			isFourOfAKind := false
			for iCount := 0; iCount < 2; iCount++ {
				count := countsSlice[iCount]
				if count == 4 {
					isFourOfAKind = true
					break
				}
			}
			if isFourOfAKind {
				// four of a kind
				handTypes[1] = append(handTypes[1], handData)
			} else {
				// full house
				handTypes[2] = append(handTypes[2], handData)
			}
		} else if len(countsSlice) == 3 {
			isThreeOfAKind := false
			for iCount := 0; iCount < 3; iCount++ {
				count := countsSlice[iCount]
				if count == 3 {
					isThreeOfAKind = true
					break
				}
			}
			if isThreeOfAKind {
				// three of a kind
				handTypes[3] = append(handTypes[3], handData)

			} else {
				// two pair
				handTypes[4] = append(handTypes[4], handData)

			}
		} else if len(countsSlice) == 4 {
			// one pair
			handTypes[5] = append(handTypes[5], handData)

		} else {
			// high card
			handTypes[6] = append(handTypes[6], handData)
		}
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
				var valA, valB int
				cardA := splitA[iCard]
				cardB := splitB[iCard]
				if val, ok := faceCardMap[cardA]; ok {
					valA = val
				}
				if val, err := utils.ParseInt(cardA); err == nil {
					valA = val
				}
				if val, ok := faceCardMap[cardB]; ok {
					valB = val
				}
				if val, err := utils.ParseInt(cardB); err == nil {
					valB = val
				}
				if valA > valB {
					return true
				}
				if valA < valB {
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
