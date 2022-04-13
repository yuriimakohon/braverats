package gui

import (
	"braverats/brp"
	"errors"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewLobbyCreatorDialog(title, confirm string, callback func(string), parent fyne.Window) dialog.Dialog {
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
	dial := dialog.NewForm(title, confirm, "Cancel", []*widget.FormItem{lobbyNameFormItem},
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
		}, parent)
	dial.Resize(fyne.NewSize(300, 100))
	return dial
}

type Player struct {
	Name  binding.String
	Ready binding.Bool
}

func NewPlayer(name string) *Player {
	player := &Player{
		Name:  binding.NewString(),
		Ready: binding.NewBool(),
	}
	player.Name.Set(name)
	return player
}

type Lobby struct {
	Name         binding.String
	FirstPlayer  *Player
	SecondPlayer *Player
}

func NewLobby() *Lobby {
	lobby := &Lobby{
		Name:         binding.NewString(),
		FirstPlayer:  NewPlayer(""),
		SecondPlayer: NewPlayer(""),
	}
	return lobby
}

// Reset players and set lobby name
func (l *Lobby) Reset(name string) error {
	err := l.Name.Set(name)
	if err != nil {
		return err
	}
	err = l.FirstPlayer.Name.Set("You")
	if err != nil {
		return err
	}
	err = l.FirstPlayer.Ready.Set(false)
	if err != nil {
		return err
	}
	err = l.ResetSecondPlayer()
	return err
}

func (l *Lobby) ResetSecondPlayer() error {
	err := l.SecondPlayer.Name.Set("Waiting for opponent")
	if err != nil {
		return err
	}
	err = l.SecondPlayer.Ready.Set(false)
	return err
}

func NewLobbyDialog(onReady func(bool), onClosed func(), lobby Lobby, parent fyne.Window) dialog.Dialog {
	lobbyNameLabel := widget.NewLabelWithData(lobby.Name)

	changeReadyLabel := func(check *widget.Check, ready bool) {
		if ready {
			check.Text = "Ready"
		} else {
			check.Text = "Not ready"
		}
	}

	playerNameLabel := widget.NewLabelWithData(lobby.FirstPlayer.Name)
	playerReadinessCheck := widget.NewCheckWithData("Not ready", lobby.FirstPlayer.Ready)
	playerReadinessCheck.OnChanged = func(ready bool) {
		changeReadyLabel(playerReadinessCheck, ready)
		onReady(ready)
	}

	anotherPlayerNameLabel := widget.NewLabelWithData(lobby.SecondPlayer.Name)
	anotherPlayerReadinessCheck := widget.NewCheckWithData("Not ready", lobby.SecondPlayer.Ready)
	anotherPlayerReadinessCheck.Disable()
	anotherPlayerReadinessCheck.OnChanged = func(ready bool) {
		changeReadyLabel(anotherPlayerReadinessCheck, ready)
	}

	readinessCnt := container.NewGridWithColumns(2, playerNameLabel, playerReadinessCheck, anotherPlayerNameLabel, anotherPlayerReadinessCheck)
	main := container.NewBorder(container.NewCenter(lobbyNameLabel), readinessCnt, nil, nil)

	dial := dialog.NewCustom("Lobby", "Leave", main, parent)
	dial.Resize(fyne.NewSize(350, 100))
	dial.SetOnClosed(func() { onClosed() })

	return dial
}
