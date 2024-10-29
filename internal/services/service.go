package services

import (
	"database/sql"
	"log"
	"time"

	"goreport/internal/models"
)

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) ValidateContracts() models.Tabelas {
	var tabela models.Tabelas
	tabela.Nome = "Contratos"

	dateBeforeRestore := "2024-07-28"
	queryMaxID := "SELECT MAX(ID) FROM CONTRATO WHERE DT_CONTRATO < TO_DATE($1, 'YYYY-MM-DD')"
	err := s.DB.QueryRow(queryMaxID, dateBeforeRestore).Scan(&tabela.BaseHomologacao.MaxId)
	if err != nil {
		log.Printf("[-] Erro ao obter Max ID de Contratos na base de restore: %v", err)
		return tabela
	}

	queryCount := "SELECT COUNT(*) FROM CONTRATO WHERE ID <= $1"
	err = s.DB.QueryRow(queryCount, tabela.BaseHomologacao.MaxId).Scan(&tabela.BaseHomologacao.Count)
	if err != nil {
		log.Printf("[-] Erro ao contar contratos na base de restore: %v", err)
		return tabela
	}

	err = s.DB.QueryRow(queryCount, tabela.BaseHomologacao.MaxId).Scan(&tabela.BaseProducao.Count)
	if err != nil {
		log.Printf("[-] Erro ao contar contratos na base de produção: %v", err)
		return tabela
	}

	tabela.Status = tabela.BaseProducao.Count <= tabela.BaseHomologacao.Count
	return tabela
}

func (s *Service) ValidateParcelas() models.Tabelas {
	var tabela models.Tabelas
	tabela.Nome = "Parcelas"

	queryMaxID := "SELECT MAX(ID) FROM CONTRATO_PARCELAS"
	err := s.DB.QueryRow(queryMaxID).Scan(&tabela.BaseHomologacao.MaxId)
	if err != nil {
		log.Printf("[-] Erro ao obter Max ID de Parcelas na base de restore: %v", err)
		return tabela
	}

	queryCount := "SELECT COUNT(*) FROM CONTRATO_PARCELAS WHERE ID < $1"
	err = s.DB.QueryRow(queryCount, tabela.BaseHomologacao.MaxId).Scan(&tabela.BaseHomologacao.Count)
	if err != nil {
		log.Printf("[-] Erro ao contar parcelas na base de restore: %v", err)
		return tabela
	}

	err = s.DB.QueryRow(queryCount, tabela.BaseHomologacao.MaxId).Scan(&tabela.BaseProducao.Count)
	if err != nil {
		log.Printf("[-] Erro ao contar parcelas na base de produção: %v", err)
		return tabela
	}

	tabela.Status = tabela.BaseProducao.Count <= tabela.BaseHomologacao.Count
	return tabela
}

func (s *Service) ValidateFuncionarios() models.Tabelas {
	var tabela models.Tabelas
	tabela.Nome = "Funcionários"

	queryLastDate := "SELECT DT_CADASTRO FROM FUNCIONARIO_DADOS_MARGEM WHERE ID = (SELECT MAX(ID) FROM FUNCIONARIO_DADOS_MARGEM)"
	var lastDate time.Time
	err := s.DB.QueryRow(queryLastDate).Scan(&lastDate)
	if err != nil {
		log.Printf("[-] Erro ao obter última data de cadastro de Funcionários: %v", err)
		return tabela
	}

	queryCount := "SELECT COUNT(*) FROM FUNCIONARIO_DADOS_MARGEM WHERE DT_CADASTRO < TO_DATE($1, 'YYYY-MM-DD HH24:MI:SS')"
	err = s.DB.QueryRow(queryCount, lastDate.Format("2006-01-02 15:04:05")).Scan(&tabela.BaseHomologacao.Count)
	if err != nil {
		log.Printf("[-] Erro ao contar funcionários na base de restore: %v", err)
		return tabela
	}

	err = s.DB.QueryRow(queryCount, lastDate.Format("2006-01-02 15:04:05")).Scan(&tabela.BaseProducao.Count)
	if err != nil {
		log.Printf("[-] Erro ao contar funcionários na base de produção: %v", err)
		return tabela
	}

	tabela.Status = tabela.BaseProducao.Count <= tabela.BaseHomologacao.Count
	return tabela
}
