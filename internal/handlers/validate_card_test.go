package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateCard_Handle(t *testing.T) {
	type fields struct {
		ValidateCardService ValidateCardService
	}
	tests := []struct {
		name     string
		input    string
		fields   func(t *testing.T) fields
		wantCode int
		wantBody string
	}{
		{
			name:  "[Positive] Valid card",
			input: `{"card_number": "5237251624778133"}`,
			fields: func(t *testing.T) fields {
				ctrl := gomock.NewController(t)

				srv := NewMockValidateCardService(ctrl)

				srv.EXPECT().ValidateCardNumber("5237251624778133").Return(nil)

				return fields{
					ValidateCardService: srv,
				}
			},
			wantCode: http.StatusOK,
			wantBody: `{"message":"valid card number"}
`,
		},
		{
			name:  "[Negative] Invalid card",
			input: `{"card_number": "1234567891011121"}`,
			fields: func(t *testing.T) fields {
				ctrl := gomock.NewController(t)
				srv := NewMockValidateCardService(ctrl)
				srv.EXPECT().ValidateCardNumber("1234567891011121").Return(errors.New("invalid card"))
				return fields{
					ValidateCardService: srv,
				}
			},
			wantCode: http.StatusBadRequest,
			wantBody: `{"message":"invalid card number"}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//подготавливаю наш сервис
			f := tt.fields(t)
			v := NewValidateCard(f.ValidateCardService)

			//подготоваливаем контекст
			req := httptest.NewRequest(
				http.MethodPost,
				"/validate/card",
				strings.NewReader(tt.input),
			)
			req.Header.Add("content-type", "application/json")
			rec := httptest.NewRecorder()
			eCTX := echo.New().NewContext(req, rec)

			//вызывваем функцию
			err := v.Handle(eCTX)
			require.NoError(t, err)

			// проверяем что ответ и код ответа совпадают
			require.Equal(t, tt.wantCode, rec.Code)
			require.Equal(t, tt.wantBody, rec.Body.String())
		})
	}
}

//e := echo.New()
//req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//rec := httptest.NewRecorder()
//c := e.NewContext(req, rec)
//h := &handler{mockDB}
//
//// Assertions
//if assert.NoError(t, h.createUser(c)) {
//assert.Equal(t, http.StatusCreated, rec.Code)
//assert.Equal(t, userJSON, rec.Body.String())
//}
