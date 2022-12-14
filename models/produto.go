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

	rows, err := db.Query("SELECT * FROM produtos ORDER BY id DESC")
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

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	defer db.Close()

	insertDb, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDb.Exec(nome, descricao, preco, quantidade)

}

func DeleteProduct(id string) {
	db := db.ConnectDB()
	defer db.Close()

	deleteProduct, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

}

func GetProductById(id string) Produto {
	db := db.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Produto{}

	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := rows.Scan(&id, &quantidade, &nome, &descricao, &preco)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Quantidade = quantidade
		product.Nome = nome
		product.Descricao = descricao
		product.Preco = preco
	}
	return product
}

func UpdateProduct(id, nome, descricao, preco, quantidade string) {
	db := db.ConnectDB()
	defer db.Close()

	update, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(nome, descricao, preco, quantidade, id)

}
