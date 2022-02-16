package data

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/ocoderr/bsc-tool/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
	"time"
)

type Token struct {
	Client      *ethclient.Client
	Symbol      string
	Address     string
	Amount      float64
	Balance     binding.String
	BalanceUsdt binding.String
	Price       binding.String
}

func (tp *Token) Start() {
	go tp.refresh()
	//go tp.refreshView()

}
func (tp *Token) Update() {
	price := GetTokenPrice(tp.Symbol)
	p := message.NewPrinter(language.English)
	tp.Price.Set(p.Sprintf("%.6f", price))
	if config.CF.FromAddress != "" {
		if strings.ToLower(tp.Symbol) == "bnb" {
			ethBalance := utils.EthBalance(config.CF.FromAddress, tp.Client)
			tp.Balance.Set(p.Sprintf("%.6f", utils.Str2Float(ethBalance)))
			tp.BalanceUsdt.Set(p.Sprintf("%.2f", utils.Str2Float(ethBalance)*price))
			tp.Amount = utils.Str2Float(ethBalance)
		} else {
			address := config.CF.BscToken[strings.ToLower(tp.Symbol)]
			_, ethBalance := utils.Erc20Balance(config.CF.FromAddress, address, tp.Client)
			tp.Balance.Set(p.Sprintf("%.6f", utils.Str2Float(ethBalance)))
			tp.BalanceUsdt.Set(p.Sprintf("%.2f", utils.Str2Float(ethBalance)*price))
			tp.Amount = utils.Str2Float(ethBalance)
		}
	} else {
		tp.Balance.Set(p.Sprintf("%.6f", 0.0))
		tp.BalanceUsdt.Set(p.Sprintf("%.2f", 0.0))
		tp.Amount = 0.0
	}

}
func (tp *Token) refresh() {

	for {
		select {
		case <-time.After(5 * time.Second):
			tp.Update()
		}
	}
}
