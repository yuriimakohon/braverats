package theme

import (
	"braverats/client/gui/assets"
	"fyne.io/fyne/v2/canvas"
)

var inscryptionTheme Theme

func initInscryptionTheme() {
	inscryptionTheme = Theme{
		ImgCardMusicianBlue:   canvas.NewImageFromResource(assets.InscryptionMusicianbluePng),
		ImgCardMusicianRed:    canvas.NewImageFromResource(assets.InscryptionMusicianredPng),
		ImgCardPrincessBlue:   canvas.NewImageFromResource(assets.InscryptionPrincessbluePng),
		ImgCardPrincessRed:    canvas.NewImageFromResource(assets.InscryptionPrincessbluePng),
		ImgCardSpyBlue:        canvas.NewImageFromResource(assets.InscryptionSpybluePng),
		ImgCardSpyRed:         canvas.NewImageFromResource(assets.InscryptionSpyredPng),
		ImgCardAssassinBlue:   canvas.NewImageFromResource(assets.InscryptionAssassinbluePng),
		ImgCardAssassinRed:    canvas.NewImageFromResource(assets.InscryptionAssassinredPng),
		ImgCardAmbassadorBlue: canvas.NewImageFromResource(assets.InscryptionAmbassadorbluePng),
		ImgCardAmbassadorRed:  canvas.NewImageFromResource(assets.InscryptionAmbassadorredPng),
		ImgCardWizardBlue:     canvas.NewImageFromResource(assets.InscryptionWizardbluePng),
		ImgCardWizardRed:      canvas.NewImageFromResource(assets.InscryptionWizardredPng),
		ImgCardGeneralBlue:    canvas.NewImageFromResource(assets.InscryptionGeneralbluePng),
		ImgCardGeneralRed:     canvas.NewImageFromResource(assets.InscryptionGeneralredPng),
		ImgCardPrinceBlue:     canvas.NewImageFromResource(assets.InscryptionPrincebluePng),
		ImgCardPrinceRed:      canvas.NewImageFromResource(assets.InscryptionPrinceredPng),
		ImgCardSuit:           canvas.NewImageFromResource(assets.InscryptionSuitPng),
		ImgBackground:         canvas.NewImageFromResource(assets.InscryptionBackgroundPng),
	}
}
