package main

import (
	"encoding/json"
	"github.com/rs/cors"
	"math/rand"
	"net/http"
	"time"
)

type SensorData struct {
	Sunlight    float64 `json:"sunlight"`
	Humidity    float64 `json:"humidity"`
	PowerOutput float64 `json:"power_output"`
}

func generateData() SensorData {
	return SensorData{
		Sunlight:    rand.Float64() * 100,
		Humidity:    rand.Float64() * 100,
		PowerOutput: rand.Float64() * 1000,
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := generateData()
	json.NewEncoder(w).Encode(data)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", dataHandler)

	// CORS 설정 추가
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
