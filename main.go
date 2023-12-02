package main

import "fmt"

func main() {
	currencyMilk := NewCurrency("$", 30) // створення через конструктор
	milk := new(Product)                 // створення через сеттери
	milk.setName("Молоко")
	milk.setPrice(35.70)
	milk.setCost(currencyMilk)
	milk.setQuantity(10)
	milk.setProducer("Наше молоко")
	milk.setWeight(11.4)

	milk.PrintProduct() // Вивід загальної інформації про товар:
	fmt.Printf("\nЦіна за одиницю товару в гривнях: %.2f", milk.GetPriceIn())
	fmt.Printf("\nЗагальна вартість всіх пачок молока: %.2f%s", milk.GetTotalPrice(), milk.getCost())
	fmt.Printf("\nЗагальна вага всіх пачок молока: %.2f\n", milk.GetTotalWeight())

	currencyChocolate := NewCurrency("₴", 1) // створення через конструктор
	chocolate := NewProduct("Шоколад", 57.11, currencyChocolate, 2, "Наш шоколад", 0.31)

	products := []*Product{milk, chocolate}
	fmt.Printf("\nВсі продукти:\n")
	PrintProducts(products)

	info, infoCost := GetProductsInfo(products)
	fmt.Printf("\nНайдешевший товар = %.2f%s\nНайдорожчий товар = %.2f%s\n", info["min"], infoCost["minCost"], info["max"], infoCost["maxCost"])

	fmt.Printf("\nВведіть характеристику продуктів:\n")
	productsArray := ReadProductsArray()
	PrintProductsMap(productsArray)
}
