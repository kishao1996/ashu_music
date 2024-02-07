package ui

import (
	"ashu_music/src"
	"ashu_music/src/play"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Init() {
	a := app.New()
	w := a.NewWindow("Ashu Music")

	list := widget.NewList(
		func() int {
			return len(src.Musics)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(src.Musics[i].Name)
		})
	list.OnSelected = func(id widget.ListItemID) {
		go play.Play(src.Musics[id].Path)
		fmt.Println(id, src.Musics[id].Path)
	}
	w.SetContent(list)

	w.ShowAndRun()
}
