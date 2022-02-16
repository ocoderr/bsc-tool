package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	settingss "github.com/ocoderr/bsc-tool/screens/settings"
)

// SyrupScreen account info

func SettingScreen(w fyne.Window) *fyne.Container {

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("BSC Token", theme.StorageIcon(), settingss.BscTokenTab(w)),
		container.NewTabItemWithIcon("Syrup Pool", theme.StorageIcon(), settingss.SyrupPoolTab(w)),
		//container.NewTabItem(config.GetText("main.tab.mnemonic"), MnemonicScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.balance"), BalanceScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.transfer"), TransferToOneAddressScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.transfers"), TransferToManyAddressScreen(w)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	content := container.NewBorder(nil, nil, nil, nil, container.NewVScroll(tabs))

	return content
}
