package database

import (
	"database/sql"
	"fmt"
	"log"

	"goreport/internal/config"

	go_ora "github.com/sijms/go-ora/v2"
)

func ConectarAoBanco(dataBaseConfig config.DataBaseConfig) *sql.DB {
	connString := go_ora.BuildUrl(
		dataBaseConfig.Host,
		dataBaseConfig.Porta,
		dataBaseConfig.Banco,
		dataBaseConfig.Usuario,
		dataBaseConfig.Senha,
		nil,
	)

	db, err := sql.Open("oracle", connString)
	if err != nil {
		log.Fatalf("[-] Erro ao conectar ao banco %s : %v", dataBaseConfig.Banco, err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("[-] Erro ao pingar o banco de dados: %v", err)
	}

	fmt.Printf("[+] Conexão com o banco %s estabelecida com sucesso\n", dataBaseConfig.Banco)
	return db
}
