package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Currencies struct {
	Usd Currency
	Uah Currency
	Pln Currency
	Czk Currency
}

type Currency struct {
	Code        string
	AlphaCode   string
	NumericCode string
	Name        string
	Rate        float64
	Date        string
	InverseRate float64
}

func main() {
	amount, _ := strconv.ParseInt(os.Args[1], 10, 64)
	from := strings.ToLower(os.Args[2])
	to := strings.ToLower(os.Args[4])

	body := request()

	var currencies Currencies

	json.Unmarshal([]byte(body), &currencies)
	result := convert(from, to, amount, currencies)
	fmt.Printf("\n%d %s is %f in %s\n\n", amount, strings.ToUpper(from), result, strings.ToUpper(to))
}

func convert(from, to string, amount int64, c Currencies) float64 {
	fromCurrency := getCurrency(from, c)
	toCurrency := getCurrency(to, c)

	if to == "eur" {
		return float64(amount) * fromCurrency.InverseRate
	}

	if from == "eur" {
		return float64(amount) * toCurrency.Rate
	}

	return float64(amount) * fromCurrency.InverseRate * toCurrency.Rate
}

func request() []byte {
	res, err := http.Get("http://www.floatrates.com/daily/eur.json")
	if err != nil {
		log.Fatal(err)
	}
	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err)
	}

	return body
}

func getCurrency(c string, currencies Currencies) Currency {
	var currency Currency

	// TODO: very ugly, need to figure out dynamic lookup
	switch c {
	case "usd":
		return currencies.Usd
	case "uah":
		return currencies.Uah
	case "pln":
		return currencies.Pln
	case "czk":
		return currencies.Czk
	}

	return currency
}
