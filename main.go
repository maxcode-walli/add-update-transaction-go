package p

import (
	"net/http"
)

//func main() {
//	fmt.Println("Running ...")
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		AddUpdateTransaction(w, r)
//	})
//
//	http.ListenAndServe(":8080", nil)
//}

func AddUpdateTransaction(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	method := req.Method

	if method == "POST" {
		handleAddTransaction(w, req)
	}
	if method == "PUT" {
		handleMarkLegitTransaction(w, req)
	}
}
