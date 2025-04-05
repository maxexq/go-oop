package main

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

func (s Sorcerer) Attack() {
	println("Sorcerer attacks with spell:", s.Spell)
}

func (w Witch) Attack() {
	println("Witch attacks with minion:", w.Minion)
}

func (h Huntress) Attack() {
	println("Huntress attacks with spear:", h.Spear)
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
	sorcerer.Attack()
	witch.Attack()
	huntress.Attack()
}
