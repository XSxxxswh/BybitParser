package github.com/XSxxxswh/BybitParser

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Request_params struct {
	Amount     string   `json:"amount"`
	AuthMaker  bool     `json:"authMaker"`
	CanTrade   bool     `json:"canTrade"`
	CurrencyId string   `json:"currencyId"`
	ItemRegion string   `json:"itemRegion"`
	Page       string   `json:"page"`
	Payment    []string `json:"payment"`
	Side       string   `json:"side"`
	Size       string   `json:"size"`
	TokenId    string   `json:"tokenId"`
	UserId     string   `json:"userId"`
}
type Res struct {
	Res Result `json:"result"`
}
type Result struct {
	Count int     `json:"count"`
	Items []Price `json:"items"`
}
type Price struct {
	Pr string `json:"price"`
}

func GetRate(currency string, amount string, size string) ([]float64, error) {
	params := Request_params{
		Amount:     amount,
		AuthMaker:  false,
		CanTrade:   false,
		CurrencyId: currency,
		ItemRegion: "2",
		Page:       "1",
		Payment:    []string{"585", "582"},
		Side:       "0",
		Size:       size,
		TokenId:    "USDT",
		UserId:     "",
	}
	var answer Res
	data, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api2.bybit.com/fiat/otc/item/online", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
	req.Header.Set("Lang", "ru-RU")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	json.Unmarshal(body, &answer)
	var a []float64
	for _, price := range answer.Res.Items {
		p, _ := strconv.ParseFloat(price.Pr, 64)
		a = append(a, p)
	}
	return a, nil
}
