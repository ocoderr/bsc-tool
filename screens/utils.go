package screens

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewIntEntryWithData() (binding.Int, *widget.Entry) {
	dataaddressNum := binding.NewInt()
	dataaddressNumStr := binding.IntToString(dataaddressNum)
	addressNumEntry := widget.NewEntryWithData(dataaddressNumStr)
	return dataaddressNum, addressNumEntry
}
func NewFloatEntryWithData() (binding.Float, *widget.Entry) {
	dataaddressNum := binding.NewFloat()
	dataaddressNumStr := binding.FloatToString(dataaddressNum)
	addressNumEntry := widget.NewEntryWithData(dataaddressNumStr)
	return dataaddressNum, addressNumEntry
}

func NewStrEntryWithData() (binding.String, *widget.Entry) {
	dataAddress := binding.NewString()
	addEntry := widget.NewEntryWithData(dataAddress)
	//widget.NewEntry()
	//addEntry.Disable()
	return dataAddress, addEntry
}
