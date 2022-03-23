package urlshort

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	fallbackResponse = "fallback"
	path             = "/test"
	url              = "https://abc.com"
)

func fallback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fallbackResponse)
}

var testData = []struct {
	testname     string
	path         string
	responseCode int
}{
	{
		testname:     "Fallback for unknown path",
		path:         "/unknown",
		responseCode: 200,
	},
	{
		testname:     "Redirects to found URL",
		path:         "/test",
		responseCode: http.StatusMovedPermanently,
	},
}

func TestMapHandler(t *testing.T) {
	pathToUrls := map[string]string{path: url}

	for _, test := range testData {
		t.Run(test.testname, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, test.path, nil)
			response := httptest.NewRecorder()

			mapHandler := MapHandler(pathToUrls, http.HandlerFunc(fallback))
			mapHandler(response, request)

			if response.Code != test.responseCode {
				t.Errorf("Expected %d Got %d", test.responseCode, response.Code)
			}
		})
	}
}

func TestYAMLHandler(t *testing.T) {

	yml := fmt.Sprintf(`
- path: %s
  url: %s
`, path, url)

	for _, test := range testData {
		t.Run(test.testname, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, test.path, nil)
			response := httptest.NewRecorder()

			yamlHandler, err := YAMLHandler([]byte(yml), http.HandlerFunc(fallback))
			if err != nil {
				t.Error("Could not create YAMLHandler")
			}

			yamlHandler(response, request)

			if response.Code != test.responseCode {
				t.Errorf("Expected %d Got %d", test.responseCode, response.Code)
			}
		})

	}

}
