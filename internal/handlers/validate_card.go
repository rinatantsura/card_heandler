//go:generate mockgen -source=validate_card.go -destination=validate_card_mock_test.go -package=handlers
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ValidateCardService interface {
	ValidateCardNumber(cardNumber string) error
}

func NewValidateCard(service ValidateCardService) ValidateCard {
	return ValidateCard{
		ValidateCardService: service,
	}
}

type ValidateCard struct {
	ValidateCardService
}

type ValidateCardRequest struct {
	CardNumber string `json:"card_number"`
}

type ValidateCardResponse struct {
	Message string `json:"message"`
}

func (v ValidateCard) Handle(eCTX echo.Context) error {
	var cardData ValidateCardRequest
	var err error
	if err = eCTX.Bind(&cardData); err != nil {
		return eCTX.String(http.StatusBadRequest, err.Error())
	}
	valid := v.ValidateCardNumber(cardData.CardNumber)
	if valid == nil {
		return eCTX.JSON(http.StatusOK, ValidateCardResponse{Message: "valid card number"})
	} else {
		return eCTX.JSON(http.StatusBadRequest, ValidateCardResponse{Message: "invalid card number"})
	}
}

//func (v ValidateCard) ValidateCardNumber(cardNumber string) error {
//	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
//
//	if cardNumber == "" {
//		return errors.New("invalid empty card number")
//	}
//
//	cardNumberSlice := []int64{}
//
//	for _, n := range cardNumber {
//
//		numeric, err := strconv.ParseInt(string(n), 10, 64)
//		if err != nil {
//			return errors.New("card number is not digit")
//		}
//		cardNumberSlice = append(cardNumberSlice, numeric)
//	}
//
//	return checkSum(cardNumberSlice)
//}
//
//func checkSum(num []int64) error {
//	var sum int64
//	j := 1
//	for i := len(num) - 1; i >= 0; i-- {
//		if j%2 == 0 {
//			num[i] = num[i] * 2
//			if num[i] >= 10 {
//				num[i] = num[i]%10 + num[i]/10
//			}
//		}
//		sum += num[i]
//		j++
//	}
//	if sum%10 == 0 {
//		return nil
//	} else {
//		return errors.New("invalid card number")
//	}
//}
//
