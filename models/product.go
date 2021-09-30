package models

import "github.com/antunesgabriel/crud/configs"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func (*Product) GetAll() ([]Product, error) {
	products := []Product{}
	product := Product{}

	database, err := configs.ConnectDatabase()

	if err != nil {
		return products, err
	}

	result, err := database.Query("SELECT * FROM products")

	if err != nil {
		return products, err
	}

	for result.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := result.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			continue
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}

	return products, nil
}
