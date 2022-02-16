package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/ocoderr/bsc-tool/data"
	"github.com/ocoderr/bsc-tool/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
	"time"
)

// SyrupScreen account info

func AccountScreen(w fyne.Window) *fyne.Container {

	AddressBingData, entryAddress := NewStrEntryWithData()
	//_, entryPrivateKey := NewStrEntryWithData()

	TotalBingData := binding.NewString()
	//PrivateKeyBingData := binding.NewString()
	//entryPrivateKey := widget.NewPasswordEntry()
	//entryPrivateKey.Bind(PrivateKeyBingData)

	gridwrap := container.NewGridWrap(fyne.NewSize(250, 200))

	//addCard("bnb", gridwrap)
	//addCard("wbnb", gridwrap)
	//addCard("usdt", gridwrap)
	//addCard("busd", gridwrap)

	//for symbol, _ := range config.CF.BscToken {
	//	if symbol == "bnb" || symbol == "wbnb" || symbol == "usdt" || symbol == "busd" {
	//		continue
	//	}
	//	addCard(symbol, gridwrap)
	//}
	if config.CF.FromAddress != "" {
		AddressBingData.Set(config.CF.FromAddress)
		go func() {
			addCard("bnb", gridwrap)
			addCard("wbnb", gridwrap)
			addCard("usdt", gridwrap)
			addCard("busd", gridwrap)

			for symbol, _ := range config.CF.BscToken {
				if symbol == "bnb" || symbol == "wbnb" || symbol == "usdt" || symbol == "busd" {
					continue
				}
				addCard(symbol, gridwrap)
			}
		}()

	}

	go RefreshTotal(TotalBingData)

	ConnectBtn := widget.NewButton("Connect", func() {
		//_, _ := PrivateKeyBingData.Get()
		//config.CF.FromAddress = fromAddress
		//config.SaveValue("FromAddress", fromAddress)
		//go func() {
		//	for _, pool := range data.SyrupPools {
		//		pool.Update()
		//	}
		//}()
	})
	QueryBtn := widget.NewButton("Query", func() {
		fromAddress, _ := AddressBingData.Get()
		if fromAddress == "" {
			return
		}
		config.CF.FromAddress = fromAddress
		config.SaveValue("FromAddress", fromAddress)
		go func() {
			addCard("bnb", gridwrap)
			addCard("wbnb", gridwrap)
			addCard("usdt", gridwrap)
			addCard("busd", gridwrap)

			for symbol, _ := range config.CF.BscToken {
				if symbol == "bnb" || symbol == "wbnb" || symbol == "usdt" || symbol == "busd" {
					continue
				}
				addCard(symbol, gridwrap)
			}
		}()

	})

	content := container.NewBorder(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{Text: config.GetText("main.field.address"), Widget: entryAddress},
			//&widget.FormItem{Text: "Private Key:", Widget: entryPrivateKey},
			&widget.FormItem{Text: "Total:", Widget: widget.NewLabelWithData(TotalBingData)},
		),
		container.NewGridWithColumns(2, ConnectBtn, QueryBtn),
	), nil, nil, nil, container.NewVScroll(gridwrap),
	)

	return content
}
func RefreshTotal(totalBinding binding.String) {
	for {
		select {
		case <-time.After(5 * time.Second):
			var total float64
			for _, pool := range data.Tokens {
				values, _ := pool.BalanceUsdt.Get()
				total += utils.Str2Float(values)
			}
			p := message.NewPrinter(language.English)
			totalBinding.Set(p.Sprintf("%.6f", total))
		}
	}
}

func addCard(symbol string, gridwrap *fyne.Container) {
	BalanceBingData := binding.NewString()
	BalanceUsdtBingData := binding.NewString()
	PriceBingData := binding.NewString()

	tradePairs := data.Token{
		Client:      data.Client,
		Symbol:      symbol,
		Balance:     BalanceBingData,
		BalanceUsdt: BalanceUsdtBingData,
		Price:       PriceBingData,
	}
	tradePairs.Update()
	if tradePairs.Amount > 0.0 {
		gridwrap.Add(widget.NewCard(strings.ToUpper(symbol), "",
			widget.NewForm(
				&widget.FormItem{Text: "Balance:", Widget: widget.NewLabelWithData(BalanceBingData)},
				&widget.FormItem{Text: "Price:", Widget: widget.NewLabelWithData(PriceBingData)},
				&widget.FormItem{Text: "Balance(USDT):", Widget: widget.NewLabelWithData(BalanceUsdtBingData)},
			)),
		)
		data.Tokens = append(data.Tokens, tradePairs)
		tradePairs.Start()
	}
}
