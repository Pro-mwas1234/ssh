package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }).Methods("GET")

    // User endpoints
    router.HandleFunc("/api/users", usersHandler).Methods("GET", "POST")
    // Transaction endpoints
    router.HandleFunc("/api/transactions", transactionsHandler).Methods("GET", "POST")
    // Auth endpoints
    router.HandleFunc("/api/auth/login", loginHandler).Methods("POST")
    router.HandleFunc("/api/auth/register", registerHandler).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Starting server on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}