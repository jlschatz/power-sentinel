package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/power-sentinel/ddns"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1"))
	})

	ddns.NewDDNSService()

	log.Println("Server listening on port: 6660")
	http.ListenAndServe(":6660", r)

}

func SystemUpNotification() {
	values := map[string]string{"status": "up", "time": time.Now().String()}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Println(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Println(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	log.Println(res)
}
