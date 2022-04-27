package client

import (
	"braverats/brp"
	"braverats/client/gui"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type match struct {
	gui      *gui.Match
	playerIn bool
	matchResult
}

func (app *App) initMatchScene() {
	matchContainer := container.NewMax()
	guiMatch := gui.NewMatch(matchContainer)

	app.match = &match{
		gui: guiMatch,
	}

	app.gui.AddScene(gui.GIDMatch, matchContainer)
}

type matchResult struct {
	sub binding.DataListener
	res brp.RoundResult
}

func (m *matchResult) Set(result brp.RoundResult) {
	switch result {
	case brp.DrawGame, brp.LoseGame, brp.WinGame:
		m.res = result
		m.sub.DataChanged()
	}
}

func (m *matchResult) AddListener(listener binding.DataListener) {
	m.sub = listener
}

func (m matchResult) RemoveListener(listener binding.DataListener) {
	m.sub = nil
}

func (app *App) initMatchDialog() {
	message := binding.NewString()
	app.match.matchResult.AddListener(binding.NewDataListener(func() {
		switch app.match.matchResult.res {
		case brp.WinGame:
			message.Set("You won the match")
		case brp.LoseGame:
			message.Set("You loosed the match")
		case brp.DrawGame:
			message.Set("Match is draw")
		}
	}))
	matchDialog := dialog.NewCustom("Match", "close", widget.NewLabelWithData(message), app.gui.W)
	app.gui.AddDialog(gui.GIDDialMatchEnd, matchDialog)
}

func (app *App) closeMatch() {
	app.gui.ShowScene(gui.GIDMainMenu)
	app.match.playerIn = false
}

func (app *App) RenderNewMatch() {
	app.match.gui.ClearMatch()
	app.match.gui.AddEnemyHandCards(8)
	app.match.gui.AddPlayerHandCards(func(id brp.CardID) bool {
		return app.PutCard(id)
	},
		brp.CardMusician, brp.CardPrincess, brp.CardSpy, brp.CardAssassin, brp.CardAmbassador, brp.CardWizard, brp.CardGeneral, brp.CardPrince)
}

func (app *App) PutCard(card brp.CardID) bool {
	_, err := app.conn.Write(brp.NewReqPutCard(card))
	app.gui.SendErrDialog(brp.ReqPutCard, err)
	ok, _ := app.receiveAndProcessResponse(brp.ReqPutCard, "")
	return ok
}

func (app *App) CardPut(packet brp.Packet) {
	faceUp, enemyCardID, err := brp.ParseEventCardPut(packet)
	if err != nil {
		app.gui.ApplicationErrDialog(err)
		return
	}

	if !faceUp {
		app.match.gui.PutCardOnEnemyTable(brp.CardUnknown)
	} else {
		app.match.gui.PutCardOnEnemyTable(enemyCardID)
	}
}

func (app *App) RoundEnded(packet brp.Packet) {
	roundResult, enemyCardID, err := brp.ParseEventRoundEnded(packet)
	if err != nil {
		app.gui.ApplicationErrDialog(err)
		return
	}

	switch roundResult {
	case brp.HoldRound:
		app.gui.SendNotification("Round is held", "")
		app.match.gui.RemoveEnemyTableCard(1)
		app.match.gui.PutCardOnEnemyTable(enemyCardID)
	case brp.LoseRound:
		app.gui.SendNotification("You lose this round", "")
		app.match.gui.RemoveEnemyTableCard(1)
		app.match.gui.PutCardOnEnemyTable(enemyCardID)
	case brp.WinRound:
		app.gui.SendNotification("You win this round", "")
		app.match.gui.RemoveEnemyTableCard(1)
		app.match.gui.PutCardOnEnemyTable(enemyCardID)
	default:
		app.match.matchResult.Set(roundResult)
		app.gui.ShowDialog(gui.GIDDialMatchEnd)
		app.match.gui.RemoveEnemyTableCard(1)
		app.match.gui.PutCardOnEnemyTable(enemyCardID)
	}
}
