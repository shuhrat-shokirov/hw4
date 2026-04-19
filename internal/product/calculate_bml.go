package product

import (
	"fmt"
	"strconv"
)

const (
	duration  = 12
	tiinToSum = 100
)

func Calculate(amount int) int {
	return amount / duration
}

func PriceFormating(price int) string {
	tiin := price % tiinToSum
	sum := price / tiinToSum
	formatted := InsertingSpace(sum)
	if tiin%tiinToSum == 0 {
		return fmt.Sprintf("%.2d", sum)
	}
	return fmt.Sprintf("%s.%2d", formatted, tiin)
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
