package main

import (
	"sort"
)

func colorMap(hand []Card) map[int]int {
	returnMap := make(map[int]int)
	for _, card := range hand {
		val, ok := returnMap[card.couleur]
		if ok {
			returnMap[card.couleur] = val + 1
		} else {
			returnMap[card.couleur] = 1
		}
	}
	return returnMap
}

func valueMap(hand []Card) map[int]int {
	returnMap := make(map[int]int)
	for _, card := range hand {
		val, ok := returnMap[card.valeur]
		if ok {
			returnMap[card.valeur] = val + 1
		} else {
			returnMap[card.valeur] = 1
		}
	}
	return returnMap
}

func CompareHands(handA []Card, handB []Card) int {

	aValue := valueMap(handA)
	bValue := valueMap(handB)

	aColor := colorMap(handA)
	bColor := colorMap(handB)

	sort.Slice(handA, func(i, j int) bool {
		return handA[i].valeur < handA[j].valeur
	})
	sort.Slice(handB, func(i, j int) bool {
		return handB[i].valeur < handB[j].valeur
	})
	// suite couleur
	lastCard := handA[0]
	count, maxValueA := 1, 0
	for _, card := range handA[1:] {

		if card.valeur == lastCard.valeur+1 && lastCard.couleur == card.couleur {
			count += 1
		} else {
			count = 0
		}
		lastCard = card
		if count == 5 {
			maxValueA = lastCard.valeur
		}
	}
	lastCard = handB[0]
	count, maxValueB := 0, 0
	for _, card := range handB[1:] {
		if card.valeur == lastCard.valeur+1 && lastCard.couleur == card.couleur {
			count += 1
		} else {
			count = 0
		}
		lastCard = card
		if count == 5 {
			maxValueB = lastCard.valeur
		}
	}
	if maxValueA > maxValueB {
		return 1
	} else if maxValueB > maxValueA {
		return -1
	}

	// carré
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 4
		bPair := bOk && bCount == 4
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}

	// full
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aBrelan := aOk && aCount == 3
		bBrelan := bOk && bCount == 3
		for val2 := 12; val2 >= 0; val2-- {
			if val == val2 {
				continue
			}
			aCount, aOk = aValue[val2]
			bCount, bOk = bValue[val2]
			aPair := aOk && aCount == 2
			bPair := bOk && bCount == 2
			if (aBrelan && aPair) && !(bPair && bBrelan) {
				return 1
			}
			if !(aBrelan && aPair) && (bPair && bBrelan) {
				return -1
			}
		}
	}

	// couleur
	for c := 0; c < 4; c++ {
		aCount, aOk := aColor[c]
		bCount, bOk := bColor[c]
		aPair := aOk && aCount == 5
		bPair := bOk && bCount == 5
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}

	// suite
	for val := 8; val >= 0; val-- {
		aFollow := true
		bFollow := true
		for i := 0; i < 5; i++ {
			_, aOk := aValue[val+i]
			aFollow = aFollow && aOk
			_, bOk := bValue[val+i]
			bFollow = bFollow && bOk
		}
		if aFollow && !bFollow {
			return 1
		}
		if bFollow && !aFollow {
			return -1
		}
	}

	// brelan
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aBrelan := aOk && aCount == 3
		bBrelan := bOk && bCount == 3
		if aBrelan && !bBrelan {
			return 1
		}
		if bBrelan && !aBrelan {
			return -1
		}
	}

	// double paire
	aPairs := make([]int, 0)
	bPairs := make([]int, 0)
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 2
		bPair := bOk && bCount == 2
		if aPair {
			aPairs = append(aPairs, val)
		}
		if bPair {
			bPairs = append(bPairs, val)
		}
	}
	if len(aPairs) == 2 && len(bPairs) < 2 {
		return 1
	}
	if len(bPairs) == 2 && len(aPairs) < 2 {
		return -1
	}
	if len(aPairs) == 2 && len(bPairs) == 2 {
		if (aPairs[0] > bPairs[0] && aPairs[0] > bPairs[1]) || (aPairs[1] > bPairs[0] && aPairs[1] > bPairs[1]) {
			return 1
		}
		if (bPairs[0] > aPairs[0] && bPairs[0] > aPairs[1]) || (bPairs[1] > aPairs[0] && bPairs[1] > aPairs[1]) {
			return -1
		}
		return 0
	}

	// paire
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 2
		bPair := bOk && bCount == 2
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}

	// carte la plus haute
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 1
		bPair := bOk && bCount == 1
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}
	return 0
}
