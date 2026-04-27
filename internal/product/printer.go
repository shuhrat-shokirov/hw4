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
	
	text = strings.ReplaceAll(text, "{price}", FormatPrice(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", FormatPrice(firstAmount))

	return text
}

func FormatPrice(tiyins int) string {
	// Делим на 100, чтобы получить сумы и остаток (тийины)
	sums := tiyins / 100
	rem := tiyins % 100

	sumStr := strconv.Itoa(sums)

	var parts []string
	for len(sumStr) > 3 {
		parts = append([]string{sumStr[len(sumStr)-3:]}, parts...)
		sumStr = sumStr[:len(sumStr)-3]
	}
	if len(sumStr) > 0 {
		parts = append([]string{sumStr}, parts...)
	}

	formattedSums := strings.Join(parts, " ")

	if rem > 0 {
		return fmt.Sprintf("%s.%02d", formattedSums, rem)
	}

	return formattedSums
}
