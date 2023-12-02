package main

type Currency struct {
	Name   string
	ExRate float64
}

// конструктор

func NewCurrency(name string, exRate float64) *Currency {
	return &Currency{
		Name:   name,
		ExRate: exRate,
	}
}

// set та get методи
func (currency *Currency) setName(name string) {
	switch name {
	case "$":
		fallthrough
	case "₴":
		currency.Name = name
	default:
		panic("Некоректна грошова одиниця")
	}
}
func (currency *Currency) getName() string {
	return currency.Name
}

func (currency *Currency) setExRate(exRate float64) {
	if exRate > 0 {
		currency.ExRate = exRate
	} else {
		panic("Некоректне значення курсу")
	}
}
func (currency *Currency) getExRate() float64 {
	return currency.ExRate
}

/* Метод переводу в гривню */

func (currency *Currency) ConvertToHryvnia(count float64) float64 {
	res := currency.getExRate() * count
	return res
}
