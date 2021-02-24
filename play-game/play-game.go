package playgame

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blackjack/deck"
)

const minBet = 10
const returnRatio = 0.5
const pause = 1
const pauseLong = 3

var startingBalance float64
var balance float64

func StartGame() {
	deck.ShuffleDeck()
	getUsername()
}

func getUsername() {
	balance = 0
	startingBalance = 0

	fmt.Println("Enter your username: ")
	var username string
	fmt.Scanln(&username)

	for startingBalance < minBet {
		fmt.Println("Enter your buy in: ")
		fmt.Scanln(&startingBalance)

		if startingBalance < minBet {
			fmt.Printf("You entered %.2f but minimum bet is %d. Please enter a higher amount\n", startingBalance, minBet)
		}
	}
	balance = startingBalance

	playGame(username)
}

func playGame(username string) {
	fmt.Printf("Username: "+username+". Current balance: %.2f\n", balance)
	fmt.Println("Enter a bet or type q or quit to leave")
	var command string
	fmt.Scanln(&command)
	strings.ToLower(command)

	if command == "q" || command == "quit" {
		fmt.Printf("Thank you for playing. You started wth %.2f and are taking %.2f home\n", startingBalance, balance)
		StartGame()
	}

	betAsInt, err := strconv.Atoi(command)
	bet := float64(betAsInt)
	if err != nil {
		fmt.Println("There was a problem understanding your input. Please enter q or quit or a number to place a bet.")
		playGame(username)
	}
	if bet > balance {
		fmt.Printf("You only have %.2f in your balance but you placed a bet of %.2f. Please try again.\n", balance, bet)
		playGame(username)
	}
	if bet < minBet {
		fmt.Printf("The minimum bet is %d. Please try again.\n", minBet)
		playGame(username)
	}

	dealCards(username, bet)
}

func dealCards(username string, bet float64) {
	fmt.Printf("%.2f bet has been placed, dealing cards now\n", bet)

	cardUser1 := deck.DrawCard()
	cardUser2 := deck.DrawCard()

	cardDealer := deck.DrawCard()

	fmt.Println("Your first card is " + cardUser1.String())
	fmt.Println("Your second card is " + cardUser2.String())
	fmt.Println("The dealer's card is " + cardDealer.String())

	cardsDrawn := []deck.Card{cardUser1, cardUser2}
	total := deck.GetValue(cardsDrawn)
	for total <= 21 {
		time.Sleep(time.Second * pause)
		if total == 21 {
			fmt.Printf("Blackjack! You win\n")
			time.Sleep(time.Second * pauseLong)
			balance += bet * returnRatio
			playGame(username)
		}

		fmt.Printf("Would you like to hit or stay? Your score is %d\n", total) // TODO Offer to split or insurance or double down
		var action string
		fmt.Scanln(&action)
		strings.ToLower(action)

		if action == "hit" || action == "h" {
			newCard := deck.DrawCard()
			fmt.Println("Your next card is " + newCard.String())

			cardsDrawn = append(cardsDrawn, newCard)
			total = deck.GetValue(cardsDrawn)
		} else if action == "stay" || action == "s" {
			fmt.Printf("You've chosen to stay, your point value is %d.\nThe dealer has a "+cardDealer.String()+" and will now draw.\n", total)
			time.Sleep(time.Second * pause)
			dealerPlay(username, bet, cardDealer, total)
		} else {
			fmt.Println("You entered an action that doesn't make sense. Enter \"hit\" or \"h\" to draw a card or enter \"stay\" or \"s\" to continue.")
		}
	}
	fmt.Printf("Bust! Your score is %d and you lose your bet\n", total)
	time.Sleep(time.Second * pauseLong)
	balance -= bet
	playGame(username)
}

func dealerPlay(username string, bet float64, cardDrawn deck.Card, pointTotal int) {
	cardsDrawn := []deck.Card{cardDrawn}
	total := deck.GetValue(cardsDrawn)
	
	for total < 16 {
		time.Sleep(time.Second * pause)
		newCard := deck.DrawCard()

		cardsDrawn = append(cardsDrawn, newCard)
		total = deck.GetValue(cardsDrawn)

		if total <= 21 && total > 15 {
			// if total == 21 && len(cardsDrawn) == 2 { // Unsure if dealer auto wins with blackjack
			// 	fmt.Printf("Dealer drew a "+newCard.String()+".\nDealer got BlackJack! You lose\n")
			// 	balance -= bet
			// }
			if total < pointTotal {
				fmt.Printf("Dealer drew a "+newCard.String()+".\nYou win your bet! Dealer got a score of %d and you beat the dealer by %d\n", total, pointTotal-total)
				balance += bet * returnRatio
			} else if total > pointTotal {
				fmt.Printf("Dealer drew a "+newCard.String()+".\nYou lose your bet! Dealer got a score of %d and beat you by %d\n", total, total-pointTotal)
				balance -= bet
			} else { // Unsure if dealer wins in he case of a tie
				fmt.Printf("Dealer drew a "+newCard.String()+".\nDealer score is %d. Score is a tie and the game will push\n", total)
			}
			time.Sleep(time.Second * pauseLong)
			playGame(username)
		} else if total > 21 {
			fmt.Printf("Dealer drew a "+newCard.String()+".\nDealer busted with a score of %d and you win your bet\n", total)
			balance += bet * returnRatio
			time.Sleep(time.Second * pauseLong)
			playGame(username)
		} else {
			fmt.Printf("Dealer drew a " + newCard.String() + "\n")
		}
	}
}
