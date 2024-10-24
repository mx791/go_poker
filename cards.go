package main

import (
	"strings"
	"fmt"
	"errors"
	"math/rand"
)

var valeurs []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Valet", "Dame", "Roi", "As"}
var couleurs []string = []string{"Piques", "Coeur", "Trefle", "Carreau"}

type Card struct {
	valeur int
	couleur int
}

func (c1 Card) Equals(c2 Card) bool {
	return c1.valeur == c2.valeur && c1.couleur == c2.couleur
}

func (c Card) ToString() string {
	return fmt.Sprintf("%s de %s", valeurs[c.valeur], couleurs[c.couleur])
}

func MakeCardFromString(cardName string) (Card, error) {
	cardName = " " + strings.ToLower(cardName) + " "
	valeur := -1
	for id, name := range valeurs {
		if strings.Contains(cardName, " " + strings.ToLower(name) + " ") {
			valeur = id
		}
	}
	if valeur == -1 {
		return Card{0,0}, errors.New("Impossible de parser la valeur")
	}
	for id, name := range couleurs {
		if strings.Contains(cardName, " " + strings.ToLower(name) + " ") {
			return Card{valeur, id}, nil
		}
	}
	return Card{0,0}, errors.New("Impossible de parser la couleur")
}

type CardGenerator struct {
	passed map[string]bool
}

func MakeCardGenerator() CardGenerator {
	return CardGenerator{make(map[string]bool)}
}

func MakeCardGeneratorNonEmpty(cards []Card) CardGenerator {
	usedMap := make(map[string]bool)
	for _, card := range cards {
		usedMap[card.ToString()] = true
	}
	return CardGenerator{usedMap}
}

func genCard() Card {
	return Card{rand.Intn(13), rand.Intn(4)}
}

func (gen *CardGenerator) Next() Card {
	card := genCard()
	cardFound, ok := gen.passed[card.ToString()]
	for cardFound && ok {
		card = genCard()
		cardFound, ok = gen.passed[card.ToString()]
	}
	gen.passed[card.ToString()] = true
	return card
}

func (gen *CardGenerator) Remove(card Card) {
	gen.passed[card.ToString()] = false
}
