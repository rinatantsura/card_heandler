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
