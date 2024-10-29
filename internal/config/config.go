package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DataBaseConfig struct {
	Usuario string `yaml:"usuario"`
	Senha   string `yaml:"senha"`
	Host    string `yaml:"host"`
	Porta   int    `yaml:"porta"`
	Banco   string `yaml:"banco"`
}

type Config struct {
	ProdDataBase    DataBaseConfig `yaml:"database-producao"`
	RestoreDataBase DataBaseConfig `yaml:"database-restore"`
}

func CarregarConfiguracoes(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("[-] Erro ao carregar as configurações: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("[-] Erro ao decodificar o arquivo de configurações: %w", err)
	}

	return &config, nil
}
