package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandler(t *testing.T) {
	t.Parallel()
	routes := map[string]string{
		"/dogs": "http://www.dogster.com/",
		"/cats": "http://www.catster.com/",
	}
	mapHandle := MapHandler(routes, http.HandlerFunc(http.NotFound))
	req, _ := http.NewRequest("GET", "http://nowherespecial.com/dogs", nil)
	w := httptest.NewRecorder()
	mapHandle.ServeHTTP(w, req)
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected a redirect")
	}
	if w.Header().Get("Location") != routes["/dogs"] {
		t.Errorf("Expected a redirect to the correct location")
	}
	req, _ = http.NewRequest("GET", "http://nowherespecial.com/ferrets", nil)
	w = httptest.NewRecorder()
	mapHandle.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected a NotFound")
	}
}
