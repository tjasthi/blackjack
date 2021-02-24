package deck

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// deckSlice will lose elements as cards are drawn, a full deck will have elements 0-51
var deckSlice []int

const numberOfDecks = 8
const numberOfCards = 52
const minimumCapacity = numberOfDecks * numberOfCards * 3 / 5 // Reshuffles the deck at 60% capacity

/*
ShuffleDeck creates a new slice and swapping elements to simluate a shuffled deck with random numbers
*/
func ShuffleDeck() {
	deckSlice = make([]int, numberOfCards*numberOfDecks)
	for i := 0; i < len(deckSlice); i++ {
		deckSlice[i] = i
	}
	for i := 0; i < len(deckSlice); i++ {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(numberOfCards * numberOfDecks)                   // Choose a random number
		deckSlice[i], deckSlice[randIndex] = deckSlice[randIndex], deckSlice[i] // Swap i with the random number
	}
}

/*
DrawCard draws a card from the deck and the deck size is reduced by 1
*/
func DrawCard() Card {
	if len(deckSlice) < minimumCapacity {
		ShuffleDeck()
	}

	cardNumber := deckSlice[0]
	deckSlice = deckSlice[1:] // Remove the first element

	return convertNumberToCard(cardNumber)
}

/*
convertNumberToCard converts a number to a card
A card is defined by % 13 or % 4 to get the card number/suit respectively
* cardNumberVal is the number representing the card from the deck
*/
func convertNumberToCard(cardNumberVal int) Card {
	// Ignore Jokers
	// if (cardNumberVal >= 52) {
	// 	return Card{CardNumber: "Joker", CardSuit: "NA", CardValue: 0}
	// }

	cardNumberVal = cardNumberVal % 52
	cardNumber := cardNumberVal % 13
	suit := cardNumberVal % 4

	suitArray := map[int]string{0: "Diamond", 1: "Club", 2: "Heart", 3: "Spade"}
	faceCardArray := map[int]string{0: "Ace", 10: "Jack", 11: "Queen", 12: "King"}

	cardValue := 0
	cardNumberString := ""

	value := cardNumber
	if cardNumber >= 10 { // Face card
		value = 9
		cardValue = value + 1
		cardNumberString = faceCardArray[cardNumber]
	} else if cardNumber == 0 { // Ace card
		cardValue = value + 1 // Value is also 11 but this will be handled with point counting
		cardNumberString = faceCardArray[cardNumber]
	} else { // All other cards
		cardValue = value + 1
		cardNumberString = strconv.Itoa(cardNumber + 1)
	}

	cardSuit := suitArray[suit]

	return Card{CardNumber: cardNumberString, CardSuit: cardSuit, CardValue: cardValue}
}

/*
Card is a struct that contains the card number (Ace (0) - King (12)), the suit, and the blackjack value of the card
* CardNumber: is the card number which is a string that is either ace, 2-10, jack, queen, king
* CardSuit: is the card suit which is either heart, spades, diamond, club
* CardValue: is the blackjack value of the card, ace is listed as 1 but this will be treated as 1 or 11 in the point counting
*/
type Card struct {
	CardNumber string
	CardSuit   string
	CardValue  int
}

func (card Card) String() string {
	str := ""
	if card.CardNumber == "Ace" {
		str = card.CardNumber + " of " + card.CardSuit + "s with a card value: 1 or 11"
	} else {
		str = card.CardNumber + " of " + card.CardSuit + "s with a card value: " + strconv.Itoa(card.CardValue)
	}
	return str
}

/*
countTotal counts the total value given a set of cards
It returns an array of all the possible values, in the case there is an ace with a value of 1/11
It will count the static total, the total of all cards that have a singular value
It will add this static total to the dynamic totals which is the value of all the aces combined
* cards: an array of cards to be counted
*/
func countTotal(cards []Card) []int {
	staticTotal := 0
	numberOfAces := 0

	for i := 0; i < len(cards); i++ {
		if cards[i].CardNumber == "Ace" {
			numberOfAces++
		} else {
			staticTotal += cards[i].CardValue
		}
	}

	possibleValues := countDynamicTotal(staticTotal, numberOfAces, make([]int, 0))

	return possibleValues
}

/*
countDynamicTotal counts the total value given a static total and the number of aces
It is a helper to CountTotal
* staticTotal: the static total or the total value without counting any aces
* numberOfAces: the number of aces to add the total
* possibleValues: the possible totals so far
*/
func countDynamicTotal(staticTotal int, numberOfAces int, possibleValues []int) []int {
	if numberOfAces == 0 {
		return []int{staticTotal}
	}
	oneSlice := append(possibleValues, countDynamicTotal(staticTotal+1, numberOfAces-1, possibleValues)...)
	elevenSlice := append(possibleValues, countDynamicTotal(staticTotal+11, numberOfAces-1, possibleValues)...)
	return append(oneSlice, elevenSlice...)
}

/*
getHighestValue takes an array of totals and returns the highest value that is less than 21
If there is no value less than 21, it will return the lowest value
* possibleValues: the possible total values
*/
func getHighestValue(possibleTotals []int) int {
	max := 0
	minOver21 := math.MaxInt64
	for i := 0; i < len(possibleTotals); i++ {
		value := possibleTotals[i]
		if (value > max && value <= 21) {
			max = value
		} else if (value > max && value < minOver21) {
			minOver21 = value
		}
	}
	if max != 0 {
		return max
	}
	return minOver21
}

/*
GetValue takes an array of cards and returns the highest value less than 21
* cards: an array of cards to be counted
*/
func GetValue(cards []Card) int {
	valuesArray := countTotal(cards)
	return getHighestValue(valuesArray)
}
