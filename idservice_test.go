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
	InitDefaultNode()
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
	InitDefaultNode()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	GetHealth(w, req)

	assert.NotEmpty(t, w)
	assert.NotEmpty(t, w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"status\": \"ok\"}", w.Body.String())

}

func TestGetIdgen(t *testing.T) {
	InitDefaultNode()
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

func TestGetIdmeta(t *testing.T) {
	InitDefaultNode()

	req := httptest.NewRequest(http.MethodGet, "/idmeta", nil)
	w := httptest.NewRecorder()
	GetIdmeta(w, req)

	assert.NotEmpty(t, w)
	assert.NotEmpty(t, w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	strBody := w.Body.String()
	assert.Contains(t, strBody, "start_time")
	assert.Contains(t, strBody, "node_id")
	assert.Contains(t, strBody, "time_bits")
	assert.Contains(t, strBody, "node_bits")
	assert.Contains(t, strBody, "count_bits")

	var result map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.NotEmpty(t, result["start_time"])
	assert.NotNil(t, result["node_id"])
	assert.NotEmpty(t, result["time_bits"])
	assert.NotEmpty(t, result["node_bits"])
	assert.NotEmpty(t, result["count_bits"])
}

// func TestGetHealthHttptServer(t *testing.T) {
// 	expected := "dummy data"
// 	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, expected)
// 	}))
// 	defer svr.Close()
// 	c := NewClient(svr.URL)
// 	res, err := c.UpperCase("anything")
// 	if err != nil {
// 		t.Errorf("expected err to be nil got %v", err)
// 	}
// 	// res: expected\r\n
// 	// due to the http protocol cleanup response
// 	res = strings.TrimSpace(res)
// 	if res != expected {
// 		t.Errorf("expected res to be %s got %s", expected, res)
// 	}
// }

// func TestGetHealthError(t *testing.T) {
// 	router := gin.New()
// 	router.GET("/health", GetHealth)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, "{\"status\":\"ok\"}", w.Body.String())
// }
