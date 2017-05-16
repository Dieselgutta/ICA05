package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

const URL = "http://api.openweathermap.org/data/2.5/weather?id=6453405&appid=2f134dd341ed0970e1bdf7bec5eac617"

type Coordinates struct {
	Lon float64
	Lat float64
}
type Measurements struct {
	Temp     float64
	Pressure float64
	Humidity float64
	Temp_min float64
	Temp_max float64
}
type Sys1 struct {
	Id      float64
	Message float64
	Country string
	Sunrise float64
	Sunset  float64
}
type Vind struct {
	Speed float64
	Deg   float64
}
type Weather struct {
	Coord Coordinates
	Main  Measurements
	Sys   Sys1
	Wind  Vind
}

func main() {
	http.HandleFunc("/", basicHandler)
	http.ListenAndServe(":8001", nil)
	log.Println("Listening...")
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	json := getData()
	profile := Weather(*json)
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decode(data []byte) *Weather {
	dat := new(Weather)
	err := json.Unmarshal(data, dat)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	dat.Main.Temp = dat.Main.Temp - 273.15
	return dat
}

func getData() *Weather {
	res, _ := http.Get(URL)
	body, _ := ioutil.ReadAll(res.Body)
	s := decode(body)
	return s
}
