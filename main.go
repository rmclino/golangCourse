package main

import "fmt"

/* Coments */

var card string

func main() {
	// Teste Cards
	cards := []string{"Ace of Diamonds", newCard(), newCard()}
	cards = append(cards, "Six of Spades")
	//card = "Ace Of Spades" /* Virgula na inicialização da variavel */

	for index, item := range cards {
		fmt.Println(index, "-", item)
	}
	size := len(cards)
	for i := 0; i < size; i++ {
		fmt.Println(i, cards[i], size)
	}
	fmt.Println(index, item := range cards)

}

func newCard() string {

	return "Five of Diamonds"
}
