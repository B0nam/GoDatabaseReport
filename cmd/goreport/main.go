package main

import (
	"fmt"
	"goreport/internal/app"
)

func main() {
	fmt.Println(`
  ____       ____                       _   
 / ___| ___ |  _ \ ___ _ __   ___  _ __| |_ 
| |  _ / _ \| |_) / _ \ '_ \ / _ \| '__| __|
| |_| | (_) |  _ <  __/ |_) | (_) | |  | |_ 
 \____|\___/|_| \_\___| .__/ \___/|_|   \__|
                      |_|

Ferramenta para gerar relatórios do banco de dadaos de restore
	`)
	app.App()
}