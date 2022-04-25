package client

import (
	"braverats/brp"
	"braverats/client/gui"

	"fyne.io/fyne/v2/container"
)

type match struct {
	gui *gui.Match
}

func (app *App) initMatchScene() {
	matchContainer := container.NewMax()
	guiMatch := gui.NewMatch(matchContainer)

	app.match = &match{
		gui: guiMatch,
	}

	app.gui.AddScene(gui.GIDMatch, matchContainer)
}

func (app *App) RenderMatch() {
	app.match.gui.AddEnemyHandCards(8)
	app.match.gui.AddPlayerHandCards(brp.CardMusician, brp.CardPrincess, brp.CardSpy, brp.CardAssassin, brp.CardAmbassador, brp.CardWizard, brp.CardGeneral, brp.CardPrince)
}
