package database

import (
	"fmt"
	"BankGo/models"
)

var db *gorm.DB

/**
 Inicialização do banco de dados
 */

 func Initialize() {
	 // Iniciando banco de dados.

	 var err error
	 dataSourceName := "root:123456@tcp(localhost:3306)/?parseTime=True"
	 db, err = gorm.Open("mysql", dataSourceName)

	 if err != nil {
		 fmt.Println(err)
		 panic("falha na conexão com o banco de dados")
	 }

	 // Criando banco de dados.
	 db.Exec("CREATE DATABASE BankGo_db")
	 db.Exec("USE BankGo_db")

	 // Migration para criação dos schemas.
	 db.AutoMigrate(&models.Account{}, &models.Transfer{})
 }