package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type DataBaseConfig struct {
	Usuario string `yml:usuario`
	Senha string `yml:senha`
	Host string `yml:host`
	Porta int `yml:porta`
	Banco string `yml:banco`
}

type Config struct {
	ProdDataBase DataBaseConfig `database-prod`
	RestoreDataBase DataBaseConfig `database-restore`
}

func CarregarConfiguracoes(filename string)(*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("[-] Erro ao carregar as configurações: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("[-] Erro ao decodificar o arquivo de configurações: %w", err)
	}

	return &config, nil
}
