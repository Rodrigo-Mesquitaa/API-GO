package domain

import (
	"encoding/json"
	"BankGo/models"
	"net/http"
)

var db *gorm.DB 

/**
// Retorna todas as transferencias já realizadas
*/

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transfers []models.Transfer
	db.Preload("Transfer").Find(&transfers)
	json.NewEncoder(w).Encode(transfers)
}

/**
Criar account
*/

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.account
	json.NewEncoder(r.Body).Decode(&account)
	db.Create(&account)
	w.?Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

// Todo

func SendTransfer() {}

func getAccounts() {}

func GetBallance() {}