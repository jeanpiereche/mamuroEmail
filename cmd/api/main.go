package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router *chi.Mux

func routers() *chi.Mux {
	router.Get("/api/emails", getEmails)

	return router
}

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	routers()
	http.ListenAndServe(":3000", Logger())
}

func getEmails(w http.ResponseWriter, r *http.Request) {
	//term := chi.URLParam(r, "term")
	termino := r.URL.Query().Get("term")
	var query = `{
        "search_type": "matchphrase",
        "query":
        {
            "term": "` + termino + `"
        },
        "from": 0,
        "max_results": 18,
        "_source": [
			"Subject","From","To","body", "Message-Id"
		]
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	respondwithJSON(w, http.StatusOK, string(body))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, msg)
}

func respondwithJSON(w http.ResponseWriter, code int, payload string) {
	//response, _ := json.Marshal(payload)
	fmt.Println(payload)
	fmt.Println("ejecutado")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write([]byte(payload))
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
