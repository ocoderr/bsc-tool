package settingss

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/spf13/viper"
	"strings"
)

// SyrupScreen account info

var BscTokenListSelectIndex int = -1

func BscTokenTab(w fyne.Window) *fyne.Container {

	KeyBingData, KeyInput := NewStrEntryWithData()
	ValueBingData, ValueInput := NewStrEntryWithData()

	list := widget.NewList(
		func() int { return len(config.CF.BscToken) },
		func() fyne.CanvasObject {

			symbol := widget.NewLabel("List item x")
			address := widget.NewLabel("List item x")
			return container.NewHBox(symbol, address)
		},
		func(index widget.ListItemID, template fyne.CanvasObject) {
			cont := template.(*fyne.Container)
			label := cont.Objects[0].(*widget.Label)
			label.SetText(config.BscTokenKeyIndex(index))

			label = cont.Objects[1].(*widget.Label)
			label.SetText(config.BscTokenValueIndex(index))
		})

	list.OnSelected = func(index widget.ListItemID) {
		BscTokenListSelectIndex = index
	}
	list.OnUnselected = func(index widget.ListItemID) {
		BscTokenListSelectIndex = -1
	}

	AddBtn := widget.NewButton("Add", func() {
		keyBing, _ := KeyBingData.Get()
		valueBing, _ := ValueBingData.Get()
		keyBing = strings.ToLower(keyBing)
		valueBing = strings.ToLower(valueBing)
		if keyBing != "" && valueBing != "" {
			config.CF.BscToken[keyBing] = valueBing
			config.SaveValue(fmt.Sprintf("bsctoken.%s", keyBing), valueBing)
			list.Refresh()
			list.UnselectAll()
			BscTokenListSelectIndex = -1
			KeyBingData.Set("")
			ValueBingData.Set("")
		}

	})
	DeleteBtn := widget.NewButton("Delete", func() {
		if BscTokenListSelectIndex >= 0 {
			key := config.BscTokenKeyIndex(BscTokenListSelectIndex)
			if key != "" {
				key = strings.ToLower(key)
				if key == "busd" || key == "wbnb" || key == "usdt" || key == "cake" {
					list.Refresh()
					list.UnselectAll()
					BscTokenListSelectIndex = -1
					dialog.ShowError(errors.New(fmt.Sprintf("can not delete %s", key)), w)
					return
				}
				delete(config.CF.BscToken, key)
				delete(viper.Get("bsctoken").(map[string]interface{}), key)
				viper.WriteConfig()
				list.Refresh()
				list.UnselectAll()
				BscTokenListSelectIndex = -1
			}
		}
	})

	content := container.NewBorder(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{Text: "Symbol:", Widget: KeyInput},
			&widget.FormItem{Text: "Address:", Widget: ValueInput},
		),
		container.NewGridWithColumns(2, AddBtn, DeleteBtn),
	), nil, nil, nil, list,
	)

	return content
}
func NewStrEntryWithData() (binding.String, *widget.Entry) {
	dataAddress := binding.NewString()
	addEntry := widget.NewEntryWithData(dataAddress)
	//widget.NewEntry()
	//addEntry.Disable()
	return dataAddress, addEntry
}
