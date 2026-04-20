package product

import (
	"fmt"
	"strings"
)

type Product struct {
	Name    string
	Brand   string
	Price   int
	InStock bool
}

const template = `===== Alifshop =====
Товар:    {name}
Бренд:    {brand}
Цена:     {price} сум
В наличии: {inStock}
Рассрочка: 12 мес → {firstAmount} сум/мес
====================`

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatMoney(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", fmt.Sprintf("%t", product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatMoney(firstAmount))
	return text
}

func formatMoney(amount int) string {
	sum := amount / 100
	tiyin := amount % 100
	s := addSpaces(fmt.Sprintf("%d", sum))
	if tiyin == 0 {
		return s
	}

	return fmt.Sprintf("%s.%02d", s, tiyin)
}

func addSpaces(s string) string {
	result := ""
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		result = string(s[i]) + result
		count++

		if count%3 == 0 && i != 0 {
			result = " " + result
		}
	}

	return result
}
