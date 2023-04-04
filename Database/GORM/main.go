package main

import (
	"fmt"

	_ "github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   string	`gorm:"primary_key"`
	Name string
	Products []Product
}

type Product struct {
	ID    string	`gorm:"primary_key"`
	Name  string
	Price float64
	CategoryID string
	Category Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&Product{}, &Category{})

	// //Create category
	// category := Category{ ID: uuid.New().String(), Name: "Eletrônicos"}
	// db.Create(&category)

	// //Create product
	// product := Product{ ID: uuid.New().String(), Name: "Televisão", Price: 3333.06, CategoryID: category.ID}
	// db.Create(&product)

	/*
	products := []product{
		{ID: uuid.New().String(), Name: "Porta", Price: 33.06},
		{ID: uuid.New().String(), Name: "Janela", Price: 33.06},
		{ID: uuid.New().String(), Name: "Cadeira", Price: 33.06},
		{ID: uuid.New().String(), Name: "Mesa", Price: 33.06},
		{ID: uuid.New().String(), Name: "Cama", Price: 33.06},
		{ID: uuid.New().String(), Name: "Fogão", Price: 33.06},
		{ID: uuid.New().String(), Name: "Geladeira", Price: 777.07},
		{ID: uuid.New().String(), Name: "Microondas", Price: 33.06},
		{ID: uuid.New().String(), Name: "Televisão", Price: 33.06},
		{ID: uuid.New().String(), Name: "Notebook", Price: 2933.06},
	}
	db.Create(&products)

	//Alter batch
	db.Model(&product{}).Where("price = ?", 33.06).Update("price", 3306)
	//Delete record
	db.Delete(&product{}, "name = ?", "Televisão")
	*/

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Printf("Produto: %v\nPreço: %.2f\nCategoria: %v\n", product.Name, product.Price, product.Category.Name)
	}

}
//Um erro que cometi foi gerar um go.mod numa subpasta de onde já tinha um go.mod
