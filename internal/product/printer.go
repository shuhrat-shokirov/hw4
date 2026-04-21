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

func formatTiin(amount int) string {
	sum := amount / 100
	tiins := amount % 100

	sumsStr := strconv.Itoa(sum)
	var result []byte
	for i, ch := range sumsStr {
		remaining := len(sumsStr) - i - 1
		result = append(result, byte(ch))
		if remaining > 0 && remaining%3 == 0 {
			result = append(result, ' ')
		}
	}
	formatted := string(result)

	if tiins != 0 {
		formatted += fmt.Sprintf(".%02d", tiins)
	}

	return formatted
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatTiin(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatTiin(firstAmount))
	return text
}
