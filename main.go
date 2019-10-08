package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Payload Response from Biso API
type Payload struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	CreatedAt string `json:"created_at"`
	Book      string `json:"book"`
	Volume    string `json:"volume"`
	Vwap      string `json:"vwap"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	Change24  string `json:"change_24"`
}

// ResponseBitso Bitso response structure
type ResponseBitso struct {
	Success bool    `json:"success"`
	Payload Payload `json:"payload"`
}

func doSomething(s string) {
	log.Println("doing something", s)
}

func polling() {
	for {
		time.Sleep(10 * time.Minute)
		//go doSomething(" Bay number one")
		getSymbolTrack("btc_mxn")
		getSymbolTrack("ltc_mxn")
	}
}

func startPolling2() {
	for {
		<-time.After(2 * time.Second)
		go doSomething(" Hello from Bay 2")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, ❤️ %s", r.URL.Path[1:])
}

func getSymbolTrack(symbol string) {
	log.Println("Asking:", symbol)

	var bitsoAPIURL = "https://api.bitso.com/v3/ticker/?book"
	url := fmt.Sprintf("%s=%s", bitsoAPIURL, symbol)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("The request has failed: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		payload := ResponseBitso{}
		json.Unmarshal(data, &payload)
		//log.Println(payload.Payload.Bid)
		saveData(&payload, symbol)

	}
}

func saveData(data *ResponseBitso, symbol string) {
	db, err := sql.Open("sqlite3", "./db/cryptohistory")
	insertPayload := "INSERT INTO tahistory(id,source,symbol,high,last_price,created_at, book, volume, vwap,low,ask, bid,change24) values (null,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(insertPayload)
	checkErr(err)

	high, err := strconv.ParseFloat(data.Payload.High, 64)
	checkErr(err)
	last, err := strconv.ParseFloat(data.Payload.Last, 64)
	checkErr(err)
	volume, err := strconv.ParseFloat(data.Payload.Volume, 64)
	checkErr(err)
	vwap, err := strconv.ParseFloat(data.Payload.Vwap, 64)
	checkErr(err)
	low, err := strconv.ParseFloat(data.Payload.Low, 64)
	checkErr(err)
	ask, err := strconv.ParseFloat(data.Payload.Ask, 64)
	checkErr(err)
	bid, err := strconv.ParseFloat(data.Payload.Bid, 64)
	checkErr(err)
	change24, err := strconv.ParseFloat(data.Payload.Change24, 64)
	checkErr(err)
	resp, err := stmt.Exec("bitso", symbol, high, last, data.Payload.CreatedAt, data.Payload.Book, volume, vwap, low, ask, bid, change24)
	checkErr(err)

	id, err := resp.LastInsertId()
	checkErr(err)

	log.Printf("Insertado %v", id)
	//log.Printf("Guardando lo que está en: %v ", data.Payload.Ask)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Hubo un error: ", err)
	}
}

func main() {

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	/* TODO:
	1. Agregar banderas de tiempo
	2. Separar lógica de base de datos en otro package
	3. Separar lógica de consulta en otro package

	*/

	// Poll for API.
	log.Println("Crypto HODLER ....")
	log.Println("Contacting API.")
	polling()
	log.Println("Bye.")
	//os.Exit(0)
}
