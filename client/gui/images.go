package gui

import (
	"braverats/client/gui/assets"

	"fyne.io/fyne/v2/canvas"
)

var (
	ImgCardMussicianBlue  *canvas.Image
	ImgCardMussicianRed   *canvas.Image
	ImgCardPrincessBlue   *canvas.Image
	ImgCardPrincessRed    *canvas.Image
	ImgCardSpyBlue        *canvas.Image
	ImgCardSpyRed         *canvas.Image
	ImgCardAssassinBlue   *canvas.Image
	ImgCardAssassinRed    *canvas.Image
	ImgCardAmbassadorBlue *canvas.Image
	ImgCardAmbassadorRed  *canvas.Image
	ImgCardWizardBlue     *canvas.Image
	ImgCardWizardRed      *canvas.Image
	ImgCardGeneralBlue    *canvas.Image
	ImgCardGeneralRed     *canvas.Image
	ImgCardPrinceBlue     *canvas.Image
	ImgCardPrinceRed      *canvas.Image
	ImgCardSuit           *canvas.Image
)

func InitImages() {
	ImgCardMussicianBlue = canvas.NewImageFromResource(assets.ResourceMusicianbluePng)
	ImgCardMussicianRed = canvas.NewImageFromResource(assets.ResourceMusicianredPng)
	ImgCardPrincessBlue = canvas.NewImageFromResource(assets.ResourcePrincessbluePng)
	ImgCardPrincessRed = canvas.NewImageFromResource(assets.ResourcePrincessredPng)
	ImgCardSpyBlue = canvas.NewImageFromResource(assets.ResourceSpybluePng)
	ImgCardSpyRed = canvas.NewImageFromResource(assets.ResourceSpyredPng)
	ImgCardAssassinBlue = canvas.NewImageFromResource(assets.ResourceAssassinbluePng)
	ImgCardAssassinRed = canvas.NewImageFromResource(assets.ResourceAssassinredPng)
	ImgCardAmbassadorBlue = canvas.NewImageFromResource(assets.ResourceAmbassadorbluePng)
	ImgCardAmbassadorRed = canvas.NewImageFromResource(assets.ResourceAmbassadorredPng)
	ImgCardWizardBlue = canvas.NewImageFromResource(assets.ResourceWizardbluePng)
	ImgCardWizardRed = canvas.NewImageFromResource(assets.ResourceWizardredPng)
	ImgCardGeneralBlue = canvas.NewImageFromResource(assets.ResourceGeneralbluePng)
	ImgCardGeneralRed = canvas.NewImageFromResource(assets.ResourceGeneralredPng)
	ImgCardPrinceBlue = canvas.NewImageFromResource(assets.ResourcePrincebluePng)
	ImgCardPrinceRed = canvas.NewImageFromResource(assets.ResourcePrinceredPng)
	ImgCardSuit = canvas.NewImageFromResource(assets.ResourceSuitPng)
}
