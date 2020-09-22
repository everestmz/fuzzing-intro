package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EncodeRequest struct {
	ToEncode string `json:"to_encode"`
}

type EncodeResponse struct {
	Encoded string `json:"encoded"`
}

func EncodeHandler(c echo.Context) error {
	request := &EncodeRequest{}
	err := c.Bind(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &EncodeResponse{
		Encoded: RunLengthEncode(request.ToEncode),
	})
}
