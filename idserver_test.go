package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthHttpt(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	GetHealth(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res == nil {
		t.Errorf("expected res to be nonnil, got nil")
	}
	if res.Status != "200 OK" {
		t.Errorf("expected %s got %v", "200 OK", res.Status)
	}
	if res.Body == nil {
		t.Errorf("expected res.Body to be nonnil, got nil")
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "{\"status\": \"ok\"}" {
		t.Errorf("expected %s got %v", "{\"status\": \"ok\"}", string(data))
	}

}

func TestGetHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	GetHealth(w, req)

	assert.NotEmpty(t, w)
	assert.NotEmpty(t, w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"status\": \"ok\"}", w.Body.String())

}

func TestGetIdgen(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/idgen", nil)
	w := httptest.NewRecorder()
	GetIdgen(w, req)

	assert.NotEmpty(t, w)
	assert.NotEmpty(t, w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "uid")
	var result map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.NotEmpty(t, result["uid"])
}

