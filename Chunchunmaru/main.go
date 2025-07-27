package main

import (
	"chunchunmaru/internal/macros"
	"chunchunmaru/internal/utilities"
	"database/sql"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mb-14/gomarkov"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var startTime time.Time
var database *sql.DB
var markovModel *gomarkov.Chain

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
	// DB
	db, err := utilities.OpenDatabase("./chunchunmaru.db")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to local database.")
		database = db
	}

	ipColumns := []string{"ip TEXT PRIMARY KEY", "queries INTEGER", "aggression INTEGER"}
	uaColumns := []string{"useragent TEXT PRIMARY KEY", "queries INTEGER", "aggression INTEGER"}
	utilities.CreateTable(database, utilities.SqlTable{
		Name:    "ipinfo",
		Columns: ipColumns,
	})

	utilities.CreateTable(database, utilities.SqlTable{
		Name:    "agentinfo",
		Columns: uaColumns,
	})

	// Markov
	if utilities.FileExists("model.json") {
		model, markerr := utilities.LoadMarkovModel()
		if markerr != nil {
			log.Fatal(markerr)
			return
		}

		markovModel = model
	}

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
			templatesDirSize, _ := utilities.DirSize("templates")
			reply := utilities.ApiTemplateInfoReply{
				Count:          len(files),
				FileNames:      filenames,
				TotalDiskUsage: templatesDirSize,
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
		case "/api/logging/queries/ip":
			table := utilities.SqlTable{
				Name:    "ipinfo",
				Columns: []string{"ip", "queries", "aggression"},
			}

			ipQueries, iperr := utilities.FetchAllIpInfo(database, &table)
			if iperr != nil {
				log.Println("Error fetching ip info ", iperr)
				handleWebError(writer, iperr)
				return
			}

			replybytes, marshalerr := json.Marshal(ipQueries)
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
		case "/api/logging/queries/useragent":
			table := utilities.SqlTable{
				Name:    "agentinfo",
				Columns: []string{"useragent", "queries", "aggression"},
			}

			uaQueries, iperr := utilities.FetchAllUserAgentInfo(database, &table)
			if iperr != nil {
				log.Println("Error fetching agent info ", iperr)
				handleWebError(writer, iperr)
				return
			}

			replybytes, marshalerr := json.Marshal(uaQueries)
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
		case "/api/logging/queries/info":

			ipTable := utilities.SqlTable{
				Name:    "ipinfo",
				Columns: []string{"ip", "queries", "aggression"},
			}
			uaTable := utilities.SqlTable{
				Name:    "agentinfo",
				Columns: []string{"useragent", "queries", "aggression"},
			}
			totalIpQueries, iperr := utilities.SumQueries(database, &ipTable)
			if iperr != nil {
				log.Println("Error summarizing ip info ", iperr)
				handleWebError(writer, iperr)
				return
			}

			totalUaQueries, uaerr := utilities.SumQueries(database, &uaTable)
			if uaerr != nil {
				log.Println("Error summarizing ua info ", uaerr)
				handleWebError(writer, uaerr)
				return
			}

			replybytes, marshalerr := json.Marshal(utilities.ApiQueryInfoReply{
				TotalQueries: totalIpQueries + totalUaQueries,
			})
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
		case "/api/markov/train":
			decoder := json.NewDecoder(request.Body)
			var data utilities.ApiMarkovTrainData
			decoderr := decoder.Decode(&data)
			if decoderr != nil {
				log.Println("Error decoding json ", decoderr)
				handleWebError(writer, decoderr)
				return
			}
			if data.Corpus != "" {
				chain := utilities.TrainMarkovModel(data.Corpus, 5, 2.0, markovModel)
				markovModel = chain
				utilities.SaveMarkovModel(markovModel)
				writer.Header().Add("Content-Type", "text/html")
				writer.Write([]byte("OK"))
			} else {
				handleWebErrorWithMessage(writer, "Corpus must not be empty.")
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

	// Variables
	clientip := strings.Split(r.RemoteAddr, ":")[0]
	userAgent := r.Header.Get("User-Agent")
	config := utilities.AppConfig.GetConfig()
	html, filename, _ := utilities.RandomHTMLFromDir("./templates")

	// Tables
	ipTable := utilities.SqlTable{
		Name:    "ipinfo",
		Columns: []string{"ip", "queries", "aggression"},
	}
	uaTable := utilities.SqlTable{
		Name:    "agentinfo",
		Columns: []string{"useragent", "queries", "aggression"},
	}

	// SQL code
	ipQueries, iperr := utilities.FetchSingleValue[int](database, &ipTable, "queries", "ip", clientip)
	if iperr == sql.ErrNoRows {
		ipQueries = 0
		log.Printf("No record found for IP %s, defaulting queries to %d\n", clientip, ipQueries)
	} else if iperr != nil {
		log.Println("Database error:", iperr)
	} else {
		log.Printf("Queries for IP %s: %d\n", clientip, ipQueries)
	}

	uaQueries, uaerr := utilities.FetchSingleValue[int](database, &uaTable, "queries", "useragent", userAgent)
	if uaerr == sql.ErrNoRows {
		uaQueries = 0
		log.Printf("No record found for User-Agent %s, defaulting queries to %d\n", userAgent, uaQueries)
	} else if uaerr != nil {
		log.Println("Database error:", uaerr)
	} else {
		log.Printf("Queries for User-Agent %s: %d\n", userAgent, uaQueries)
	}

	ipValues := []interface{}{clientip, ipQueries + 1, (ipQueries + 1) / config.QueriesPerAggression}
	ipuperr := utilities.UpsertRow(database, ipTable, ipValues)
	if ipuperr != nil {
		log.Println(ipuperr)
		return
	}

	uaValues := []interface{}{userAgent, uaQueries + 1, (uaQueries + 1) / config.QueriesPerAggression}
	uauperr := utilities.UpsertRow(database, uaTable, uaValues)
	if uauperr != nil {
		log.Println(uauperr)
		return
	}

	// Aggression code
	var templateAggression int
	if (uaQueries+1)/config.QueriesPerAggression > (ipQueries+1)/config.QueriesPerAggression {
		// UA has higher aggression level
		templateAggression = (uaQueries + 1) / config.QueriesPerAggression
	} else if (uaQueries+1)/config.QueriesPerAggression < (ipQueries+1)/config.QueriesPerAggression {
		// IP has higher aggression level
		templateAggression = (ipQueries + 1) / config.QueriesPerAggression
	} else if (uaQueries+1)/config.QueriesPerAggression == (ipQueries+1)/config.QueriesPerAggression {
		// Both have the same aggression, default to IP
		templateAggression = (ipQueries + 1) / config.QueriesPerAggression
	}

	// Website delay
	randomDelay := utilities.RandomDuration(time.Duration(config.MinDelay), time.Duration(config.MaxDelay))
	log.Printf("Waiting with delay %fs\n", randomDelay.Seconds())
	time.Sleep(randomDelay)

	log.Printf("Serving template: %s with aggression %d\n", filename, templateAggression)
	template, err := macros.BuildTemplate(filename, html)
	if err != nil {
		log.Printf("Error building template: %s", err)
		return
	}
	err = template.Execute(w, macros.TemplateInput{Aggression: templateAggression})
	if err != nil {
		if strings.Contains(err.Error(), "An established connection was aborted by the software in your host machine.") {
			log.Println("Error executing template: Client browser aborted the request.")
		} else {
			log.Printf("Error executing template: %s", err)
		}
	}

	//if r.URL.Path != "/" {
	//http.NotFound(w, r)
	//return
	//}
}
