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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",       // 로컬 개발 환경
			"http://133.186.228.94:31030", // 배포 환경
		},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	http.ListenAndServe(":8080", handler)
}
