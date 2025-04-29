package pkg

import (
	"errors"
	"strconv"
	"strings"
)

type CardNumberValidator struct {
}

func (c CardNumberValidator) ValidateCardNumber(cardNumber string) error {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	if cardNumber == "" {
		return errors.New("invalid empty card number")
	}

	cardNumberSlice := []int64{}

	for _, n := range cardNumber {

		numeric, err := strconv.ParseInt(string(n), 10, 64)
		if err != nil {
			return errors.New("card number is not digit")
		}
		cardNumberSlice = append(cardNumberSlice, numeric)
	}

	return checkSum(cardNumberSlice)
}
