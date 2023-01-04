package client

import (
	"braverats/brp"
	"braverats/client/gui"
	"braverats/client/gui/assets"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type match struct {
	gui      *gui.Match
	playerIn bool
}

func (app *App) initMatchScene() {
	background := canvas.NewImageFromResource(assets.ResourceCastlePng)
	background.Translucency = 0.6
	matchContainer := container.NewMax(background)
	guiMatch := gui.NewMatch(matchContainer)

	app.match = &match{
		gui: guiMatch,
	}

	app.gui.AddScene(gui.GIDMatch, matchContainer)
}

func (app *App) initMatchEndDialog() {
	matchEndDialog := gui.NewMatchEndDialog(&app.match.gui.MatchResult,
		func() {
			app.LeaveLobby()
			app.gui.ShowScene(gui.GIDMainMenu)
		}, app.gui.W)
	app.gui.AddDialog(gui.GIDDialMatchEnd, matchEndDialog)
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

	app.match.gui.PlayerHand.Enable()

	app.match.gui.EnemyHand.PopCards(1)
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

	app.match.gui.RemoveEnemyTableCard(1)
	app.match.gui.PutCardOnEnemyTable(enemyCardID)
	app.match.gui.RedrawTable(roundResult)
	switch roundResult {
	case brp.HeldRound:
		app.gui.SendNotification("Hold", "Round is held")
	case brp.LoosedRound:
		app.gui.SendNotification("Lose", "You loosed this round")
	case brp.WonRound:
		app.gui.SendNotification("Win", "You won this round")
	default:
		app.match.gui.MatchResult.Set(roundResult)
		app.gui.ShowDialog(gui.GIDDialMatchEnd)
	}
}
