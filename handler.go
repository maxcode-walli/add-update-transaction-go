package p

import (
	"encoding/json"
	"log"
	"net/http"
	events "project.maxcode.net/tfs/Apollo/Blizzard/_git/walli-events.git"
)

func handleAddTransaction(w http.ResponseWriter, req *http.Request) {
	var transaction events.TransactionUpdatedEventV1

	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = publish(ProjectID, TestTopic, transaction, map[string]string{"pigeon.eventType": events.TransactionUpdatedEventType})
	if err != nil {
		log.Println("error")
	}

	w.WriteHeader(http.StatusOK)
}

func handleMarkLegitTransaction(w http.ResponseWriter, req *http.Request) {
	var score events.TransactionAnomalyScoreCalculatedEventV1

	err := json.NewDecoder(req.Body).Decode(&score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = publish(ProjectID, TransactionScoresTopic, score, map[string]string{"pigeon.eventType": events.TransactionAnomalyScoreCalculatedEventType})
	if err != nil {
		log.Println("error")
	}

	w.WriteHeader(http.StatusOK)
}
