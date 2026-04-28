package product

import (
	"fmt"
	"strconv"
	"strings"
)

// formatMoney преобразует тийины в красивую строку (например, 12 990 000.99)
func formatMoney(tiins int) string {
	sum := tiins / 100 // Целые сумы
	rem := tiins % 100 // Остаток (тийины)

	// Превращаем сумы в строку и добавляем пробелы каждые 3 символа
	s := strconv.Itoa(sum)
	var parts []string
	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		parts = append([]string{s[start:i]}, parts...)
	}
	formattedSum := strings.Join(parts, " ")

	// Если есть тийины, добавляем их через точку
	if rem > 0 {
		return fmt.Sprintf("%s.%02d", formattedSum, rem)
	}
	return formattedSum
}

// Converter выводит карточку товара в нужном формате
func Converter(name, brand string, price int, isAvailable bool, installments int) {
	// Рассчитываем ежемесячный платеж (цена / кол-во месяцев)
	monthly := price / installments

	fmt.Println("===== Alifshop =====")
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %s сум\n", formatMoney(price))
	fmt.Printf("В наличии: %v\n", isAvailable)
	fmt.Printf("Рассрочка: %d мес → %s сум/мес\n", installments, formatMoney(monthly))
	fmt.Println("====================")
}
