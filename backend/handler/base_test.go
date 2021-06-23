package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/handle_err"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseHandler_error(t *testing.T) {
	base := baseHandler{}

	w := httptest.NewRecorder()
	respErr := fmt.Errorf("%w: bad request", domain.BadInput)
	base.error(w, respErr)

	assert.EqualValues(t, w.Code, http.StatusBadRequest)

	var resp handle_err.ResponseError
	err := json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.NoError(t, err)
	assert.EqualValues(t, respErr.Error(), resp.Msg)
	assert.EqualValues(t, http.StatusBadRequest, resp.Sts)

	w = httptest.NewRecorder()
	regErr := errors.New("this ia standard regular error")
	base.error(w, regErr)

	assert.EqualValues(t, w.Code, http.StatusInternalServerError)
}

func TestBaseHandler_respond(t *testing.T) {
	base := baseHandler{}

	type fake struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	w := httptest.NewRecorder()
	f := fake{1, "mat"}
	base.respond(w, http.StatusCreated, f)

	j, err := json.Marshal(f)
	assert.NoError(t, err)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())
}
