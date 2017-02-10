package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "path"

    "github.com/vined/lbCurrency/services"
)

func getRatesHandler(w http.ResponseWriter, r *http.Request) {

    responseJson, err := json.Marshal(
        services.GetRatesResponse(path.Base(r.URL.Path)))

    if err != nil {
        panic(err)
    }

    fmt.Fprintf(w, "%s", responseJson)
}

func main() {

    http.Handle("/", http.FileServer(http.Dir("./web/public/")))
    http.HandleFunc("/getRates/", getRatesHandler)
    http.ListenAndServe(":8001", nil)
}
