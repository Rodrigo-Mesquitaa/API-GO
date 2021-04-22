package routes

import (
	"BankGo/domain"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

/**
Inicialização de rotas
*/

func Initialize() {
	router := mux.NewRouter()
	//obtém  a lista de transferencias
	router.HandleFunc("/transfers", domain.GetTransfers).Methods("GET")
	// faz transferencia de um Account para outro.
	router.HandleFunc("/transfers", domain.SendTransfer).Methods("POST")
	// obtém a lista de contas.
	router.HandleFunc("/accounts", domain.getAccounts).Methods("GET")
	// obtém o saldo da conta.
	router.HandleFunc("/accounts/{account_id}/balance", domain.GetBallance).Methods("GET")
	
	// criar uma conta.
	router.HandleFunc("/accounts", domain.CreateAccount).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
