package gui

import (
	"braverats/brp"
	"errors"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type LobbyCreatorDialog struct {
	dialog.Dialog
}

func NewLobbyCreatorDialog(title, confirm string, callback func(string), parent fyne.Window) *LobbyCreatorDialog {
	lobbyNameString := binding.NewString()
	lobbyNameEntry := widget.NewEntryWithData(lobbyNameString)
	lobbyNameEntry.SetPlaceHolder("Lobby name")
	lobbyNameEntry.Validator = func(s string) error {
		if len(s) == 0 || len(s) > brp.MaxLobbyNameLen {
			return errors.New("")
		}
		return nil
	}
	lobbyNameFormItem := widget.NewFormItem("", lobbyNameEntry)
	dial := &LobbyCreatorDialog{
		Dialog: dialog.NewForm(title, confirm, "Cancel", []*widget.FormItem{lobbyNameFormItem},
			func(confirm bool) {
				if !confirm {
					return
				}
				name, err := lobbyNameString.Get()
				if err != nil {
					log.Println(err)
					return
				}
				callback(name)
			}, parent),
	}
	dial.Resize(fyne.NewSize(300, 100))
	return dial
}
