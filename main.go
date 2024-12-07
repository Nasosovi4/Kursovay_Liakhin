package main

import (
	sklad "sklad/app"
	"sklad/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	
	app := app.New()
	app.SetIcon(sklad.Icon())
	window := app.NewWindow("Система учетов товара на складе") // новое окно
	app.Settings().SetTheme(theme.DarkTheme())
	window.Resize(fyne.NewSize(800, 600))
	window.SetMaster()
	window.CenterOnScreen() // перемещение на центр экрана

	
	appStruct := sklad.NewApp()
	appStruct.Window = window
	appStruct.EnterWindow()
	appStruct.DB = db
	window.ShowAndRun() // запуск окна
}
