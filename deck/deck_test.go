package deck

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShuffleDeck(t *testing.T) {
	Convey("Testing ShuffleDeck", t, func() {
		Convey("Testing the size of the deck and that the elements are unique", func() {
			ShuffleDeck()
			set := make(map[int]struct{})
			for i := 0; i < len(deckSlice); i++ {
				set[deckSlice[i]] = struct{}{}
			}
			So(numberOfCards*numberOfDecks, ShouldEqual, len(deckSlice))
			So(len(deckSlice), ShouldEqual, len(set))
		})
	})
}

func TestDrawCard(t *testing.T) {
	Convey("Testing DrawCard", t, func() {
		Convey("Testing the size of the deck reduces by one when the card is drawn", func() {
			ShuffleDeck()
			fullDeck := len(deckSlice)
			DrawCard()
			So(fullDeck-1, ShouldEqual, len(deckSlice))
		})
		Convey("Testing that the first card from the deck is returned", func() {
			ShuffleDeck()
			deckSlice[0] = 0
			card := DrawCard()
			So(card.String(), ShouldEqual, "Ace of Diamonds. Card value: 1 or 11")
		})
	})
}

func TestConvertNumberToCard(t *testing.T) {
	Convey("Testing convertNumberToCard", t, func() {
		Convey("Testing the CardNumber, CardSuit, and CardValue are converted properly", func() {
			Convey("Testing ace of diamonds", func() {
				card := convertNumberToCard(0)
				So(card.String(), ShouldEqual, "Ace of Diamonds. Card value: 1 or 11")
			})
			Convey("Testing ace of hearts", func() {
				card := convertNumberToCard(26)
				So(card.String(), ShouldEqual, "Ace of Hearts. Card value: 1 or 11")
			})
			Convey("Testing king of diamonds", func() {
				card := convertNumberToCard(12)
				So(card.String(), ShouldEqual, "King of Diamonds. Card value: 10")
			})
			Convey("Testing king of hearts", func() {
				card := convertNumberToCard(38)
				So(card.String(), ShouldEqual, "King of Hearts. Card value: 10")
			})
			Convey("Testing ace of diamonds with higher number", func() {
				card := convertNumberToCard(52)
				So(card.String(), ShouldEqual, "Ace of Diamonds. Card value: 1 or 11")
			})
		})
	})
}

func TestString(t *testing.T) {
	Convey("Testing String", t, func() {
		Convey("Testing the card values are mapped correctly", func() {
			Convey("Testing ace of diamonds", func() {
				card := Card{CardNumber: "Ace", CardSuit: "Diamond", CardValue: 1}
				So(card.String(), ShouldEqual, "Ace of Diamonds. Card value: 1 or 11")
			})
			Convey("Testing king of hears", func() {
				card := Card{CardNumber: "King", CardSuit: "Heart", CardValue: 10}
				So(card.String(), ShouldEqual, "King of Hearts. Card value: 10")
			})
		})
	})
}

func TestCountTotal(t *testing.T) {
	Convey("Testing CountTotal", t, func() {
		Convey("Testing the correct total is returned to the user", func() {
			Convey("Testing static cards", func() {
				Convey("Testing 1 card", func() {
					card1 := Card{CardNumber: "5", CardSuit: "Diamond", CardValue: 5}
					possibleTotals := countTotal([]Card{card1})
					So(possibleTotals, ShouldResemble, []int{5})
				})
				Convey("Testing multiple cards", func() {
					card1 := Card{CardNumber: "5", CardSuit: "Diamond", CardValue: 5}
					card2 := Card{CardNumber: "2", CardSuit: "Heart", CardValue: 2}
					card3 := Card{CardNumber: "King", CardSuit: "Spade", CardValue: 10}

					possibleTotals := countTotal([]Card{card1, card2, card3})
					So(possibleTotals, ShouldResemble, []int{17})
				})
			})
			Convey("Testing dynamic cards", func() {
				Convey("Testing 1 ace", func() {
					card1 := Card{CardNumber: "Ace", CardSuit: "Club", CardValue: 1}
					possibleTotals := countTotal([]Card{card1})
					sort.Ints(possibleTotals)
					So(possibleTotals, ShouldResemble, []int{1, 11})
				})
				Convey("Testing 2 aces", func() {
					card1 := Card{CardNumber: "Ace", CardSuit: "Club", CardValue: 1}
					card2 := Card{CardNumber: "Ace", CardSuit: "Heart", CardValue: 1}
					possibleTotals := countTotal([]Card{card1, card2})
					sort.Ints(possibleTotals)
					So(possibleTotals, ShouldResemble, []int{2, 12, 12, 22})
				})
				Convey("Testing 3 aces", func() {
					card1 := Card{CardNumber: "Ace", CardSuit: "Club", CardValue: 1}
					card2 := Card{CardNumber: "Ace", CardSuit: "Heart", CardValue: 1}
					card3 := Card{CardNumber: "Ace", CardSuit: "Spade", CardValue: 1}
					possibleTotals := countTotal([]Card{card1, card2, card3})
					sort.Ints(possibleTotals)
					So(possibleTotals, ShouldResemble, []int{3, 13, 13, 13, 23, 23, 23, 33})
				})
			})
			Convey("Testing dynamic cards mixed with static cards", func() {
				Convey("Testing 1 ace, 1 static card", func() {
					card1 := Card{CardNumber: "5", CardSuit: "Diamond", CardValue: 5}
					card2 := Card{CardNumber: "Ace", CardSuit: "Heart", CardValue: 1}
					possibleTotals := countTotal([]Card{card1, card2})
					sort.Ints(possibleTotals)
					So(possibleTotals, ShouldResemble, []int{6, 16})
				})
				Convey("Testing 2 aces, 1 static card", func() {
					card1 := Card{CardNumber: "5", CardSuit: "Diamond", CardValue: 5}
					card2 := Card{CardNumber: "Ace", CardSuit: "Heart", CardValue: 1}
					card3 := Card{CardNumber: "Ace", CardSuit: "Club", CardValue: 1}
					possibleTotals := countTotal([]Card{card1, card2, card3})
					sort.Ints(possibleTotals)
					So(possibleTotals, ShouldResemble, []int{7, 17, 17, 27})
				})
				Convey("Testing 2 aces, 2 static cards", func() {
					card1 := Card{CardNumber: "5", CardSuit: "Diamond", CardValue: 5}
					card2 := Card{CardNumber: "King", CardSuit: "Spade", CardValue: 10}
					card3 := Card{CardNumber: "Ace", CardSuit: "Heart", CardValue: 1}
					card4 := Card{CardNumber: "Ace", CardSuit: "Club", CardValue: 1}
					possibleTotals := countTotal([]Card{card1, card2, card3, card4})
					sort.Ints(possibleTotals)
					So(possibleTotals, ShouldResemble, []int{17, 27, 27, 37})
				})
			})
		})
	})
}

func TestCountDynamicTotal(t *testing.T) {
	Convey("Testing dynamic total", t, func() {
		Convey("Empty Test", func() {
			possibleTotals := countDynamicTotal(0, 0, make([]int, 0))
			sort.Ints(possibleTotals)
			So(possibleTotals, ShouldResemble, []int{0})
		})
		Convey("Single ace", func() {
			possibleTotals := countDynamicTotal(0, 1, make([]int, 0))
			sort.Ints(possibleTotals)
			So(possibleTotals, ShouldResemble, []int{1, 11})
		})
		Convey("Two aces", func() {
			possibleTotals := countDynamicTotal(0, 2, make([]int, 0))
			sort.Ints(possibleTotals)
			So(possibleTotals, ShouldResemble, []int{2, 12, 12, 22})
		})
		Convey("Static, 1 ace", func() {
			possibleTotals := countDynamicTotal(13, 1, make([]int, 0))
			sort.Ints(possibleTotals)
			So(possibleTotals, ShouldResemble, []int{14, 24})
		})
	})
}

func TestGetHighestValue(t *testing.T) {
	Convey("Testing getHighestValue", t, func() {
		Convey("1 value", func() {
			possibleTotals := []int{1}
			So(getHighestValue(possibleTotals), ShouldEqual, 1)
		})
		Convey("2 values", func() {
			possibleTotals := []int{1, 15}
			So(getHighestValue(possibleTotals), ShouldEqual, 15)
		})	
		Convey("2 values, 1 over 21", func() {
			possibleTotals := []int{1, 25}
			So(getHighestValue(possibleTotals), ShouldEqual, 1)
		})	
		Convey("2 values, 2 over 21", func() {
			possibleTotals := []int{29, 25}
			So(getHighestValue(possibleTotals), ShouldEqual, 25)
		})	
	})
}
