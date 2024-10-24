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


func PairComparatorTest(t *testing.T) {
	cardA, _ := MakeCardFromString("5 de trefle")
	cardB, _ := MakeCardFromString("5 de carreau")

	cardC, _ := MakeCardFromString("8 de trefle")
	cardD, _ := MakeCardFromString("7 de carreau")

	if CompareHands([]Card{cardA,cardB}, []Card{cardC,cardD}) != 1 {
		t.Fatalf("Erreur avec la paire")
	}

	cardC, _ = MakeCardFromString("9 de trefle")
	cardD, _ = MakeCardFromString("9 de carreau")
	if CompareHands([]Card{cardA,cardB}, []Card{cardC,cardD}) != -1 {
		t.Fatalf("Erreur avec la paire")
	}
}


func BrelanComparatorTest(t *testing.T) {
	cardA, _ := MakeCardFromString("5 de trefle")
	cardB, _ := MakeCardFromString("5 de carreau")
	cardC, _ := MakeCardFromString("5 de carreau")

	cardD, _ := MakeCardFromString("8 de trefle")
	cardE, _ := MakeCardFromString("7 de carreau")
	cardF, _ := MakeCardFromString("6 de carreau")

	if CompareHands([]Card{cardA,cardB}, []Card{cardC,cardD}) != 1 {
		t.Fatalf("Erreur avec le brelan")
	}

	cardD, _ = MakeCardFromString("8 de trefle")
	cardE, _ = MakeCardFromString("8 de carreau")
	cardF, _ = MakeCardFromString("6 de carreau")

	if CompareHands([]Card{cardA,cardB}, []Card{cardC,cardD}) != 1 {
		t.Fatalf("Erreur avec le brelan")
	}

	cardD, _ = MakeCardFromString("As de trefle")
	cardE, _ = MakeCardFromString("As de carreau")
	cardF, _ = MakeCardFromString("As de carreau")

	if CompareHands([]Card{cardA,cardB,cardC}, []Card{cardD,cardE,cardF}) != 1 {
		t.Fatalf("Erreur avec le brelan")
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
	if proba < 0.85 || proba > 0.9 {
		t.Fatalf("Erreur de calcul, %f", proba)
	}
}