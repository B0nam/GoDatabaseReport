package app

import (
	"fmt"
	"goreport/internal/database"
	"goreport/config/config"
)

func App() {
	fmt.Println(`[+] Script iniciado`)

	config, err := config.CarregarConfiguracoes()
	if err != nil {
		log.Fatalf("[-] Erro ao carregar as configurações")
	}
	
	prodDatabase := database.ConectarAoBanco(config.ProdDataBase)
	restoreDatabase := database.ConectarAoBanco(config.RestoreDataBase)
}