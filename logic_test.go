package main

import (
	"testing"
)

func TestMakeCard(t *testing.T) {
	card, err := MakeCardFromString("Dame de Carreau")
	if err != nil || card.valeur != 10 {
		t.Fatalf("Erreur de parsing")
	}

	card, err = MakeCardFromString("5 de trefle")
	if err != nil || card.valeur != 3 {
		t.Fatalf("Erreur de parsing")
	}
}

func TestPairComparator(t *testing.T) {
	cardA, _ := MakeCardFromString("5 de trefle")
	cardB, _ := MakeCardFromString("5 de carreau")

	cardC, _ := MakeCardFromString("8 de trefle")
	cardD, _ := MakeCardFromString("7 de carreau")

	if CompareHands([]Card{cardA, cardB}, []Card{cardC, cardD}) != 1 {
		t.Fatalf("Erreur avec la paire")
	}

	cardC, _ = MakeCardFromString("9 de trefle")
	cardD, _ = MakeCardFromString("9 de carreau")
	if CompareHands([]Card{cardA, cardB}, []Card{cardC, cardD}) != -1 {
		t.Fatalf("Erreur avec la paire")
	}
}

func TestBrelanComparator(t *testing.T) {
	cardA, _ := MakeCardFromString("5 de trefle")
	cardB, _ := MakeCardFromString("5 de carreau")
	cardC, _ := MakeCardFromString("5 de carreau")
	cardD, _ := MakeCardFromString("8 de trefle")

	if CompareHands([]Card{cardA, cardB, cardC}, []Card{cardC, cardA, cardD}) != 1 {
		t.Fatalf("Le brelan doit être > paire")
	}

	cardD, _ = MakeCardFromString("As de trefle")
	cardE, _ := MakeCardFromString("As de carreau")
	cardF, _ := MakeCardFromString("As de carreau")

	if CompareHands([]Card{cardA, cardB, cardC}, []Card{cardD, cardE, cardF}) != -1 {
		t.Fatalf("Le brelan d'As doit etre > au brelan de 5")
	}
}

func TestQintFlushComparator(t *testing.T) {

	cardA, _ := MakeCardFromString("5 de trefle")
	cardB, _ := MakeCardFromString("4 de trefle")
	cardC, _ := MakeCardFromString("3 de trefle")
	cardD, _ := MakeCardFromString("2 de trefle")
	cardE, _ := MakeCardFromString("6 de trefle")

	if CompareHands([]Card{cardA, cardB, cardC, cardD, cardE}, []Card{cardC, cardD, cardE}) != 1 {
		t.Fatalf("Erreur avec le quint flush 1")
	}

	cardF, _ := MakeCardFromString("As de trefle")
	cardG, _ := MakeCardFromString("As de carreau")
	cardH, _ := MakeCardFromString("As de piques")
	cardI, _ := MakeCardFromString("As de coeur")

	if CompareHands([]Card{cardA, cardB, cardC, cardD, cardE}, []Card{cardF, cardG, cardH, cardI}) != 1 {
		t.Fatalf("Erreur avec le quint flush")
	}
}

func TestSimulate(t *testing.T) {
	cardA, _ := MakeCardFromString("5 de trefle")
	cardB, _ := MakeCardFromString("As de trefle")
	proba := simulate([]Card{cardA, cardB}, []Card{}, 5000)
	if proba < 0.5 || proba > 0.60 {
		t.Fatalf("Erreur de calcul, %f", proba)
	}

	cardA, _ = MakeCardFromString("As de coeur")
	cardC, _ := MakeCardFromString("As de Carreau")
	proba = simulate([]Card{cardA, cardB}, []Card{cardC}, 5000)
	if proba < 0.9 {
		t.Fatalf("Erreur de calcul, %f", proba)
	}

	proba = simulate([]Card{cardA, cardB}, []Card{}, 5000)
	if proba < 0.8 || proba > 0.9 {
		t.Fatalf("Erreur de calcul, %f", proba)
	}
}
