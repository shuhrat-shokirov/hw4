package product

import (
	"strconv"
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

func formatSum(n int) string {
	str := strconv.Itoa(n / 100)
	result := ""

	for i, ch := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += " "
		}
		result += string(ch)
	}
	res := n % 100
	if res > 0 {
		result += "."
		if res < 10 {
			result += "0"
		}
		result += strconv.Itoa(n % 100)
	}
	return result
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatSum(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatSum(firstAmount))
	return text
}
