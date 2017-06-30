package main

import (
	"fmt"
	"github.com/enjekt/panda/commons"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"net/http"
)

var db BlackvaultDB = NewBlackvaultDB()

type AddService int
type GetService int

func (self *AddService) Add(r *http.Request, msg *commons.TokenPadMessage, result *commons.Result) error {
	fmt.Println("Add method being called")
	fmt.Printf("Received message %s \n", msg.ToString())
	token := commons.InitToken(msg.Token)
	pad := commons.InitPad(msg.Pad)
	db.UpsertTokenPad(token, pad)
	*result = commons.Result{Response: "Received token/pad"}
	return nil
}

func (self *GetService) GetPad(r *http.Request, token *commons.Token, result *commons.Pad) error {
	fmt.Println("Get method being called")
	fmt.Printf("Received message %s \n", token.ToString())

	pad := db.GetPad(token)
	fmt.Printf("Result Pad %s \n", pad.ToString())
	*result = commons.Pad{}
	result.ID = pad.ToString()
	return nil
}

func startServer() {

	router := mux.NewRouter()
	/*	router.HandleFunc("/datastore/{token}/{pad}", AddPadForTokenEndpoint).Methods("PUT")
		router.HandleFunc("/pad/{token}", GetPadForTokenEndpoint).Methods("GET")
		router.HandleFunc("/token/{token}", DeleteTokenEndpoint).Methods("DELETE")*/

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(new(GetService), "GetService")
	s.RegisterService(new(AddService), "AddService")
	router.Handle("/rpc", s)
	http.ListenAndServe("localhost:10000", router)
}

func main() {
	startServer()
}

/*

func GetPadForTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	tokenParam := params["token"]
	fmt.Println("Token received lookup: ", tokenParam)
	//_ = json.NewDecoder(req.Body).Decode(&token)
	token := commons.InitToken(tokenParam)
	pad := db.GetPad(token)

	fmt.Println("This is the pad: ", pad)
	log.Println("This is the pad: ", pad)
}
func AddPadForTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	tokenParam := params["token"]
	padParam := params["pad"]
	fmt.Println("Token received lookup: ", tokenParam)
	//_ = json.NewDecoder(req.Body).Decode(&token)
	token := commons.InitToken(tokenParam)
	pad := commons.InitPad(padParam)

	db.UpsertTokenPad(token, pad)

}
func DeleteTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)

}
*/
