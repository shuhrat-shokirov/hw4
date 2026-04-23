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

//func Converter(product Product, firstAmount int) string {
//	text := strings.ReplaceAll(template, "{name}", product.Name)
//	text = strings.ReplaceAll(text, "{brand}", product.Brand)
//	text = strings.ReplaceAll(text, "{price}", strconv.Itoa(product.Price))
//	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
//	text = strings.ReplaceAll(text, "{firstAmount}", strconv.Itoa(firstAmount))
//	return text
//}

func FormatPrice(amount int) string {
	sum := amount / 100
	tiyin := amount % 100
	if amount < 0 && tiyin != 0 {
		tiyin = -tiyin
	}

	s := strconv.Itoa(sum)
	var result []string

	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{s[start:i]}, result...)
	}
	formattedSum := strings.Join(result, " ")

	if tiyin != 0 {
		return fmt.Sprintf("%s.%02d", formattedSum, tiyin)
	}
	return formattedSum
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", FormatPrice(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", FormatPrice(firstAmount))

	return text
}
