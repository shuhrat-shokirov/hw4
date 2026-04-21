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

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatAmount(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatAmount(firstAmount))
	return text
}

func formatAmount(tiyin int) string {
	sum := tiyin / 100
	fraction := tiyin % 100

	formattedSum := formatThousands(sum)
	if fraction == 0 {
		return formattedSum
	}

	return formattedSum + "." + twoDigits(fraction)
}

func formatThousands(n int) string {
	if n == 0 {
		return "0"
	}

	digits := strconv.Itoa(n)
	result := ""

	for len(digits) > 3 {
		part := digits[len(digits)-3:]
		result = " " + part + result
		digits = digits[:len(digits)-3]
	}

	result = digits + result

	return result
}

func twoDigits(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}

	return strconv.Itoa(n)
}
