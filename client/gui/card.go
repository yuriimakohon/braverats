package gui

import (
	"braverats/brp"
	"braverats/client/gui/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type cardRenderer struct {
	c     *Card
	image *canvas.Image
}

func (cr cardRenderer) Destroy() {
}

func (cr cardRenderer) Layout(size fyne.Size) {
	cr.image.Resize(size)
}

func (cr cardRenderer) MinSize() fyne.Size {
	return cr.image.MinSize()
}

func (cr cardRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{cr.image}
}

func (cr cardRenderer) Refresh() {
}

func (c *Card) CreateRenderer() fyne.WidgetRenderer {
	return &cardRenderer{c: c, image: &c.image}
}

type Card struct {
	widget.BaseWidget
	brp.CardID

	image canvas.Image

	OnTap func()
}

func (c *Card) Resize(size fyne.Size) {
	c.BaseWidget.Resize(size)
	c.image.Resize(size)
}

func (c *Card) Tapped(event *fyne.PointEvent) {
	if c.OnTap != nil {
		c.OnTap()
	}
}

func (c *Card) SetMinSize(size fyne.Size) {
	c.image.SetMinSize(size)
}

func NewCard(id brp.CardID) *Card {
	var resource fyne.Resource
	switch id {
	case brp.CardUnknown:
		resource = assets.ResourceSuitPng
	case brp.CardMusician:
		resource = assets.ResourceMusicianbluePng
	case brp.CardPrincess:
		resource = assets.ResourcePrincessbluePng
	case brp.CardSpy:
		resource = assets.ResourceSpybluePng
	case brp.CardAssassin:
		resource = assets.ResourceAssassinbluePng
	case brp.CardAmbassador:
		resource = assets.ResourceAmbassadorbluePng
	case brp.CardWizard:
		resource = assets.ResourceWizardbluePng
	case brp.CardGeneral:
		resource = assets.ResourceGeneralbluePng
	case brp.CardPrince:
		resource = assets.ResourcePrincebluePng
	}
	image := canvas.NewImageFromResource(resource)
	image.ScaleMode = canvas.ImageScalePixels
	image.FillMode = canvas.ImageFillContain

	card := &Card{
		CardID: id,
		image:  *image,
	}
	card.image.SetMinSize(fyne.NewSize(90, 120))

	card.ExtendBaseWidget(card)

	return card
}

type TableCard struct {
	*Card
	OnMouseIn  func()
	OnMouseOut func()
}

func (t TableCard) MouseIn(event *desktop.MouseEvent) {
	if t.OnMouseIn != nil {
		t.OnMouseIn()
	}
}

func (t TableCard) MouseOut() {
	if t.OnMouseOut != nil {
		t.OnMouseOut()
	}
}

func (t TableCard) MouseMoved(event *desktop.MouseEvent) {
}

func OnMouseInTableStandard(card *TableCard) {
	cardSize := card.Size()
	card.image.Resize(fyne.NewSize(cardSize.Width*1.1, cardSize.Height*1.1))
	imageSize := card.image.Size()
	card.image.Move(fyne.NewPos(cardSize.Width/2-imageSize.Width/2, cardSize.Height/2-imageSize.Height/2))
}

func OnMouseOutTableStandard(card *TableCard) {
	card.image.Resize(card.Size())
	card.image.Move(fyne.NewPos(0, 0))
}

func NewTableCard(id brp.CardID) *TableCard {
	card := &TableCard{
		Card: NewCard(id),
	}
	card.OnMouseIn = func() {
		OnMouseInTableStandard(card)
	}
	card.OnMouseOut = func() {
		OnMouseOutTableStandard(card)
	}

	return card
}

type PlayerCard struct {
	*Card

	OnMouseIn  func()
	OnMouseOut func()

	hoverSize  fyne.Size
	hoverPos   fyne.Position
	defaultPos fyne.Position
}

func (c *PlayerCard) MouseIn(event *desktop.MouseEvent) {
	c.OnMouseIn()
	c.image.Move(c.hoverPos)
	c.image.Refresh()
}

func (c *PlayerCard) MouseOut() {
	c.OnMouseOut()
	c.image.Move(c.defaultPos)
	c.image.Refresh()
}

func (c *PlayerCard) MouseMoved(event *desktop.MouseEvent) {
}

func NewPlayerCard(id brp.CardID) *PlayerCard {
	card := &PlayerCard{
		Card:       NewCard(id),
		defaultPos: fyne.NewPos(0, 0),
	}
	size := fyne.NewSize(100, 150)
	card.SetMinSize(size)
	card.hoverPos = fyne.NewPos(0, -size.Height/6)

	card.ExtendBaseWidget(card)

	return card
}

type ShowCard struct {
	*Card
}

func NewShowCard() *ShowCard {
	card := &ShowCard{
		Card: NewCard(brp.CardUnknown),
	}
	card.image.Hide()
	card.SetMinSize(fyne.NewSize(200, 300))

	return card
}

func (c *ShowCard) Show() {
	c.image.Show()
}

func (c *ShowCard) ShowRecourse(resource fyne.Resource) {
	c.image.Resource = resource
	c.image.Refresh()
	c.Show()
}

func (c *ShowCard) Hide() {
	c.image.Hide()
}
