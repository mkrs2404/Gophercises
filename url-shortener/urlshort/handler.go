package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		if path, ok := pathsToUrls[request.URL.Path]; ok {
			fmt.Println("Matched: ", path)
			http.Redirect(writer, request, path, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYaml(yml []byte) ([]map[string]string, error) {
	m := make([]map[string]string, 1, 1)
	err := yaml.Unmarshal(yml, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func buildMap(parsedYaml []map[string]string) map[string]string {
	finalMap := make(map[string]string)
	for _, m := range parsedYaml {
		key := m["path"]
		finalMap[key] = m["url"]
	}
	return finalMap
}
