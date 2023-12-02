package main

import (
	"fmt"
	"math"
)

type Product struct {
	Name     string
	Price    float64
	Cost     *Currency
	Quantity uint64
	Producer string
	Weight   float64
}

// конструктор

func NewProduct(name string, price float64, cost *Currency, quantity uint64, producer string, weight float64) *Product {
	return &Product{
		Name:     name,
		Price:    price,
		Cost:     cost,
		Quantity: quantity,
		Producer: producer,
		Weight:   weight,
	}
}

// set та get методи
func (product *Product) setName(name string) {
	if len(name) > 1 && len(name) < 255 {
		product.Name = name
	} else {
		panic("Некоректна назва продукту")
	}
}
func (product *Product) getName() string {
	return product.Name
}

func (product *Product) setPrice(price float64) {
	if price > 0 {
		product.Price = price
	} else {
		panic("Некоректне значення ціни")
	}
}
func (product *Product) getPrice() float64 {
	return product.Price
}

func (product *Product) setCost(cost *Currency) {
	product.Cost = cost
}
func (product *Product) getCost() string {
	return product.Cost.Name
}

func (product *Product) setQuantity(quantity uint64) {
	if quantity > 0 {
		product.Quantity = quantity
	} else {
		panic("Некоректне значення кількості")
	}
}
func (product *Product) getQuantity() uint64 {
	return product.Quantity
}

func (product *Product) setProducer(producer string) {
	if len(producer) > 1 && len(producer) < 255 {
		product.Producer = producer
	} else {
		panic("Некоректне значення назви компанії")
	}
}
func (product *Product) getProducer() string {
	return product.Producer
}

func (product *Product) setWeight(weight float64) {
	if weight > 0 {
		product.Weight = weight
	} else {
		panic("Некоректне значення ваги")
	}
}
func (product *Product) getWeight() float64 {
	return product.Weight
}

func (product *Product) GetPriceIn() float64 {
	if product.Cost.Name == "₴" {
		return product.getPrice()
	} else {
		return product.Cost.ConvertToHryvnia(product.getPrice())
	}
}
func (product *Product) GetTotalPrice() float64 {
	res := product.getPrice() * float64(product.getQuantity())
	return res
}
func (product *Product) GetTotalWeight() float64 {
	res := product.getWeight() * float64(product.getQuantity())
	return res
}
func (product *Product) PrintProduct() {
	fmt.Printf("\n*************\nІнформація про `%s`\nЦіна за 1 одиницю: %.2f%s\nВ наявності: %d (одиниць товару)\nВиробник: `%s`\nВага за 1 одиницю: %.2f\n*************\n", product.getName(), product.getPrice(), product.getCost(), product.getQuantity(), product.getProducer(), product.getWeight())
}
func PrintProducts(products []*Product) {
	for _, p := range products {
		fmt.Println(p)
		p.PrintProduct()
	}
}
func PrintProductsMap(products map[int]*Product) {
	for _, p := range products {
		fmt.Println(p)
		p.PrintProduct()
	}
}
func GetProductsInfo(products []*Product) (map[string]float64, map[string]string) {
	info := make(map[string]float64, 2)
	infoCost := make(map[string]string, 2)
	var (
		maxPriceHryvnia = float64(math.MinInt64)
		minPriceHryvnia = float64(math.MaxInt64)

		maxPrice = float64(math.MinInt64)
		minPrice = float64(math.MaxInt64)

		minCost string
		maxCost string
	)
	for _, i := range products {
		if i.GetPriceIn() > maxPriceHryvnia {
			maxPriceHryvnia = i.GetPriceIn()
			maxPrice = i.getPrice()
			maxCost = i.getCost()
		}
		if i.GetPriceIn() < minPriceHryvnia {
			minPriceHryvnia = i.GetPriceIn()
			minPrice = i.getPrice()
			minCost = i.getCost()
		}
	}
	info["min"] = minPrice
	info["max"] = maxPrice
	infoCost["minCost"] = minCost
	infoCost["maxCost"] = maxCost

	return info, infoCost
}
func ReadProductsArray() map[int]*Product {
	continueInputValues := true
	currencies := make(map[int]*Currency)
	products := make(map[int]*Product)
	var (
		name     string
		price    float64
		cost     string
		exRate   float64
		quantity uint64
		producer string
		weight   float64
		add      = "n"
	)
	for i := 0; continueInputValues == true; i++ {
		fmt.Printf("Введіть назву продукту:_")
		fmt.Scanln(&name)
		fmt.Printf("Введіть ціну (без валюти для однієї одиниці товару):_")
		fmt.Scanln(&price)
		fmt.Printf("Введіть валюту, в якій буде зберігатись ціна продукту ($ або ₴):_")
		fmt.Scanln(&cost)
		fmt.Printf("Введіть курс для цієї валюти:_")
		fmt.Scanln(&exRate)
		fmt.Printf("Введіть кількість продуктів:_")
		fmt.Scanln(&quantity)
		fmt.Printf("Введіть назву компанії-виробника (без пробілів):_")
		fmt.Scanln(&producer)
		fmt.Printf("Введіть вагу (для однієї одиниці товару):_")
		fmt.Scanln(&weight)
		fmt.Printf("Чи бажаєте Ви додати ще товар (y/n):_")
		fmt.Scanln(&add)
		currencies[i] = new(Currency)
		currencies[i].setName(cost)
		currencies[i].setExRate(exRate)

		products[i] = new(Product)
		products[i].setName(name)
		products[i].setPrice(price)
		products[i].setCost(currencies[i])
		products[i].setQuantity(quantity)
		products[i].setProducer(producer)
		products[i].setWeight(weight)

		if add == "n" {
			continueInputValues = false
			break
		}
	}
	return products
}
