package handler

import (
	"encoding/json"
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseHandler_error(t *testing.T) {
	base := baseHandler{}

	w := httptest.NewRecorder()
	respErr := domain.ApiError{Msg: "Bad request", Typ: domain.BadInput}
	base.error(w, respErr)

	assert.EqualValues(t, w.Code, http.StatusBadRequest)

	expectedResp, err := json.Marshal(newResponseError(respErr.Msg, http.StatusBadRequest))
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedResp), w.Body.String())

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
