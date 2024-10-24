package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"strconv"
)

type Dto struct {
	PersonalCards []string `json:"personal_cards"`
	CommonCards []string `json:"common_cards"`
}

func evaluate(w http.ResponseWriter, r *http.Request) {

	// parsing de la requête
	var inputs Dto
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Impossible de parser la requête")
		log.Println("Parsing error", r.Body)
		return
	}

	// parsing des cartes
	usedCards := make(map[string]bool)
	myCards := make([]Card, len(inputs.PersonalCards))
	for id, c := range inputs.PersonalCards {
		card, err := MakeCardFromString(c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Erreur lors du parsing de la carte: %s, %s", c, err)
			return
		} else if _, ok := usedCards[card.ToString()]; ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Carte duppliquée: %s", card.ToString())
		} else {
			myCards[id] = card
			usedCards[card.ToString()] = true
		}
	}
	if len(myCards) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Nombre de cartes personnelles différent de 2")
		return
	}

	commonCards := make([]Card, len(inputs.CommonCards))
	for id, c := range inputs.CommonCards {
		card, err := MakeCardFromString(c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Erreur lors du parsing de la carte: %s, %s", c, err)
			return
		} else if _, ok := usedCards[card.ToString()]; ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Carte duppliquée: %s", card.ToString())
		} else {
			commonCards[id] = card
			usedCards[card.ToString()] = true
		}
	}
	if len(commonCards) > 5 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Nombre de cartes communes supérieur à 5")
		return
	}

	projectId := r.URL.Query().Get("eval_count")
	i, err := strconv.Atoi(projectId)
	if err != nil {
		i = 5000
	}
	
	value := simulate(myCards, commonCards, i)
	log.Println("Cartes perso:", myCards, ", cartes communes: ", commonCards, ", proba: ", value)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%f", value)
}


func simulate(myCards []Card, commonCards []Card, evalCount int) float64 {
	target_count := float64(evalCount)
	games_won := 0.0
	for i:=0.0; i<target_count; i+=1.0 {
		gen := MakeCardGeneratorNonEmpty(append(myCards, commonCards...))
		opponentGame := []Card{gen.Next(), gen.Next()}
		completeCommon := make([]Card, 5-len(commonCards))
		for i, _ := range completeCommon {
			completeCommon[i] = gen.Next()
		}
		result := CompareHands(
			append(myCards, append(commonCards, completeCommon...)...),
			append(opponentGame, append(commonCards, completeCommon...)...))
		if result == 1 {
			games_won += 1.0
		}
	}
	return games_won / target_count
}


func main() {
	log.Println("Server starting")
	http.Handle("/evaluate", http.HandlerFunc(evaluate))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
