package gui

import "C"
import (
	"braverats/brp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Match struct {
	scene                *fyne.Container
	enemyHandContainer   *fyne.Container
	playerHandContainer  *fyne.Container
	showCard             *ShowCard
	playerTableContainer *fyne.Container
	enemyTableContainer  *fyne.Container
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

func (m *Match) AddPlayerHandCards(ids ...brp.CardID) {
	for _, id := range ids {
		card := NewPlayerCard(id)
		card.OnTap = func() {
			m.showCard.Hide()
			m.playerHandContainer.Remove(card)
			m.playerHandContainer.Refresh()
			m.PutCardOnPlayerTable(card.CardID)
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
	card := NewTableCard(id)
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
	card := NewTableCard(id)
	card.OnMouseIn = nil
	card.OnMouseOut = nil
	m.enemyTableContainer.Add(card)
	m.enemyTableContainer.Refresh()
}

func (m *Match) AddEnemyHandCards(count uint8) {
	for i := uint8(0); i < count; i++ {
		card := NewCard(brp.CardUnknown)

		card.OnTap = func() {
			m.enemyHandContainer.Remove(card)
			m.enemyHandContainer.Refresh()
			m.PutCardOnEnemyTable(card.CardID)
		}
		m.enemyHandContainer.Add(card)
	}
}

func (m *Match) RemoveEnemyCard(count uint8) {
	for i := uint8(0); i < count; i++ {
		m.enemyHandContainer.Remove(m.enemyHandContainer.Objects[0])
	}
}
