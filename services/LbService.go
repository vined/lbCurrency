package services

import (
    "encoding/xml"
    "io/ioutil"
    "net/http"

    "github.com/vined/lbCurrency/models"
    "github.com/vined/lbCurrency/utils"
)

const LB_URL = "https://www.lb.lt/webservices/fxrates/FxRates.asmx/getFxRates?tp=EU"


func getLbRatesForDate(date string) (map[string]float32, error) {

    response, err := http.Get(LB_URL + "&dt=" + date)
    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    lbRates := models.LbRatesResponse{}
    err = xml.Unmarshal(responseData, &lbRates)

    if err != nil {
        return nil, err
    }

    return utils.ToRatesMap(lbRates), nil;
}

func GetRatesResponse(dateStr string) models.RatesResponse {

    if !utils.IsDate(dateStr) {
        return models.RatesResponse{"", nil, "Invalid date"}
    }

    ratesMap, err := getLbRatesForDate(dateStr)
    if err != nil {
        panic(err)
    }
    return models.RatesResponse{dateStr, ratesMap, ""}
}

