package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All systems stable"))
	})

	// r.Post("/alert", func(w http.ResponseWriter, req *http.Request) {
	// 	decoder := json.NewDecoder(req.Body)
	// 	var a bot.Alert
	// 	err := decoder.Decode(&a)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	c := db.C("Alerts")
	// 	if err := c.Insert(a); err != nil {
	// 		w.Write([]byte(fmt.Sprintf("Failed reporting alert with the following error:\n %v", err.Error())))
	// 	}else{
	// 		w.Write([]byte("Successfully reported alert"))
	// 	}
	// })

	log.Println("Server listening on port: 6660")
	http.ListenAndServe(":6660", r)

}
