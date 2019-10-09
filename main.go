package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func polling() {
	for {
		//interval := int64(*i)
		//log.Println(interval)
		time.Sleep(2 * time.Minute)
		//go doSomething(" Bay number one")
		go getSymbolTrack("btc_mxn")
		go getSymbolTrack("ltc_mxn")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, ❤️ %s", r.URL.Path[1:])
}

func handlerTracker(w http.ResponseWriter, r *http.Request) {
	symbols, ok := r.URL.Query()["sym"]

	if !ok || len(symbols[0]) < 1 {
		log.Println("Symbol needed to get information")
		return
	}

	symbol := symbols[0]
	w.Header().Set("Content-Type", "application/json")
	lastPrices := getData(symbol, 5)

	json.NewEncoder(w).Encode(lastPrices)

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

func checkErr(err error) {
	if err != nil {
		log.Fatal("Hubo un error: ", err)
	}
}

func main() {

	//intervalTime := flag.Int("i", 10, "Update Interval")
	//flag.Parse()

	/* TODO:
	1. Agregar banderas de tiempo
	2. Separar lógica de base de datos en otro package
	3. Separar lógica de consulta en otro package

	*/

	// Poll for API.
	log.Println("Crypto HODLER ....")
	log.Println("Contacting API.")
	go polling()
	//resjson, _ := json.Marshal(getData("btc_mxn", 15))
	//fmt.Println(string(resjson))
	http.HandleFunc("/", handler)
	http.HandleFunc("/track/", handlerTracker)

	log.Println("Server working on port 8080...")
	http.ListenAndServe(":8080", nil)
	log.Println("Bye.")
	//os.Exit(0)
}
