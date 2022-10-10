package main

import (
	"assigment3/services"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	go services.UpdateWeather()
	http.HandleFunc("/", getWeather)
	http.ListenAndServe(":9000", nil)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	dataWeather, err := ioutil.ReadFile("data.json")
	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	statusWeather := services.StatusWeather{}
	errUn := json.Unmarshal(dataWeather, &statusWeather)
	if errUn != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	water := statusWeather.Status.Water
	wind := statusWeather.Status.Wind
	var statusWater string
	var statusWind string

	if water <= 5 {
		statusWater = "Aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	if wind <= 6 {
		statusWind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}

	resultWeather := services.ResultWeather{}
	resultWeather.Water = water
	resultWeather.Wind = wind
	resultWeather.StatusWater = statusWater
	resultWeather.StatusWind = statusWind

	tpl, errTmpl := template.ParseFiles("static/index.html")
	if errTmpl != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": errTmpl.Error(),
		})
		return
	}
	tpl.Execute(w, resultWeather)

}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
