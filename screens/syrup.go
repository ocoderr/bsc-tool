package screens

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/ocoderr/bsc-tool/data"
	"strings"
)

// SyrupScreen account info

func SyrupScreen(w fyne.Window) *fyne.Container {

	AddressBingData, entryAddress := NewStrEntryWithData()
	//dataPrivateKey, entryPrivateKey := NewStrEntryWithData()

	gridwrap := container.NewGridWrap(fyne.NewSize(350, 350))

	for _, symbol := range config.SyrupPoolKeys() {
		totalBingData := binding.NewString()
		RewardBingData := binding.NewString()
		RewardUsdtBingData := binding.NewString()
		StakeBingData := binding.NewString()
		ApyBingData := binding.NewString()
		PriceBingData := binding.NewString()
		EstimationBingData := binding.NewString()
		tradePairs := data.SyrupPool{
			Client:     data.Client,
			PoolName:   symbol,
			Total:      totalBingData,
			Reward:     RewardBingData,
			RewardUsdt: RewardUsdtBingData,
			Stake:      StakeBingData,
			Apy:        ApyBingData,
			Price:      PriceBingData,
			Estimation: EstimationBingData,
		}
		if !tradePairs.IsActive() {
			continue
		}
		tradePairs.Start()
		data.SyrupPools = append(data.SyrupPools, tradePairs)

		var rewardToken string
		if strings.ToLower(symbol) == "cake-auto" {
			rewardToken = "CAKE"
		} else {
			rewardToken = symbol
		}

		gridwrap.Add(widget.NewCard(strings.ToUpper(symbol), "",
			widget.NewForm(
				&widget.FormItem{Text: "Total(CAKE):", Widget: widget.NewLabelWithData(totalBingData)},
				&widget.FormItem{Text: "APY:", Widget: widget.NewLabelWithData(ApyBingData)},
				&widget.FormItem{Text: "Price:", Widget: widget.NewLabelWithData(PriceBingData)},
				&widget.FormItem{Text: "Estimated(1K Cake/Day):", Widget: widget.NewLabelWithData(EstimationBingData)},

				&widget.FormItem{Text: "Stake(CAKE):", Widget: widget.NewLabelWithData(StakeBingData)},
				&widget.FormItem{Text: fmt.Sprintf("Reward(%s):", rewardToken), Widget: widget.NewLabelWithData(RewardBingData)},
				&widget.FormItem{Text: "Reward(USDT):", Widget: widget.NewLabelWithData(RewardUsdtBingData)},
			)),
		)
	}
	if config.CF.FromAddress != "" {
		AddressBingData.Set(config.CF.FromAddress)
		go func() {
			for _, pool := range data.SyrupPools {
				pool.Update()
			}
		}()
	}

	GoBtn := widget.NewButton("Connect", func() {
		fromAddress, _ := AddressBingData.Get()
		if fromAddress == "" {
			return
		}
		config.CF.FromAddress = fromAddress
		config.SaveValue("FromAddress", fromAddress)
		go func() {
			for _, pool := range data.SyrupPools {
				pool.Update()
			}
		}()
	})

	//center := container.NewMax(textArea)
	//center.Resize(fyne.NewSize(400,1000))

	content := container.NewBorder(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{Text: config.GetText("main.field.address"), Widget: entryAddress},
		),
		GoBtn,
	), nil, nil, nil, container.NewVScroll(gridwrap),
	)

	return content
}
