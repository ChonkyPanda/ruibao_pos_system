package main

import "ruibao_pos_system/db"
import "ruibao_pos_system/ui"

import "fyne.io/fyne/v2"
import "fyne.io/fyne/v2/app"
// import "fyne.io/fyne/v2/widget"


func main() {
	db.InitDB()
	db.CreateUser("Admin", "Jasonxie175", true)

	a := app.NewWithID("com.ruibao.pos")
	w := a.NewWindow("RuiBao POS System")
	w.Resize(fyne.NewSize(1000, 700))

	ui.ShowInventory(w, nil)

	// ui.ShowLogin(w, func(user *db.User) {
	// 	// Swap this out for a real dashboard page later, e.g. ui.ShowDashboard(w, user)
	// 	w.SetContent(widget.NewLabel("Welcome, " + user.Name + "!"))
	// })

	w.ShowAndRun()
}