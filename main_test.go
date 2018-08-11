package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

// TestMain_Index test that / is reachable
func TestMain_Index(t *testing.T) {
	router := SetUpRouter()
	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestMain_ValidateGetProviders test our providers validator
func TestMain_ValidateGetProviders(t *testing.T) {
	// state is an allowed parameter
	params := make(map[string][]string, 0)
	state := make([]string, 1)
	state[0] = "AL"
	params["state"] = state
	err := ValidateGetProviders(params)
	assert.Nil(t, err)

	// we must have one parameter
	params1 := make(map[string][]string, 0)
	err = ValidateGetProviders(params1)
	assert.NotNil(t, err)

	// unknown parameters are not allowed
	params2 := make(map[string][]string, 0)
	s := make([]string, 1)
	s[0] = "AL"
	params2["badparam"] = s
	err = ValidateGetProviders(params2)
	assert.NotNil(t, err)

}
