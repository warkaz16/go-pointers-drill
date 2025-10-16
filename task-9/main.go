package main

import "fmt"

type StudentCard struct {
	Owner  string
	Active bool
}

func issue(card *StudentCard, owner string) { // заполняет Owner и включает Active
	card.Owner = owner
	card.Active = true
}

func block(card *StudentCard) { // ставит Active = false
	card.Active = false
}

func rename(card *StudentCard, newOwner string) {
	card.Owner = newOwner
}

func main() {

	card := StudentCard{}
	fmt.Println("До:", card)

	issue(&card, "Алиев Али")
	fmt.Println("После issue:", card)

	rename(&card, "Интукод Студент")
	fmt.Println("После rename:", card)

	block(&card)
	fmt.Println("После block:", card)
}
