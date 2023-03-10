package p

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleAddTransaction(w http.ResponseWriter, req *http.Request) {
	var transaction TransactionUpdatedEventV1

	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = publish(ProjectID, TestTopic, transaction, map[string]string{"pigeon.eventType": TransactionUpdatedEventType})
	if err != nil {
		log.Println("error")
	}

	w.WriteHeader(http.StatusOK)
}

func handleMarkLegitTransaction(w http.ResponseWriter, req *http.Request) {
	var score TransactionAnomalyScoreCalculatedEventV1

	err := json.NewDecoder(req.Body).Decode(&score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = publish(ProjectID, TransactionScoresTopic, score, map[string]string{"pigeon.eventType": TransactionAnomalyScoreCalculatedEventType})
	if err != nil {
		log.Println("error")
	}

	w.WriteHeader(http.StatusOK)
}
