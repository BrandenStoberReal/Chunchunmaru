package main

import (
	"chunchunmaru/internal/macros"
	"chunchunmaru/internal/utilities"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func handleWebError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func handleWebErrorWithMessage(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusInternalServerError)
}

func main() {
	// Entrypoint
	log.Println("Welcome to Chunchunmaru!")
	log.Printf("Found %d words in words.txt\n", utilities.WordCount())
	log.Printf("Random word of the day: %s\n", utilities.RandomWord())

	// HTTP stuff. Higher handlers take priority
	http.HandleFunc("/config", utilities.AppConfig.ConfigSetAPI)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/", indexHandler)
	log.Printf("Listening on port %d", utilities.AppConfig.GetConfig().Port)
	log.Printf("Open http://localhost:%d in the browser", utilities.AppConfig.GetConfig().Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", utilities.AppConfig.GetConfig().Port), nil))
}

func apiHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// GET api methods
		writer.Header().Add("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/server/info":
			// Provides generic server info to the client
			reply := utilities.ApiBaseInfoReply{
				AppVersion: "1.0.0",
				Uptime:     uptime().Seconds(),
				Os:         cases.Title(language.English, cases.Compact).String(runtime.GOOS),
				Arch:       cases.Title(language.English, cases.Compact).String(runtime.GOARCH),
			}
			replybytes, marshalerr := json.Marshal(reply)
			if marshalerr != nil {
				log.Println("Error marshalling json ", marshalerr)
				handleWebError(writer, marshalerr)
				return
			}
			_, writeerr := writer.Write(replybytes)
			if writeerr != nil {
				log.Println("Error writing json ", writeerr)
				handleWebError(writer, writeerr)
				return
			}
			break
		case "/api/templates/info":
			// Provides misc template information to the client
			filenames := make([]string, 0)
			files, readerr := os.ReadDir("templates")
			if readerr != nil {
				log.Println("Error reading templates dir ", readerr)
				handleWebError(writer, readerr)
				return
			}
			for _, file := range files {
				filenames = append(filenames, file.Name())
			}
			reply := utilities.ApiTemplateInfoReply{
				Count:     len(files),
				FileNames: filenames,
			}
			replybytes, marshalerr := json.Marshal(reply)
			if marshalerr != nil {
				log.Println("Error marshalling json ", marshalerr)
				handleWebError(writer, marshalerr)
				return
			}
			_, writeerr := writer.Write(replybytes)
			if writeerr != nil {
				log.Println("Error writing json ", writeerr)
				handleWebError(writer, writeerr)
				return
			}
			break
		}
		break
	case http.MethodPost:
		// POST api methods
		switch request.URL.Path {
		case "/api/templates/upload":
			// Allow a client to upload a template to the web server for serving
			decoder := json.NewDecoder(request.Body)
			var data utilities.ApiUploadTemplateData
			decoderr := decoder.Decode(&data)
			if decoderr != nil {
				log.Println("Error decoding json ", decoderr)
				handleWebError(writer, decoderr)
				return
			}
			if data.FileName == "" && data.ContentBase64 == "" {
				decodedhtml, base64err := base64.StdEncoding.DecodeString(data.ContentBase64)
				if base64err != nil {
					log.Println("Error decoding base64 ", base64err)
					handleWebError(writer, base64err)
					return
				}
				writefilerr := os.WriteFile("templates/"+data.FileName, decodedhtml, 0644)
				if writefilerr != nil {
					log.Println("Error writing file ", writefilerr)
					handleWebError(writer, writefilerr)
					return
				}
				writer.Header().Add("Content-Type", "text/html")
				writer.Write([]byte("OK"))
			} else {
				handleWebErrorWithMessage(writer, "Both JSON fields must not be empty.")
				return
			}
			break
		case "/api/templates/delete":
			decoder := json.NewDecoder(request.Body)
			var data utilities.ApiDeleteTemplateData
			decoderr := decoder.Decode(&data)
			if decoderr != nil {
				log.Println("Error decoding json ", decoderr)
				handleWebError(writer, decoderr)
				return
			}
			if data.FileName != "" {
				if utilities.FileExists("templates/" + data.FileName) {
					delfileerr := os.Remove("templates/" + data.FileName)
					if delfileerr != nil {
						log.Println("Error deleting file ", delfileerr)
						handleWebError(writer, delfileerr)
						return
					}
					writer.Header().Add("Content-Type", "text/html")
					writer.Write([]byte("OK"))
				} else {
					handleWebErrorWithMessage(writer, "File does not exist.")
					return
				}
			} else {
				handleWebErrorWithMessage(writer, "JSON field \"fileName\" must not be empty.")
				return
			}
			break
		}
		break
	default:
		// Client used an unsupported request method
		handleWebErrorWithMessage(writer, "Unsupported method \""+request.Method+"\".")
		break
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api") || strings.HasPrefix(r.URL.Path, "/config") {
		return
	}
	config := utilities.AppConfig.GetConfig()
	log.Printf("Waiting with delay %fs\n", time.Duration(config.Delay).Seconds())
	time.Sleep(time.Duration(config.Delay))
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
