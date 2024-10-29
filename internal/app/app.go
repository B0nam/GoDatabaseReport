package app

import (
	"fmt"
	"log"

	"goreport/internal/config"
	"goreport/internal/database"
	"goreport/internal/services"
)

func App() {
	fmt.Println(`[+] Script iniciado`)

	config, err := config.CarregarConfiguracoes("config.yml")
	if err != nil {
		log.Fatalf("[-] Erro ao carregar as configurações: %v", err)
	}

	prodDatabase := database.ConectarAoBanco(config.ProdDataBase)
	restoreDatabase := database.ConectarAoBanco(config.RestoreDataBase)

	if prodDatabase != nil && restoreDatabase != nil {
		serviceProd := services.NewService(prodDatabase)
		serviceRestore := services.NewService(restoreDatabase)

		contratos := serviceRestore.ValidateContracts()
		fmt.Printf("Resultados da validação de %s:\n", contratos.Nome)
		fmt.Printf("Produção: Count = %d, MaxId = %d\n", contratos.BaseProducao.Count, contratos.BaseProducao.MaxId)
		fmt.Printf("Restauração: Count = %d, MaxId = %d\n", contratos.BaseHomologacao.Count, contratos.BaseHomologacao.MaxId)
		fmt.Printf("Status: %v\n\n", contratos.Status)

		parcelas := serviceRestore.ValidateParcelas()
		fmt.Printf("Resultados da validação de %s:\n", parcelas.Nome)
		fmt.Printf("Produção: Count = %d, MaxId = %d\n", parcelas.BaseProducao.Count, parcelas.BaseProducao.MaxId)
		fmt.Printf("Restauração: Count = %d, MaxId = %d\n", parcelas.BaseHomologacao.Count, parcelas.BaseHomologacao.MaxId)
		fmt.Printf("Status: %v\n\n", parcelas.Status)

		funcionarios := serviceRestore.ValidateFuncionarios()
		fmt.Printf("Resultados da validação de %s:\n", funcionarios.Nome)
		fmt.Printf("Produção: Count = %d\n", funcionarios.BaseProducao.Count)
		fmt.Printf("Restauração: Count = %d\n", funcionarios.BaseHomologacao.Count)
		fmt.Printf("Status: %v\n", funcionarios.Status)

		contratosProd := serviceProd.ValidateContracts()
		fmt.Printf("Resultados da validação de %s na produção:\n", contratosProd.Nome)
		fmt.Printf("Produção: Count = %d, MaxId = %d\n", contratosProd.BaseProducao.Count, contratosProd.BaseProducao.MaxId)
		fmt.Printf("Restauração: Count = %d, MaxId = %d\n", contratosProd.BaseHomologacao.Count, contratosProd.BaseHomologacao.MaxId)
		fmt.Printf("Status: %v\n\n", contratosProd.Status)

		parcelasProd := serviceProd.ValidateParcelas()
		fmt.Printf("Resultados da validação de %s na produção:\n", parcelasProd.Nome)
		fmt.Printf("Produção: Count = %d, MaxId = %d\n", parcelasProd.BaseProducao.Count, parcelasProd.BaseProducao.MaxId)
		fmt.Printf("Restauração: Count = %d, MaxId = %d\n", parcelasProd.BaseHomologacao.Count, parcelasProd.BaseHomologacao.MaxId)
		fmt.Printf("Status: %v\n\n", parcelasProd.Status)

		funcionariosProd := serviceProd.ValidateFuncionarios()
		fmt.Printf("Resultados da validação de %s na produção:\n", funcionariosProd.Nome)
		fmt.Printf("Produção: Count = %d\n", funcionariosProd.BaseProducao.Count)
		fmt.Printf("Restauração: Count = %d\n", funcionariosProd.BaseHomologacao.Count)
		fmt.Printf("Status: %v\n", funcionariosProd.Status)
	}
}
