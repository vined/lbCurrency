package models

import (
    "encoding/xml"
)

type LbCurrencyAmount struct {
    LbCurrency string `xml:"Ccy"`
    LbAmount float32 `xml:"Amt"`
}

type LbRate struct {
    LbCurrencyAmount []LbCurrencyAmount `xml:"CcyAmt"`
}

type LbRatesResponse struct {
    XMLName xml.Name `xml:"FxRates"`
    LbRate []LbRate `xml:"FxRate"`
}

