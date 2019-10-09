package main

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

//CurrentPrices Estructur for the current price and time of creation
type CurrentPrices struct {
	Bid       float32 `json:"bid"`
	CreatedAt string  `json:"created_at"`
}
