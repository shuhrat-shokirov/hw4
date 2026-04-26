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

func formatTiyin(amount int) string {
	sums := amount / 100
	tiin := amount % 100

	sumsStr := strconv.Itoa(sums)
	var intFormatted strings.Builder
	offset := len(sumsStr) % 3
	for i, ch := range sumsStr {
		if i != 0 && (i-offset)%3 == 0 {
			intFormatted.WriteByte(' ')
		}
		intFormatted.WriteRune(ch)
	}

	if tiin == 0 {
		return intFormatted.String()
	}

	tiinStr := strconv.Itoa(tiin)
	if len(tiinStr) == 1 {
		tiinStr = "0" + tiinStr
	}
	return intFormatted.String() + "." + tiinStr
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatTiyin(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", formatTiyin(firstAmount))
	return text
}
