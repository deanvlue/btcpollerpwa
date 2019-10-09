package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

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

func getData(symbol string, top int) []CurrentPrices {
	db, err := sql.Open("sqlite3", "./db/cryptohistory")
	checkErr(err)

	if symbol == "" {
		symbol = "btc_mxn"
	}

	if top < 1 {
		top = 5
	}

	queryString := fmt.Sprintf("SELECT bid, created_at FROM tahistory WHERE book= \"%s\" ORDER BY created_at desc LIMIT %d", symbol, top)
	rows, err := db.Query(queryString)
	checkErr(err)

	var last5 []CurrentPrices
	var currentBid CurrentPrices

	for rows.Next() {
		err := rows.Scan(&currentBid.Bid, &currentBid.CreatedAt)
		checkErr(err)
		last5 = append(last5, currentBid)
		//fmt.Printf("%f, %s\n", bid, createdAt)
	}

	//fmt.Printf("%v", last5)
	rows.Close()

	return last5

}
