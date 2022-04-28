package gui

import "C"
import (
	"braverats/brp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Match struct {
	scene *fyne.Container

	enemyHandContainer  *fyne.Container
	playerHandContainer *fyne.Container

	showCard             *ShowCard
	playerTableContainer *fyne.Container
	enemyTableContainer  *fyne.Container

	roundsOnHold uint8
}

func NewMatch(parentScene *fyne.Container) *Match {
	enemyHandContainer := container.NewHBox()
	playerHandContainer := container.NewHBox()

	playerTableContainer := container.NewHBox()
	enemyTableContainer := container.NewHBox()
	tableContainer := container.New(tableLayout{}, enemyTableContainer, playerTableContainer)

	showCard := NewShowCard()

	mainContainer := container.NewBorder(container.NewCenter(enemyHandContainer), container.NewCenter(playerHandContainer), nil, showCard, tableContainer)
	parentScene.Add(mainContainer)

	return &Match{
		scene:                parentScene,
		enemyHandContainer:   enemyHandContainer,
		playerHandContainer:  playerHandContainer,
		showCard:             showCard,
		playerTableContainer: playerTableContainer,
		enemyTableContainer:  enemyTableContainer,
	}
}

func (m *Match) ClearMatch() {
	clearContainer(m.enemyHandContainer)
	clearContainer(m.playerHandContainer)
	clearContainer(m.playerTableContainer)
	clearContainer(m.enemyTableContainer)
}

func clearContainer(container *fyne.Container) {
	for _, child := range container.Objects {
		container.Remove(child)
	}
}

func (m *Match) AddPlayerHandCards(f func(id brp.CardID) bool, ids ...brp.CardID) {
	for _, id := range ids {
		card := NewPlayerCard(id, false)
		card.OnTap = func() {
			m.PutCardOnPlayerTable(card.CardID)
			if f(card.CardID) {
				m.showCard.Hide()
				m.playerHandContainer.Remove(card)
				m.playerHandContainer.Refresh()
			} else {
				m.RemovePlayerTableCard(1)
			}
		}
		card.OnMouseIn = func() {
			m.showCard.ShowRecourse(card.image.Resource)
		}
		card.OnMouseOut = func() {
			m.showCard.Hide()
		}
		m.playerHandContainer.Add(card)
	}
}

func (m *Match) PutCardOnPlayerTable(id brp.CardID) {
	card := NewTableCard(id, false)
	card.OnMouseIn = func() {
		OnMouseInTableStandard(card)
		m.showCard.ShowRecourse(card.image.Resource)
	}
	card.OnMouseOut = func() {
		OnMouseOutTableStandard(card)
		m.showCard.Hide()
	}
	m.playerTableContainer.Add(card)
	m.playerTableContainer.Refresh()
}

func (m *Match) PutCardOnEnemyTable(id brp.CardID) {
	card := NewTableCard(id, true)
	if id == brp.CardUnknown {
		card.OnMouseIn = nil
		card.OnMouseOut = nil
	} else {
		card.OnMouseIn = func() {
			OnMouseInTableStandard(card)
			m.showCard.ShowRecourse(card.image.Resource)
		}
		card.OnMouseOut = func() {
			OnMouseOutTableStandard(card)
			m.showCard.Hide()
		}
	}
	m.enemyTableContainer.Add(card)
	m.enemyTableContainer.Refresh()
}

func (m *Match) AddEnemyHandCards(count uint8) {
	for i := uint8(0); i < count; i++ {
		m.enemyHandContainer.Add(NewCard(brp.CardUnknown, false))
	}
}

func (m *Match) RemovePlayerTableCard(count uint8) {
	removeTableCard(m.playerTableContainer, count)
}

func (m *Match) RemoveEnemyHandCard(count uint8) {
	for i := uint8(0); i < count; i++ {
		m.enemyHandContainer.Remove(m.enemyHandContainer.Objects[0])
	}
}

func (m *Match) RemoveEnemyTableCard(count uint8) {
	removeTableCard(m.enemyTableContainer, count)
}

func removeTableCard(container *fyne.Container, count uint8) {
	for i := uint8(0); i < count; i++ {
		size := len(container.Objects)
		if size == 0 {
			return
		}
		container.Remove(container.Objects[size-1])
	}
}

func (m *Match) RedrawTable(result brp.RoundResult) {
	switch result {
	case brp.HoldRound:
		m.HoldCards()
		m.roundsOnHold++
	case brp.WinRound, brp.WinGame:
		m.PlayerWonCard(m.roundsOnHold + 1)
		m.EnemyLoseCard(m.roundsOnHold + 1)
		m.roundsOnHold = 0
	case brp.LoseRound, brp.LoseGame:
		m.PlayerLoseCard(m.roundsOnHold + 1)
		m.EnemyWonCard(m.roundsOnHold + 1)
		m.roundsOnHold = 0
	}
}

func (m *Match) HoldCards() {
	setTranslucency(m.enemyTableContainer, 1, 0.2)
	setTranslucency(m.playerTableContainer, 1, 0.2)
}

func (m *Match) EnemyWonCard(count uint8) {
	setTranslucency(m.enemyTableContainer, count, 0)
}

func (m *Match) PlayerWonCard(count uint8) {
	setTranslucency(m.playerTableContainer, count, 0)
}

func (m *Match) EnemyLoseCard(count uint8) {
	setTranslucency(m.enemyTableContainer, count, 0.4)
}

func (m *Match) PlayerLoseCard(count uint8) {
	setTranslucency(m.playerTableContainer, count, 0.4)
}

func setTranslucency(tableContainer *fyne.Container, count uint8, translucency float64) {
	size := len(tableContainer.Objects)
	for i := 1; uint8(i) <= count && i <= size; i++ {
		card := tableContainer.Objects[size-i].(*TableCard)
		card.image.Translucency = translucency
	}
}
