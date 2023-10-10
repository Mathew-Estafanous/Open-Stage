package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseHandler_error(t *testing.T) {
	w := httptest.NewRecorder()
	respErr := fmt.Errorf("%w: bad request", domain.BadInput)
	respondWithError(w, respErr)

	assert.EqualValues(t, w.Code, http.StatusBadRequest)

	var resp ResponseError
	err := json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.NoError(t, err)
	assert.EqualValues(t, respErr.Error(), resp.Msg)
	assert.EqualValues(t, http.StatusBadRequest, resp.Sts)

	w = httptest.NewRecorder()
	regErr := errors.New("this ia standard regular respondWithError")
	respondWithError(w, regErr)

	assert.EqualValues(t, w.Code, http.StatusInternalServerError)
}

func TestBaseHandler_respond(t *testing.T) {
	type fake struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	w := httptest.NewRecorder()
	f := fake{1, "mat"}
	respondWithCode(w, http.StatusCreated, f)

	j, err := json.Marshal(f)
	assert.NoError(t, err)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())
}
