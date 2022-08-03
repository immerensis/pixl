package ui

import (
	"fyne.io/fyne/v2"
	"github.com/immerensis/pixl/apptype"
	"github.com/immerensis/pixl/pxcanvas"
	"github.com/immerensis/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
