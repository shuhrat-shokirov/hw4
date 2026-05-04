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
	text = strings.ReplaceAll(text, "{price}", formatMoney(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatMoney(firstAmount))
	return text
}
func formatMoney(amount int) string {
	sums := amount / 100
	tiyin := amount % 100
	s := strconv.Itoa(sums)

	var result []byte
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		result = append([]byte{s[i]}, result...)
		count++

		if count%3 == 0 && i != 0 {
			result = append([]byte{' '}, result...)
		}
	}
	formatted := string(result)

	if tiyin > 0 {
		return fmt.Sprintf("%s.%02d", formatted, tiyin)
	}
	return formatted
}
