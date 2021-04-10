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
	respErr := domain.NewResponseError("A Bad request", http.StatusBadRequest)
	base.error(w, respErr)
	j, err := json.Marshal(respErr)
	assert.NoError(t, err)

	assert.EqualValues(t, w.Code, respErr.Sts)
	assert.JSONEq(t, string(j), w.Body.String())

	w = httptest.NewRecorder()
	regErr := errors.New("this ia standard regular error")
	base.error(w, regErr)

	assert.EqualValues(t, w.Code, http.StatusInternalServerError)
}

func TestBaseHandler_respond(t *testing.T) {
	base := baseHandler{}

	type fake struct {
		Id int `json:"id"`
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
