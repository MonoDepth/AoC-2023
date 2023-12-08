package main

func selectNotEmpty(arr []string, callback func(string) string) []string {
	var result []string

	for _, item := range arr {
		itm := callback(item)
		if itm != "" {
			result = append(result, itm)
		}
	}

	return result
}

func mapTo[K any, V comparable](arr []K, callback func(item K) V) []V {
	var result []V

	for _, item := range arr {
		result = append(result, callback(item))
	}

	return result
}

func filter[K any](arr []K, callback func(item K) bool) []K {
	var result []K

	for _, item := range arr {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

func forEach[K any](arr []K, callback func(item K)) {
	for _, item := range arr {
		callback(item)
	}
}

func pow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}

var CardMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var JokerMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var HandType = map[string]int{
	"High Card":       1,
	"One Pair":        2,
	"Two Pair":        3,
	"Three of a kind": 4,
	"Full house":      5,
	"Four of a kind":  6,
	"Five of a kind":  7,
}

type Hand struct {
	Cards    []int
	OrgCards []string
	Type     int
	Bid      int
}

func convertToHand(cards []string, bid int, jokerMode bool) Hand {
	var hand Hand
	hand.Cards = mapTo(cards, func(card string) int {
		if jokerMode {
			c, _ := JokerMap[card]
			return c
		} else {
			c, _ := CardMap[card]
			return c
		}
	})

	hand.Type = getHandType(hand.Cards, jokerMode)
	hand.OrgCards = cards
	hand.Bid = bid
	return hand
}

func getHandType(cards []int, jokerMode bool) int {
	sameCardCnt := make(map[int]int)
	jokerCnt := 0
	for _, card := range cards {
		if card == 1 && jokerMode {
			jokerCnt++
		} else {
			sameCardCnt[card]++
		}
	}

	if jokerCnt > 0 {
		highestCnt := 0
		highestCard := 0
		for card, cnt := range sameCardCnt {
			if cnt > highestCnt {
				highestCnt = cnt
				highestCard = card
			}
		}

		sameCardCnt[highestCard] += jokerCnt
	}

	switch len(sameCardCnt) {
	case 1:
		return HandType["Five of a kind"]
	case 2:
		for _, cnt := range sameCardCnt {
			if cnt == 4 {
				return HandType["Four of a kind"]
			}
		}
		return HandType["Full house"]
	case 3:
		for _, cnt := range sameCardCnt {
			if cnt == 3 {
				return HandType["Three of a kind"]
			}
		}
		return HandType["Two Pair"]
	case 4:
		return HandType["One Pair"]
	default:
		return HandType["High Card"]
	}
}
