package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// RootHandler handles the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Welcome to patrick webserver. Practice API"))
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// GetHandler handles the /get endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        response := map[string]string{"message": "Hello from patrickwebserver.com"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// PostHandler handles the /post endpoint
func PostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var data map[string]interface{}
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// PostFormHandler handles the /postform endpoint
func PostFormHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        formData := make(map[string]string)
        for key, values := range r.PostForm {
            formData[key] = values[0]
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(formData)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/", RootHandler)
    http.HandleFunc("/get", GetHandler)
    http.HandleFunc("/post", PostHandler)
    http.HandleFunc("/postform", PostFormHandler)

    port := 8000
    fmt.Printf("Server is listening on port %d\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
