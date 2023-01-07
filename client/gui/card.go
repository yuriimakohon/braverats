package gui

import (
	"braverats/brp"
	"braverats/client/gui/theme"

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

func NewCard(id brp.CardID, team bool) *Card {
	var image *canvas.Image
	switch id {
	default:
		image = theme.Current().ImgCardSuit
	case brp.CardMusician:
		if team {
			image = theme.Current().ImgCardMusicianRed
			break
		}
		image = theme.Current().ImgCardMusicianBlue
	case brp.CardPrincess:
		if team {
			image = theme.Current().ImgCardPrincessRed
			break
		}
		image = theme.Current().ImgCardPrincessBlue
	case brp.CardSpy:
		if team {
			image = theme.Current().ImgCardSpyRed
			break
		}
		image = theme.Current().ImgCardSpyBlue
	case brp.CardAssassin:
		if team {
			image = theme.Current().ImgCardAssassinRed
			break
		}
		image = theme.Current().ImgCardAssassinBlue
	case brp.CardAmbassador:
		if team {
			image = theme.Current().ImgCardAmbassadorRed
			break
		}
		image = theme.Current().ImgCardAmbassadorBlue
	case brp.CardWizard:
		if team {
			image = theme.Current().ImgCardWizardRed
			break
		}
		image = theme.Current().ImgCardWizardBlue
	case brp.CardGeneral:
		if team {
			image = theme.Current().ImgCardGeneralRed
			break
		}
		image = theme.Current().ImgCardGeneralBlue
	case brp.CardPrince:
		if team {
			image = theme.Current().ImgCardPrinceRed
			break
		}
		image = theme.Current().ImgCardPrinceBlue
	}
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
	ImageStdSize fyne.Size
	OnMouseIn    func()
	OnMouseOut   func()
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
	card.ImageStdSize = card.image.Size()
	cardSize := card.Size()
	newImageSize := card.ImageStdSize
	newImageSize.Height *= float32(1.1)
	newImageSize.Width *= float32(1.1)
	card.image.Resize(newImageSize)
	card.image.Move(fyne.NewPos(cardSize.Width/2-newImageSize.Width/2, cardSize.Height/2-newImageSize.Height/2))
}

func OnMouseOutTableStandard(card *TableCard) {
	card.image.Resize(card.ImageStdSize)
	cardSize := card.Size()
	card.image.Move(fyne.NewPos(cardSize.Width/2-card.ImageStdSize.Width/2, cardSize.Height/2-card.ImageStdSize.Height/2))
}

func NewTableCard(id brp.CardID, team bool) *TableCard {
	card := &TableCard{
		Card: NewCard(id, team),
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

func NewPlayerCard(id brp.CardID, team bool) *PlayerCard {
	card := &PlayerCard{
		Card:       NewCard(id, team),
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
		Card: NewCard(brp.CardUnknown, false),
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
