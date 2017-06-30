package main

import (
	"encoding/json"
	"github.com/enjekt/panda/commons"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db WhitevaultDB = NewWhitevaultDB()

func GetPANforTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token := commons.InitToken(params["token"])
	log.Println("Token received lookup: ", token)

	panda := db.GetPanda(token)
	pad := GetPad(token)
	log.Printf("This is the panda:%s", panda.ID)
	log.Printf("This is the pad:%s", pad.ID)

	pan := commons.InitPan(panda, pad)
	log.Printf("This is the pan string value: %s \n", pan.ToString())

	json.NewEncoder(w).Encode(*pan)
}

func AddPANEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	pan := commons.NewPan(params["pan"])

	log.Println("Add PAN: " + pan.ToString())
	pad := commons.NewPad()
	panda := commons.NewPanda(pan, pad)
	token := commons.NewToken(pan)
	log.Println("Created new pad: ", pad)

	log.Println("Created new panda: ", panda)
	db.UpsertTokenPanda(token, panda)
	SendTokenAndPad(token, pad)
	log.Println("Returning new token: ", *token)
	json.NewEncoder(w).Encode(*token)

}

func DeleteTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token := commons.InitToken(params["token"])
	//TODO Add the delete logic to DB and to BlackVault call.
	json.NewEncoder(w).Encode(token)

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/pan/{token}", GetPANforTokenEndpoint).Methods("GET")
	router.HandleFunc("/token/{token}", DeleteTokenEndpoint).Methods("DELETE")
	router.HandleFunc("/pan/{pan}", AddPANEndpoint).Methods("PUT")
	//Start server and listen for errors...
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}*/
