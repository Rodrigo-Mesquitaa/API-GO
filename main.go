package main

import (
	database "BankGo/database"
	"BankGo/routes"
)
/** 
Bootstrap
*/

func main() {
	// Iniciando conexão com o banco de dados
	database.Initialize()
	// Iniciando rotas
	routes.Initialize()
}