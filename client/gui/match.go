package gui

import "C"
import (
	"braverats/brp"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Match struct {
	scene *fyne.Container

	PlayerHand *Hand
	EnemyHand  *Hand

	showCard             *ShowCard
	playerTableContainer *fyne.Container
	enemyTableContainer  *fyne.Container

	MatchResult matchResult

	roundsOnHold uint8
}

type matchResult struct {
	sub binding.DataListener
	res brp.RoundResult
}

func (m *matchResult) Set(result brp.RoundResult) {
	switch result {
	case brp.DrawGame, brp.LoosedGame, brp.WonGame:
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

func NewMatch(parentScene *fyne.Container) *Match {
	enemyHandContainer := container.NewHBox()
	playerHandContainer := container.NewHBox()

	playerHand := NewHand(playerHandContainer)
	enemyHand := NewHand(enemyHandContainer)

	playerTableContainer := container.NewHBox()
	enemyTableContainer := container.NewHBox()
	tableContainer := container.New(tableLayout{}, enemyTableContainer, playerTableContainer)

	showCard := NewShowCard()

	mainContainer := container.NewBorder(container.NewCenter(enemyHandContainer), container.NewCenter(playerHandContainer), nil, showCard, tableContainer)
	parentScene.Add(mainContainer)

	return &Match{
		scene:                parentScene,
		PlayerHand:           playerHand,
		EnemyHand:            enemyHand,
		showCard:             showCard,
		playerTableContainer: playerTableContainer,
		enemyTableContainer:  enemyTableContainer,
	}
}

func (m *Match) ClearMatch() {
	clearContainer(m.PlayerHand.container)
	clearContainer(m.EnemyHand.container)
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
			m.PlayerHand.Disable()
			if f(card.CardID) {
				m.showCard.Hide()
				m.PlayerHand.RemoveCard(card)
			} else {
				m.RemovePlayerTableCard(1)
				m.PlayerHand.Enable()
			}
		}
		card.OnMouseIn = func() {
			m.showCard.ShowRecourse(card.image.Resource)
		}
		card.OnMouseOut = func() {
			m.showCard.Hide()
		}
		m.PlayerHand.AddCard(card)
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
		m.EnemyHand.AddCard(NewCard(brp.CardUnknown, false))
	}
}

func (m *Match) RemovePlayerTableCard(count uint8) {
	removeTableCard(m.playerTableContainer, count)
}

func (m *Match) RemoveEnemyHandCard(count uint8) {
	m.EnemyHand.PopCards(count)
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
	lastPlayerCard := m.playerTableContainer.Objects[len(m.playerTableContainer.Objects)-1].(*TableCard)
	lastEnemyCard := m.enemyTableContainer.Objects[len(m.enemyTableContainer.Objects)-1].(*TableCard)

	if lastPlayerCard.Card.CardID == brp.CardSpy && lastEnemyCard.CardID != brp.CardSpy {
		m.PlayerHand.Disable()
	} else {
		m.PlayerHand.Enable()
	}
	switch result {
	case brp.HeldRound:
		m.HoldCards()
		m.roundsOnHold++
	case brp.WonRound, brp.WonGame:
		m.PlayerWonCard(m.roundsOnHold + 1)
		m.EnemyLoseCard(m.roundsOnHold + 1)
		m.roundsOnHold = 0
	case brp.LoosedRound, brp.LoosedGame:
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
	setTranslucency(m.enemyTableContainer, count, 0.6)
}

func (m *Match) PlayerLoseCard(count uint8) {
	setTranslucency(m.playerTableContainer, count, 0.6)
}

func setTranslucency(tableContainer *fyne.Container, count uint8, translucency float64) {
	size := len(tableContainer.Objects)
	for i := 1; uint8(i) <= count && i <= size; i++ {
		card := tableContainer.Objects[size-i].(*TableCard)
		card.image.Translucency = translucency
	}
}

func NewMatchEndDialog(matchResult *matchResult, onClosed func(), parent fyne.Window) dialog.Dialog {
	message := binding.NewString()
	matchResult.AddListener(binding.NewDataListener(func() {
		switch matchResult.res {
		case brp.WonGame:
			message.Set("You won the match")
		case brp.LoosedGame:
			message.Set("You loosed the match")
		case brp.DrawGame:
			message.Set("Match is draw")
		}
	}))
	label := widget.NewLabelWithData(message)
	dial := dialog.NewCustom("Match", "Leave lobby", label, parent)

	dial.SetOnClosed(func() { onClosed() })

	return dial
}
