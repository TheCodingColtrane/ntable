package main

import (
	repository "houx/Repository"
)

func main() {
	repository.Get("https://www.noticiasagricolas.com.br/cotacoes/cafe/cafe-arabica-mercado-fisico-tipo-6-7")
}
