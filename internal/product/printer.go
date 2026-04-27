package product

import (
	"fmt"
	"strconv"
	"strings"
)

func formatMoney(tiins int) string {
	sum := tiins / 100
	remainder := tiins % 100

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

	if remainder > 0 {
		return fmt.Sprintf("%s.%02d", formattedSum, remainder)
	}

	return formattedSum
}

func Converter(name, brand string, price int, isAvailable bool, installments int) {
	monthly := price / installments

	fmt.Println("===== Alifshop =====")
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %s сум\n", formatMoney(price))
	fmt.Printf("В наличии: %v\n", isAvailable)
	fmt.Printf("Рассрочка: %d мес → %s сум/мес\n", installments, formatMoney(monthly))
	fmt.Println("====================")
}
