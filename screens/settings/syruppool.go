package settingss

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/ocoderr/bsc-tool/config"
	"github.com/spf13/viper"
	"strings"
)

// SyrupScreen account info

var SyrupPoolListSelectIndex int = -1

func SyrupPoolTab(w fyne.Window) *fyne.Container {

	KeyBingData, KeyInput := NewStrEntryWithData()
	ValueBingData, ValueInput := NewStrEntryWithData()

	list := widget.NewList(
		func() int { return len(config.CF.SyrupPool) },
		func() fyne.CanvasObject {

			symbol := widget.NewLabel("List item x")
			address := widget.NewLabel("List item x")
			return container.NewHBox(symbol, address)
		},
		func(index widget.ListItemID, template fyne.CanvasObject) {
			cont := template.(*fyne.Container)
			label := cont.Objects[0].(*widget.Label)
			label.SetText(config.SyrupPoolKeyIndex(index))

			label = cont.Objects[1].(*widget.Label)
			label.SetText(config.SyrupPoolValueIndex(index))
		})

	list.OnSelected = func(index widget.ListItemID) {
		SyrupPoolListSelectIndex = index
	}
	list.OnUnselected = func(index widget.ListItemID) {
		SyrupPoolListSelectIndex = -1
	}

	AddBtn := widget.NewButton("Add", func() {
		keyBing, _ := KeyBingData.Get()
		valueBing, _ := ValueBingData.Get()
		keyBing = strings.ToLower(keyBing)
		valueBing = strings.ToLower(valueBing)
		if keyBing != "" && valueBing != "" {
			config.CF.SyrupPool[keyBing] = valueBing
			config.SaveValue(fmt.Sprintf("bsctoken.%s", keyBing), valueBing)
			list.Refresh()
			list.UnselectAll()
			BscTokenListSelectIndex = -1
			KeyBingData.Set("")
			ValueBingData.Set("")
		}

	})
	DeleteBtn := widget.NewButton("Delete", func() {
		if SyrupPoolListSelectIndex >= 0 {
			key := config.SyrupPoolKeyIndex(SyrupPoolListSelectIndex)
			if key != "" {
				key = strings.ToLower(key)
				if key == "cake-auto" || key == "cake" {
					list.Refresh()
					list.UnselectAll()
					SyrupPoolListSelectIndex = -1
					dialog.ShowError(errors.New(fmt.Sprintf("can not delete %s", key)), w)
					return
				}
				delete(config.CF.SyrupPool, key)
				delete(viper.Get("syruppool").(map[string]interface{}), key)
				viper.WriteConfig()
				list.Refresh()
				list.UnselectAll()
				SyrupPoolListSelectIndex = -1
			}
		}
	})

	content := container.NewBorder(container.NewVBox(
		widget.NewForm(
			&widget.FormItem{Text: "Pool:", Widget: KeyInput},
			&widget.FormItem{Text: "Address:", Widget: ValueInput},
		),
		container.NewGridWithColumns(2, AddBtn, DeleteBtn),
	), nil, nil, nil, list,
	)

	return content
}
