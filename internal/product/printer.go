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

func formatMoney(tiin int) string {
	sum := tiin / 100
	tiyin := tiin % 100

	intPart := formatThousands(sum)

	if tiyin == 0 {
		return intPart
	}

	return intPart + "." + twoDigits(tiyin)
}

func formatThousands(n int) string {
	s := strconv.Itoa(n)
	var result []byte
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		result = append([]byte{s[i]}, result...)
		count++

		if count%3 == 0 && i != 0 {
			result = append([]byte{' '}, result...)
		}
	}

	return string(result)
}

func twoDigits(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatMoney(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatMoney(firstAmount))
	return text
}
