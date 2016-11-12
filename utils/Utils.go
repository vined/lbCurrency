package utils

import (
    "time"

    "github.com/vined/lbCurrency/models"
)

const LT_DATE_FORMAT = "2006-01-02"


func IsDate(dateStr string) bool {
    _, err := time.Parse(LT_DATE_FORMAT, dateStr)
    if err != nil {
        return false
    }
    return true
}

func ToRatesMap(rates models.LbRatesResponse) map[string]float32 {
    ratesMap := make(map[string]float32)

    for _, rate := range rates.LbRate {
        for _, currAmt := range rate.LbCurrencyAmount {
            if currAmt.LbCurrency != "EUR" {
                ratesMap[currAmt.LbCurrency] = currAmt.LbAmount
            }
        }
    }
    return ratesMap
}

