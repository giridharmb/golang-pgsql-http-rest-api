package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
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

/*
Please Refer To This For Populating The Table With Rows:

https://github.com/giridharmb/PostgreSQL-README

*/

var (
	PGSQLHost string = "my-pgsql-host.company.com"
	PGSQLUser string = "testuser"
	PGSQLPass string = "testpassword"
	PGSQLDB   string = "test-db"
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

		router.HandleFunc("/api/v1/getDataFromPGSQL", HandlerGetDataFromPGSQLTable).Methods("GET")
		/*
		   curl -H "accept:application/json" -H "content-type:application/json" -X GET http://localhost:9000/api/v1/getDataFromPGSQL 2>/dev/null | python -m json.tool
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

/* ******************************* PGSQL Functions | start ********************************************** */

/*

Query table 't_random' - to see the column names

test-db=> select * from t_random limit 10;
 random_num |    random_float     |               md5
------------+---------------------+----------------------------------
        642 | 0.04866303424917717 | 4148034f7ff59deddce0507cf5df6136
        299 | 0.42088215715087784 | 50cf40c1cd51fa0c9bcda04d19118637
        452 |  0.5040329323515316 | d83b110f8c24dc876feb39e1d32fc9a9
        558 |   0.680895277194864 | bd5e6593d1133df5abbe2dcd7ed41b22
        436 |  0.9993317095718943 | b0e288ee1d4d5ea7c50995091222526c
        143 |  0.4972750511722168 | 1ecf469484a66c646f33f1ad26d2c120
        987 |  0.7588865512592768 | 04705f29ec44fede34940cd954acf20c
        337 |  0.9723930089149349 | fe0f5b228256a75beccfe7156af05616
        321 |   0.637456665167111 | 5522e09f9a97abd695bf821ce1eca750
        896 |  0.9731822082784944 | 34726f12c5bae77be4b3bb0293ccb38a
(10 rows)
*/

/*
DBRecord ...
*/
type DBRecord struct {
	RandomNumber int64   `json:"random_num"`   // table column name -> "random_num"
	RandomFloat  float64 `json:"random_float"` // table column name -> "random_float"
	MD5Hash      string  `json:"md5"`          // table column name -> "md5"
}

func GetDataFromPGSQLTable() ([]DBRecord, error) {
	defer elapsed("__FUNC__: GetDataFromPGSQLTable")()
	errorMessage := ""
	outputData := make([]DBRecord, 0)
	postgresConn := fmt.Sprintf("postgres://%v:%v@%v:5432/%v", PGSQLUser, PGSQLPass, PGSQLHost, PGSQLDB)
	db, err := pgx.Connect(context.Background(), postgresConn)
	if err != nil {
		errorMessage = fmt.Sprintf("unable to connect to database: %v", err.Error())
		log.Printf(errorMessage)
		return outputData, errors.New(errorMessage)
	}
	defer func() {
		_ = db.Close(context.Background())
	}()

	///////// READ MULTIPLE ROWS ////////////

	log.Printf("QUERYING ALL ROWS...")

	rows, err := db.Query(context.Background(), "select "+
		" random_num,"+
		" random_float,"+
		" md5"+
		" from t_random")
	if err != nil {
		errorMessage = fmt.Sprintf("could not perform a select from the table : %v", err.Error())
		log.Printf(errorMessage)
		return outputData, errors.New(errorMessage)
	}
	defer func() {
		rows.Close()
	}()

	counter := 0

	for rows.Next() {
		var rowData DBRecord
		err = rows.Scan(
			&rowData.RandomNumber,
			&rowData.RandomFloat,
			&rowData.MD5Hash,
		)
		if err != nil {
			errorMessage = fmt.Sprintf("got an error during row scan : %v", err.Error())
			log.Printf(errorMessage)
			return outputData, errors.New(errorMessage)
		}

		/* *********** pretty print the output **************** */
		vmDataByteArray, _ := json.MarshalIndent(rowData, "", "    ")
		fmt.Println(string(vmDataByteArray))

		outputData = append(outputData, rowData)

		counter++
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		errorMessage = fmt.Sprintf("got an error during row iteration : %v", err.Error())
		log.Printf(errorMessage)
		return outputData, errors.New(errorMessage)
	}

	log.Printf("after scanning all the rows, counter => %v", counter)

	return outputData, nil
}

/*
HandlerGetDataFromPGSQLTable ...
*/
func HandlerGetDataFromPGSQLTable(w http.ResponseWriter, r *http.Request) {
	//defer elapsed("__FUNC__: HandlerGetDataFromPGSQLTable")()
	returnResult := make(map[string]interface{})
	returnResult["error"] = ""
	returnResult["serverResponse"] = make([]interface{}, 0)
	data, err := GetDataFromPGSQLTable()
	if err != nil {
		returnResult["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&returnResult)
		return
	}
	returnResult["serverResponse"] = data
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&returnResult)
}

/* ******************************* PGSQL Functions | end ********************************************** */

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
