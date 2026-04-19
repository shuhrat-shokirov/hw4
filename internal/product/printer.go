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

const (
	template = `===== Alifshop =====
Товар:    {name}
Бренд:    {brand}
Цена:     {price} сум
В наличии: {inStock}
Рассрочка: 12 мес → {firstAmount} сум/мес
====================`
)

func formatSum(tiin int) string {
	sums := tiin / 100
	tiins := tiin % 100

	s := fmt.Sprintf("%d", sums)
	result := ""
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += " "
		}
		result += string(ch)
	}

	if tiins == 0 {
		return result
	}
	return fmt.Sprintf("%s.%02d", result, tiins)
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatSum(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", fmt.Sprintf("%v", product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatSum(firstAmount))
	return text
}
