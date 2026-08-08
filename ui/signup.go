package ui

import (
	"errors"
	"ruibao_pos_system/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

func ShowSignUp(w fyne.Window, onSuccess func(user *db.User)) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	errorLabel := widget.NewLabel("")
	errorLabel.Hide()

	signUpButton := widget.NewButton("Create Account", func() {
		user, err := db.GetUserByName(usernameEntry.Text)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// A real DB error, not just "no user found"
			errorLabel.SetText("Unable to check username, try again")
			errorLabel.Show()
			return
		}

		if user != nil {
			errorLabel.SetText("Username already exists")
			errorLabel.Show()
			return
		}

		newUser, err := db.CreateUser(usernameEntry.Text, passwordEntry.Text, false)
		if newUser != nil && err != nil {
			errorLabel.SetText("Unable to create user")
			errorLabel.Show()
			return
		}

		errorLabel.Hide()
		ShowLogin(w, onSuccess)
	})

	backToLogin := widget.NewButton("Already have an account? Log in", func() {
		ShowLogin(w, onSuccess)
	})

	form := container.NewVBox(
		widget.NewLabelWithStyle("RuiBao POS - Sign Up", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		usernameEntry,
		passwordEntry,
		errorLabel,
		signUpButton,
		backToLogin,
	)

	w.SetContent(container.NewCenter(form))
}