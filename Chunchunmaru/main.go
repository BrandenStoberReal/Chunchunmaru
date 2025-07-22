package main

import (
	"chunchunmaru/internal/macros"
	"chunchunmaru/internal/utilities"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	// Entrypoint
	log.Println("Welcome to Chunchunmaru!")
	log.Printf("Found %d words in words.txt\n", utilities.WordCount())
	log.Printf("Random word of the day: %s\n", utilities.RandomWord())

	// HTTP stuff
	http.HandleFunc("/config", utilities.AppConfig.ConfigSetAPI)
	http.HandleFunc("/", indexHandler)
	log.Printf("Listening on port %d", utilities.AppConfig.GetConfig().Port)
	log.Printf("Open http://localhost:%d in the browser", utilities.AppConfig.GetConfig().Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", utilities.AppConfig.GetConfig().Port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html, filename, _ := utilities.RandomHTMLFromDir("./templates")
	aggression := rand.Intn(101)
	log.Printf("Serving template: %s with aggression %d\n", filename, aggression)
	template, err := macros.BuildTemplate(filename, html)
	if err != nil {
		log.Printf("Error building template: %s", err)
		return
	}
	err = template.Execute(w, macros.TemplateInput{Aggression: aggression})
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}

	//if r.URL.Path != "/" {
	//http.NotFound(w, r)
	//return
	//}
}
