package models

type RatesResponse struct {
    Date string `json:"date"`
    Rates map[string]float32 `json:"rates"`
    Error string `json:"error,omitempty"`
}
