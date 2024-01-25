package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var suits = [4]byte{67, 68, 72, 83}
var rating = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

type Card struct {
	rating int
	suit   byte
}

func getCards(c [2]string) ([2]Card, bool) {
	var cards [2]Card
	var err error
	for i := 0; i < 2; i++ {
		cards[i].suit = byte(c[i][1])
		if cards[i].rating, err = strconv.Atoi(c[i][:1]); err != nil {
			switch c[i][:1] {
			case "T":
				cards[i].rating = 10
			case "J":
				cards[i].rating = 11
			case "Q":
				cards[i].rating = 12
			case "K":
				cards[i].rating = 13
			case "A":
				cards[i].rating = 14
			}
		}
	}
	if cards[0].rating < cards[1].rating {
		cards[0], cards[1] = cards[1], cards[0]
	}
	return cards, cards[0].rating == cards[1].rating
}

func cardTostring(c Card) (str string) {
	switch c.rating {
	case 10:
		str = "T"
	case 11:
		str = "J"
	case 12:
		str = "Q"
	case 13:
		str = "K"
	case 14:
		str = "A"
	default:
		str = strconv.Itoa(c.rating)
	}
	str += string(c.suit)
	return str + "\n"
}

func cardsPrint(noCardsInDeck map[Card]bool, r []int) string {
	var str string
	var count int
	for i := range r {
		for j := range suits {
			if _, ok := noCardsInDeck[Card{r[i], suits[j]}]; !ok {
				str += cardTostring(Card{r[i], suits[j]})
				count++
			}
		}
	}
	return strconv.Itoa(count) + "\n" + str
}

func processingPair(myCards [2]Card, n int, input *bufio.Reader) string {
	var c [2]string
	set := 2
	pair := true
	exclude := make(map[Card]bool)
	exclude[myCards[0]] = true
	exclude[myCards[1]] = true
	for n > 0 {
		fmt.Fscan(input, &c[0], &c[1])
		cards, ok := getCards(c)
		exclude[cards[0]] = true
		exclude[cards[1]] = true
		if ok {
			if myCards[0].rating == cards[0].rating {
				set = 0
			}
			if myCards[0].rating < cards[0].rating {
				pair = false
			}
			if myCards[0].rating > cards[0].rating {
				for i := range suits {
					exclude[Card{cards[0].rating, suits[i]}] = true
				}
			}
		} else {
			if cards[0].rating == myCards[0].rating || cards[1].rating == myCards[0].rating {
				set--
			}
			if cards[0].rating > myCards[0].rating {
				for i := range suits {
					exclude[Card{cards[0].rating, suits[i]}] = true
				}
			}
			if cards[1].rating > myCards[0].rating {
				for i := range suits {
					exclude[Card{cards[1].rating, suits[i]}] = true
				}
			}
		}
		n--
	}
	if !pair && set == 0 {
		return "0\n"
	}
	if !pair && set != 0 {
		return cardsPrint(exclude, []int{myCards[0].rating})
	}
	return cardsPrint(exclude, rating)
}

func processingNoPair(myCards [2]Card, n int, input *bufio.Reader) string {
	var c [2]string
	var str string
	high := true
	onlyPair := false
	maxRating := 0
	exclude := make(map[Card]bool)
	exclude[myCards[0]] = true
	exclude[myCards[1]] = true
	for n > 0 {
		fmt.Fscan(input, &c[0], &c[1])
		// fmt.Println(c[0], c[1])
		cards, ok := getCards(c)
		exclude[cards[0]] = true
		exclude[cards[1]] = true
		if ok {
			if myCards[0].rating == cards[0].rating || myCards[1].rating == cards[0].rating ||
				myCards[0].rating > cards[0].rating {
				onlyPair = true
				for i := range suits {
					exclude[Card{cards[0].rating, suits[i]}] = true
				}
				if myCards[1].rating < cards[0].rating {
					for i := range suits {
						exclude[Card{myCards[1].rating, suits[i]}] = true
					}
				}
			}
			if myCards[0].rating < cards[0].rating {
				str = "0\n"
			}
		} else {
			if maxRating < cards[0].rating {
				maxRating = cards[0].rating
			}
			if cards[0].rating > myCards[0].rating {
				high = false
			}
			if !(cards[0].rating == myCards[0].rating || cards[0].rating == myCards[1].rating) {
				for i := range suits {
					exclude[Card{cards[0].rating, suits[i]}] = true
				}
			}
			if !(cards[1].rating == myCards[0].rating || cards[1].rating == myCards[1].rating) {
				for i := range suits {
					exclude[Card{cards[1].rating, suits[i]}] = true
				}
			}
		}
		n--
	}
	if str != "" {
		return str
	}
	if onlyPair {
		return cardsPrint(exclude, []int{myCards[0].rating, myCards[1].rating})
	}
	if !high {
		r := []int{myCards[0].rating, myCards[1].rating}
		for maxRating < 14 {
			maxRating++
			r = append(r, maxRating)
		}
		return cardsPrint(exclude, r)
	}
	return cardsPrint(exclude, rating)
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var t, n int
	fmt.Fscan(input, &t)
	for t > 0 {
		var c [2]string
		fmt.Fscan(input, &n)
		fmt.Fscan(input, &c[0], &c[1])
		myCards, isPair := getCards(c)
		if isPair {
			// fmt.Println(t, "Pair")
			output.WriteString(processingPair(myCards, n-1, input))
		} else {
			// fmt.Println(t, "NoPair")
			output.WriteString(processingNoPair(myCards, n-1, input))
		}
		t--
	}
}
