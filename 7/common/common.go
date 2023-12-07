package common

import (
	"fmt"
	"strconv"
)

const (
	HandFiveOfAKind  = 6 - iota
	HandFourOfAKind  = 6 - iota
	HandFullHouse    = 6 - iota
	HandThreeOfAKind = 6 - iota
	HandTwoPair      = 6 - iota
	HandOnePair      = 6 - iota
	HandHighCard     = 6 - iota
)

const (
	CardJocker = iota
	CardTwo
	CardThree
	CardFour
	CardFive
	CardSix
	CardSeven
	CardEight
	CardNine
	CardTen
	CardJockey
	CardQueen
	CardKing
	CardAce
)

type CamelCardHand struct {
	HandType int
	Cards    [5]int
	Bid      int
}

func ParseInput(lines []string, v1 bool) []CamelCardHand {
	hands := make([]CamelCardHand, len(lines))
	for idx, line := range lines {
		hand := CamelCardHand{}

		cardsSeen := make(map[int]int)
		for i, r := range line[:5] {
			cardValue := RuneToValue(r, v1)
			hand.Cards[i] = cardValue
			cardsSeen[cardValue] += 1
		}

		hand.Bid = MustParseToInt(line[6:])
		hand.HandType = DetermineHandTypeV2(cardsSeen)
		//if v1 {
		//	hand.HandType = DetermineHandTypeV1(cardsSeen)
		//} else {
		//	hand.HandType = DetermineHandTypeV2(cardsSeen)
		//}
		hands[idx] = hand
	}

	return hands
}

func RuneToValue(card rune, v1 bool) int {
	switch card {
	case 'A':
		return CardAce
	case 'K':
		return CardKing
	case 'Q':
		return CardQueen
	case 'J':
		if v1 {
			return CardJockey
		}
		return CardJocker
	case 'T':
		return CardTen
	case '9':
		return CardNine
	case '8':
		return CardEight
	case '7':
		return CardSeven
	case '6':
		return CardSix
	case '5':
		return CardFive
	case '4':
		return CardFour
	case '3':
		return CardThree
	case '2':
		return CardTwo
	}
	panic("are we playing with pokemon cards? > " + string(card) + " <")
}

func MustParseToInt(numberString string) int {
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic("this is not a number: " + numberString)
	}
	return number
}

func DetermineHandTypeV2(cardsSeen map[int]int) int {
	keyCount := len(cardsSeen)
	if keyCount == 1 {
		return HandFiveOfAKind
	}

	offset := cardsSeen[CardJocker]

	if keyCount == 4 {
		// jockeys are not the pair => increase to three of a kind
		if offset != 0 {
			return HandThreeOfAKind
		}
		return HandOnePair
	}

	if keyCount == 5 {
		// either still highcard or raised to one pair, if a jockey is found
		return HandHighCard + offset
	}

	// full house, 4 of a kind
	if keyCount == 2 {
		// if one of the two card types is a jockey, it's a five of a kind
		if offset > 0 {
			return HandFiveOfAKind
		}

		for _, count := range cardsSeen {
			if count == 4 || count == 1 {
				return HandFourOfAKind
			}
			return HandFullHouse
		}
	}

	//2 pair, 3 of a kind
	for _, count := range cardsSeen {
		if count == 1 {
			continue
		}

		// at least one pair => 2 pair
		if count == 2 {
			if offset == 1 {
				// the last single is a jockey
				return HandFullHouse
			} else if offset == 2 {
				// one of the pairs are jockeys
				return HandFourOfAKind
			}
			return HandTwoPair
		}

		// no pair => 3 of a kind
		if offset > 0 {
			return HandFourOfAKind
		}
		return HandThreeOfAKind
	}

	panic(fmt.Sprintf("This hand matches no rule: %v", cardsSeen))
}
