package web

import (
	"cyoa/cyoa"
	"cyoa/templates"
	"cyoa/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Parser *flag.FlagSet = flag.NewFlagSet("web", flag.ExitOnError)
var file *string = Parser.String("file", "gopher.json", "the JSON file with the CYOA story")

func handleIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Index")
}

func Run() {
	fmt.Printf("Using the story in %s \n", *file)

	f, err := os.Open(*file)

	if err != nil {
		utils.HandleError("Failed to read file %s, error: %s", *file, err)
	}

	story, storyCreationError := cyoa.NewStory(f)

	if storyCreationError != nil {
		utils.HandleError(storyCreationError.Error())
	}

	handler := cyoa.NewHandler(*story, cyoa.WithTemplate(templates.StyledTemplate))

	log.Print("Server running on localhost:3000")

	http.ListenAndServe(":3000", handler)

	fmt.Printf("%+v", story)
}
