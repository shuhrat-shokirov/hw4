package product

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	duration  = 12
	tiinToSum = 100
)

func Calculate(amount int) int {
	return amount / duration
}
func InsertingSpace(price int) string {
	s := strconv.Itoa(price)
	var result string
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += " "
		}
		result += string(ch)
	}
	return result
}
func PriceFormating(price int) string {
	tiin := price % 100
	sum := price / 100
	var sum2 int
	formatted := InsertingSpace(sum)
	if tiin == 0 || price == sum2 {
		return formatted
	}
	return fmt.Sprintf("%s.%02d", formatted, tiin)
}
func formatInput(input string) (string, error) {
	input = strings.ReplaceAll(input, " ", "")

	if strings.Contains(input, ".") {
		// Already in sums — just format the thousands part
		parts := strings.SplitN(input, ".", 2)
		n, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}
		return InsertingSpace(n) + "." + parts[1], nil
	}

	// No decimal — treat as tiyins, divide by 100
	tiyins, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}
	return PriceFormating(tiyins), nil
}
