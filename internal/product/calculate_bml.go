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
	if tiin == 0 {
		return formatted
	}
	return fmt.Sprintf("%s.%02d", formatted, tiin)
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
