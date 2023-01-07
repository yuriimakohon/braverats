package theme

import (
	"braverats/client/gui/assets"
	"fyne.io/fyne/v2/canvas"
)

var defaultTheme Theme

func initDefaultTheme() {
	defaultTheme = Theme{
		ImgCardMusicianBlue:   canvas.NewImageFromResource(assets.DefaultMusicianbluePng),
		ImgCardMusicianRed:    canvas.NewImageFromResource(assets.DefaultMusicianredPng),
		ImgCardPrincessBlue:   canvas.NewImageFromResource(assets.DefaultPrincessbluePng),
		ImgCardPrincessRed:    canvas.NewImageFromResource(assets.DefaultPrincessredPng),
		ImgCardSpyBlue:        canvas.NewImageFromResource(assets.DefaultSpybluePng),
		ImgCardSpyRed:         canvas.NewImageFromResource(assets.DefaultSpyredPng),
		ImgCardAssassinBlue:   canvas.NewImageFromResource(assets.DefaultAssassinbluePng),
		ImgCardAssassinRed:    canvas.NewImageFromResource(assets.DefaultAssassinredPng),
		ImgCardAmbassadorBlue: canvas.NewImageFromResource(assets.DefaultAmbassadorbluePng),
		ImgCardAmbassadorRed:  canvas.NewImageFromResource(assets.DefaultAmbassadorredPng),
		ImgCardWizardBlue:     canvas.NewImageFromResource(assets.DefaultWizardbluePng),
		ImgCardWizardRed:      canvas.NewImageFromResource(assets.DefaultWizardredPng),
		ImgCardGeneralBlue:    canvas.NewImageFromResource(assets.DefaultGeneralbluePng),
		ImgCardGeneralRed:     canvas.NewImageFromResource(assets.DefaultGeneralredPng),
		ImgCardPrinceBlue:     canvas.NewImageFromResource(assets.DefaultPrincebluePng),
		ImgCardPrinceRed:      canvas.NewImageFromResource(assets.DefaultPrinceredPng),
		ImgCardSuit:           canvas.NewImageFromResource(assets.DefaultSuitPng),
		ImgBackground:         canvas.NewImageFromResource(assets.DefaultBackgroundPng),
	}
}
