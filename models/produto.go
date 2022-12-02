package models

import "github.com/ewertonribeiro/go-loja/db"

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

func GetAllProducts() []Produto {
	db := db.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	products := []Produto{}

	for rows.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := rows.Scan(&id, &quantidade, &nome, &descricao, &preco)
		if err != nil {
			panic(err.Error())

		}

		p.Id = id
		p.Quantidade = quantidade
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco

		products = append(products, p)

	}

	return products
}
