package product

import (
	"fmt"
	"strconv"
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

// форматирование числа с пробелами
func formatNumber(n int) string {
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

// главная функция форматирования цены
func formatPrice(price int) string {
	sums := price / 100
	tiyin := price % 100

	main := formatNumber(sums)

	if tiyin == 0 {
		return main
	}

	return fmt.Sprintf("%s.%02d", main, tiyin)
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatPrice(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatPrice(firstAmount))
	return text
}
