package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tidwall/gjson"
)

func main() {
	operationPtr := flag.String("operation", "api", "operation (ex: api|test)")

	flag.Parse()

	router := mux.NewRouter().StrictSlash(false)

	currentPath, err := os.Executable()
	if err != nil {
		log.Fatalf("could not get current executable path: %v", err.Error())
	}

	runningPathDir := path.Dir(currentPath)
	staticWebDirectory := runningPathDir + "/web_ui"
	log.Printf("staticWebDirectory : %v", staticWebDirectory)

	httpDir := http.Dir(staticWebDirectory)
	httpFileServer := http.FileServer(httpDir)
	httpFileHandler := http.StripPrefix("/home", httpFileServer)
	router.PathPrefix("/home").Handler(httpFileHandler)

	fmt.Println("operation:", *operationPtr)

	if *operationPtr == "" {
		fmt.Println("please provide -operation <operation>")
		return
	}

	corsObject := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},   // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	if *operationPtr == "test" {
		log.Println("hello world !")
	} else if *operationPtr == "api" {

		fmt.Println("StartingAPI")

		corsRouterHandler := corsObject.Handler(router)

		srv := &http.Server{
			Addr:    ":9000",
			Handler: corsRouterHandler,
		}

		router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
			_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		})

		/*
		   curl -H "accept:application/json" -H "content-type:application/json" -X GET http://localhost:9000/api/health 2>/dev/null | python -m json.tool
		*/

		router.HandleFunc("/api/v1/getData", HandlerGetData).Methods("GET")
		/*
		   curl -H "accept:application/json" -H "content-type:application/json" -X GET http://localhost:9000/api/v1/getData 2>/dev/null | python -m json.tool
		*/

		router.HandleFunc("/api/v1/processData", HandlerProcessData).Methods("POST")
		/*
		   curl -H "accept:application/json" -H "content-type:application/json" -d '{"length1": 15, "length2": 25}' -X POST http://localhost:9000/api/v1/processData 2>/dev/null | python -m json.tool
		*/

		log.Fatal(srv.ListenAndServe())

	} else {
		log.Printf("Nothing To Do. Exiting ...")
		os.Exit(0)
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("(%s) took (%v) to execute", what, time.Since(start))
	}
}

func logAndReturnError(msg string, err error) error {
	message := ""
	if err != nil {
		message = fmt.Sprintf("(error) : (%v) : %v", msg, err.Error())
		log.Println(message)
		return errors.New(message)
	}
	return nil
}

func generateRandomata() (data map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in generateRandomata() Error: %v", r)
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("defer : unknown panic in generateRandomata()")
			}
		}
	}()

	data = make(map[string]interface{})
	err = json.Unmarshal([]byte(RandomString), &data)
	if err != nil {
		return data, logAndReturnError("could not decode json string", err)
	}
	return data, nil
}

func HandlerGetData(w http.ResponseWriter, r *http.Request) {
	defer elapsed("__FUNC__: HandlerGetData")()
	returnResult := make(map[string]interface{})
	returnResult["error"] = ""
	returnResult["serverResponse"] = make(map[string]interface{})

	randomData, err := generateRandomata()
	if err != nil {
		returnResult["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
		return
	}
	returnResult["serverResponse"] = randomData
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&returnResult)
}

func HandlerProcessData(w http.ResponseWriter, r *http.Request) {
	defer elapsed("__FUNC__: HandlerProcessData")()
	errorMessage := ""
	returnResult := make(map[string]interface{})
	returnResult["error"] = ""
	returnResult["serverResponse"] = make([]interface{}, 0)

	result := make([]interface{}, 0)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage = fmt.Sprintf("Could not read request body !")
		log.Printf(errorMessage)
		returnResult["error"] = errorMessage
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
	}

	bodyStr := string(body)

	var myInterface interface{} = bodyStr
	_, ok := myInterface.(string)
	if !ok {
		errorMessage = fmt.Sprintf("HandlerProcessData : String Assertion Failed !")
		log.Printf(errorMessage)
		returnResult["error"] = errorMessage
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
		return
	}

	parsedJSON, ok := gjson.Parse(bodyStr).Value().(map[string]interface{})
	if !ok {
		errorMessage = fmt.Sprintf("HandlerProcessData: Could not convert to a map !")
		log.Printf(errorMessage)
		returnResult["error"] = errorMessage
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
		return
	}

	/* needed parameters */
	length1, ok := parsedJSON["length1"].(float64)
	if !ok {
		errorMessage = fmt.Sprintf("HandlerProcessData : 'length1' is missing in the http request !")
		log.Printf(errorMessage)
		returnResult["error"] = errorMessage
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
		return
	}
	log.Printf("HandlerProcessData : length1 : (%v)", length1)

	/* *************************************************************************************************** */

	length2, ok := parsedJSON["length2"].(float64)
	if !ok {
		errorMessage = fmt.Sprintf("HandlerProcessData : 'length2' is missing in the http request !")
		log.Printf(errorMessage)
		returnResult["error"] = errorMessage
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
		return
	}
	log.Printf("HandlerProcessData : length2 : (%v)", length2)

	/* *************************************************************************************************** */

	str1 := GenerateRandomString(int64(length1))
	str2 := GenerateRandomString(int64(length2))

	result = append(result, str1)
	result = append(result, str2)

	returnResult["serverResponse"] = result
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&returnResult)
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(lengthOfString int64) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, lengthOfString)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

const RandomString = `
{
    "_id": "6219dd7e3fd9fb007a4d6fe7",
    "about": "Laborum pariatur fugi",
    "address": "460 Dahlgreen Place, Fedora, Kansas, 9718",
    "age": 21,
    "balance": "$1,995.17",
    "company": "KAGE",
    "email": "bridgettegriffin@kage.com",
    "eyeColor": "blue",
    "favoriteFruit": "strawberry",
    "friends": [
        {
            "id": 0,
            "name": "Serena Davis"
        }
    ],
    "gender": "female",
    "greeting": "Hello, Bridgette Griffin! You have 5 unread messages.",
    "guid": "b4dd0bf4-dbff-44bd-a1af-a5a6fb44231a",
    "index": 0,
    "isActive": "false",
    "latitude": 12.471153,
    "longitude": 118.329655,
    "name": "Bridgette Griffin",
    "phone": "+1 (811) 510-3853",
    "picture": "http://placehold.it/32x32",
    "registered": "2018-07-28T02:25:46 +07:00",
    "tags": [
        "labore",
        "magna",
        "do"
    ]
}
`
