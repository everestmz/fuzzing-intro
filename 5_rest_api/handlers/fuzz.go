package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func Fuzz(data []byte) int {
	if len(data) < 3 {
		// We'd like some data we could potentially actually encode
		return -1
	}

	fuzzPayload := &EncodeRequest{
		ToEncode: string(data),
	}

	// Build up a fake echo context, complete with user context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	bs, err := json.Marshal(fuzzPayload)
	if err != nil {
		panic(err)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bs))

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	EncodeHandler(ctx)
	return 0
}
