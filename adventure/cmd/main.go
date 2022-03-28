package main

import (
	"encoding/json"
	"github.com/mkrs2404/gophercises/adventure"
	"log"
	"net/http"
	"os"
)

func main() {
	data, err := os.ReadFile("story.json")
	if err != nil {
		panic(err)
	}
	story := make(adventure.Story)

	if err = json.Unmarshal(data, &story); err != nil {
		panic(err)
	}
	//tmp := `Hello World!`
	//t := template.Must(template.New("").Parse(tmp))
	//h := adventure.NewHandler(story, adventure.WithTemplate(t))
	h := adventure.NewHandler(story)
	log.Fatal(http.ListenAndServe("localhost:3030", h))
}
