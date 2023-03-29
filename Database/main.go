package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"	
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "azvdouglas:vasco@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {	
		panic(err)
	}
	product := NewProduct("Porta", 33.06)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

//Antes de inserir o produto, é necessário criar a tabela no banco de dados. Para isso, execute o seguinte comando no MySQL:
//CREATE TABLE products (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));
