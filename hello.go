package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CitiesResponse struct {
	Cities []string `json:"cities"`
}

func CityHandler(res http.ResponseWriter, req *http.Request) {
	citiesResponse := &CitiesResponse{
		Cities: []string{
			"San Francisco",
			"Amsterdam",
			"Berlin",
			"New York",
			"Tokyo",
			"Kyoto",
			"Osaka",
			"Nagasaki",
			"Naha",
			"London",
			"Paris",
			"Seoul",
			"Austin",
		},
	}
	data, _ := json.MarshalIndent(citiesResponse, "", "  ")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	log.Println("Listening on this host: http://localhost:5005")

	http.HandleFunc("/cities.json", CityHandler)
	//err := http.ListenAndServe(":5005", nil)
	//if err != nil {
	//	log.Fatal("Unable to listen on :5005: ", err)
	//}
}