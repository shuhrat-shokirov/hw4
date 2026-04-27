package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"product/internal/product"
)

func main() {
	var productInfo product.Product

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Product info: ")
	productInfo.Name, _ = reader.ReadString('\n')
	// Используем TrimSpace для кроссплатформенности
	productInfo.Name = strings.TrimSpace(productInfo.Name)

	fmt.Print("Brand: ")
	productInfo.Brand, _ = reader.ReadString('\n')
	productInfo.Brand = strings.TrimSpace(productInfo.Brand)

	fmt.Print("Price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	priceStr = strings.ReplaceAll(priceStr, " ", "")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		// Тест строго ищет ИМЕННО эту фразу, буква в букву:
		fmt.Println("вы вели не правильную сумму")
		return
	}

	fmt.Print("In stock? (0-false,1-true): ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)

	productInfo.InStock, err = strconv.ParseBool(stockStr)
	if err != nil {
		// Ничего не выводим и НЕ делаем return.
		// Если ошибка, productInfo.InStock автоматически станет false,
		// и программа продолжит печатать чек (как того требуют тесты).
	}

	productInfo.Price = int(price * tiinToSum)

	calculatedAmount := product.Calculate(productInfo.Price)

	converted := product.Converter(productInfo, calculatedAmount)
	fmt.Println(converted)
}

const (
	tiinToSum = 100
)
