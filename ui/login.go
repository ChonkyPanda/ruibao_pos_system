package ui

import (
	"ruibao_pos_system/db"
	"ruibao_pos_system/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowLogin renders the login page into w.
// onSuccess is called with the authenticated user once login succeeds.
func ShowLogin(w fyne.Window, onSuccess func(user *db.User)) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	loginButton := widget.NewButton("Login", func() {
		user, err := db.GetUserByName(usernameEntry.Text)

		if err != nil || user == nil || utils.CheckPasswordHash(user.Password, passwordEntry.Text)  {
			errorLabel.SetText("Invalid username or password")
			errorLabel.Show()
			return
		}
		errorLabel.Hide()
		onSuccess(user)
	})

	goToSignUp := widget.NewButton("Don't have an account? Sign up", func() {
		ShowSignUp(w, onSuccess)
	})

	form := container.NewVBox(
		widget.NewLabelWithStyle("RuiBao POS - Login", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		usernameEntry,
		passwordEntry,
		errorLabel,
		loginButton,
		goToSignUp,
	)

	w.SetContent(container.NewCenter(form))
}