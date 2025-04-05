package main

import "fmt"

type (
	PlayerClass interface {
		Attack()
	}

	Sorcerer struct {
		Spell string
	}

	Witch struct {
		Minion string
	}

	Huntress struct {
		Spear string
	}
)

func PlayerAttack(p PlayerClass) {
	p.Attack()
}

func (s Sorcerer) Attack() {
	fmt.Println("Sorcerer attacks with spell:", s.Spell)
}

func (w Witch) Attack() {
	fmt.Println("Witch attacks with minion:", w.Minion)
}

func (h Huntress) Attack() {
	fmt.Println("Huntress attacks with spear:", h.Spear)
}

func main() {
	// Creating instances of Sorcerer, Witch, and Huntress
	sorcerer := Sorcerer{
		Spell: "Fireball",
	}

	witch := Witch{
		Minion: "Bomb Thrower",
	}

	huntress := Huntress{
		Spear: "Longbow",
	}

	// Calling the Attack method for each class
	PlayerAttack(sorcerer)
	PlayerAttack(witch)
	PlayerAttack(huntress)
}
