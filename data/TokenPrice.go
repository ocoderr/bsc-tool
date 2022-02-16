package data

import (
	"github.com/ocoderr/bsc-tool/config"
	"strings"
	"time"
)

var priceMap map[string]float64

func StartRefreshPrices() {
	priceMap = make(map[string]float64)

	tokenPrice := TokenPrice("bnb", Client)
	priceMap[strings.ToLower("bnb")] = tokenPrice

	for symbol, _ := range config.CF.BscToken {
		if symbol == "busd" || symbol == "usdt" {
			priceMap[strings.ToLower(symbol)] = 1.0
		} else {
			tokenPrice := TokenPrice(symbol, Client)
			priceMap[strings.ToLower(symbol)] = tokenPrice
		}
	}

	for {
		select {
		case <-time.After(10 * time.Second):
			tokenPrice := TokenPrice("bnb", Client)
			priceMap[strings.ToLower("bnb")] = tokenPrice

			for symbol, _ := range config.CF.BscToken {
				if symbol == "busd" || symbol == "usdt" {
					priceMap[strings.ToLower(symbol)] = 1.0
				} else {
					tokenPrice := TokenPrice(symbol, Client)
					priceMap[strings.ToLower(symbol)] = tokenPrice
				}
			}
		}
	}
}
func GetTokenPrice(token string) float64 {
	age, ok := priceMap[strings.ToLower(token)]
	if !ok {
		age = TokenPrice(token, Client)
	}
	return age
}
