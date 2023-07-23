package repository

import (
	"houx/models"
	"houx/utils"
	"io"
	"net/http"
	"strings"
)

func Get(url string) ([]models.Table, error) {
	//example urls:
	//https://www.noticiasagricolas.com.br/cotacoes/cafe/cafe-arabica-mercado-fisico-tipo-6-7
	//"https://books.toscrape.com/catalogue/a-light-in-the-attic_1000/index.html
	res, err := http.Get("https://pt.wikipedia.org/wiki/Alemanha")
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	convertedBody := strings.Replace(string(content), "\n", "", -1)
	tables := utils.ParseTable(convertedBody)
	return tables, nil
}
