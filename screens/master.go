package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"github.com/ocoderr/bsc-tool/config"
)

// Master master window
func Master(a fyne.App) fyne.Window {
	w := a.NewWindow(config.GetText("app.title"))
	w.SetMaster()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(config.GetText("main.tab.account"), theme.StorageIcon(), AccountScreen(w)),
		container.NewTabItemWithIcon(config.GetText("main.tab.syrup"), theme.StorageIcon(), SyrupScreen(w)),
		container.NewTabItemWithIcon(config.GetText("main.tab.setting"), theme.StorageIcon(), SettingScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.mnemonic"), MnemonicScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.balance"), BalanceScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.transfer"), TransferToOneAddressScreen(w)),
		//container.NewTabItem(config.GetText("main.tab.transfers"), TransferToManyAddressScreen(w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	menu := fyne.NewMenu(config.GetText("menu.language"),
		fyne.NewMenuItem(config.GetText("menu.language.zh"), func() {
			config.SaveValue("language", "ZH")
			dialog.ShowConfirm(config.GetText("menu.show.title"), config.GetText("menu.show.message"), func(b bool) {
				if b {
					w.Close()
				}
			}, w)

		}), fyne.NewMenuItem(config.GetText("menu.language.en"), func() {

			config.SaveValue("language", "EN")
			dialog.ShowConfirm(config.GetText("menu.show.title"), config.GetText("menu.show.message"), func(b bool) {
				if b {
					w.Close()
				}
			}, w)

		}))

	w.SetMainMenu(fyne.NewMainMenu(menu))
	w.SetContent(tabs)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(1500, 800))
	w.ShowAndRun()

	return w
}
