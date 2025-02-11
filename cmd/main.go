package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	e := echo.New()
	e.POST("/", cardHandler)
	e.Logger.Fatal(e.Start(":1323"))

	var c Context
	c.Body = "хуй"

}

type CardDataRequest struct {
	CardNumber string `json:"card_number"`
}

type CardDataResponse struct {
	Message string `json:"message"`
}

func cardHandler(c echo.Context) error {
	var cardData CardDataRequest
	var err error
	body, err := io.ReadAll(c.Request().Body)
	fmt.Println(string(body))
	if err = c.Bind(&cardData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	valid := ValidateCardNumber(cardData.CardNumber)
	if valid == nil {
		return c.JSON(http.StatusOK, CardDataResponse{Message: "valid card number"})
	} else {
		return c.JSON(http.StatusBadRequest, CardDataResponse{Message: "invalid card number"})
	}
}

func ValidateCardNumber(cardNumber string) error {
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

func checkSum(num []int64) error {
	var sum int64
	j := 1
	for i := len(num) - 1; i >= 0; i-- {
		if j%2 == 0 {
			num[i] = num[i] * 2
			if num[i] >= 10 {
				num[i] = num[i]%10 + num[i]/10
			}
		}
		sum += num[i]
		j++
	}
	if sum%10 == 0 {
		return nil
	} else {
		return errors.New("invalid card number")
	}
}

type Context struct {
	Body string
}

func (c Context) Bind(i interface{}) error {
	err := json.Unmarshal([]byte(c.Body), i)
	if err != nil {
		return err
	}

	return nil
}
