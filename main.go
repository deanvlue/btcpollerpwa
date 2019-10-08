package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func doSomething(s string) {
	log.Println("doing something", s)
}

func polling() {
	for {
		time.Sleep(1 * time.Minute)
		//go doSomething(" Bay number one")
		go getSymbolTrack("btc_mxn")
		go getSymbolTrack("ltc_mxn")
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

	var bitsoapi = "https://api.bitso.com/v3/ticker/?book"
	url := fmt.Sprintf("%s=%s", bitsoapi, symbol)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("The request has failed: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		log.Println(string(data))
	}

}

func main() {

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	// Poll for API.
	log.Println("Crypto HODLER ....")
	log.Println("Contacting API.")
	polling()
	log.Println("Bye.")
	//os.Exit(0)
}
