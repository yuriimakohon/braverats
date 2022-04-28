package gui

import (
	"braverats/brp"

	"fyne.io/fyne/v2"
)

type Hand struct {
	container      *fyne.Container
	disabled       bool
	onTapFunctions map[brp.CardID]func()
}

func NewHand(container *fyne.Container) *Hand {
	return &Hand{
		container:      container,
		onTapFunctions: make(map[brp.CardID]func()),
	}
}

func (h *Hand) AddCard(card fyne.CanvasObject) {
	h.container.Add(card)
}

func (h *Hand) RemoveCard(card fyne.CanvasObject) {
	h.container.Remove(card)
	h.container.Refresh()
}

func (h *Hand) PopCards(count uint8) {
	for i := uint8(0); i < count; i++ {
		size := len(h.container.Objects)
		if size == 0 {
			return
		}
		h.container.Remove(h.container.Objects[size-1])
	}
}

func (h *Hand) Disable() {
	if h.disabled {
		return
	}

	for _, object := range h.container.Objects {
		card, ok := object.(*PlayerCard)
		if ok {
			card.image.Translucency = 0.3
			h.onTapFunctions[card.CardID] = card.OnTap
			card.OnTap = nil
		}
	}
	h.disabled = true
}

func (h *Hand) Enable() {
	if !h.disabled {
		return
	}

	for _, object := range h.container.Objects {
		card, ok := object.(*PlayerCard)
		if ok {
			card.image.Translucency = 0
			card.OnTap = h.onTapFunctions[card.CardID]
		}
	}
	h.disabled = false
}

func (h *Hand) Disabled() bool {
	return h.disabled
}
