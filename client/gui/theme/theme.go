package theme

import (
	"fyne.io/fyne/v2/canvas"
)

type Theme struct {
	ImgCardMusicianBlue   *canvas.Image
	ImgCardMusicianRed    *canvas.Image
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
	ImgBackground         *canvas.Image
}

var currentTheme *Theme

func Current() *Theme {
	return currentTheme
}

func Init() {
	initDefaultTheme()
	initInscryptionTheme()
	currentTheme = &defaultTheme
}
