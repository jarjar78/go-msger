package ui

import (
	"fmt"
	mn "mynet"
	"net"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type UI struct {
	UdpAddr   *net.UDPAddr
	App       fyne.App
	answers   *widget.Entry
	window    fyne.Window
	userlist  fyne.Window
	loginForm *widget.Box
	mainBox   *widget.Box
	User      string
}

func (ui *UI) CreateApp() {
	ui.App = app.New()
	ui.window = ui.App.NewWindow("Messenger")
	ui.CreateLoginForm()
	size := fyne.NewSize(300, 400)
	ui.window.Resize(size)
}

func (ui *UI) CreateLoginForm() {
	login := widget.NewEntry()
	ui.loginForm = widget.NewVBox(
		widget.NewLabel("Your name"),
		login,
		widget.NewButton("Login", func() {
			ui.User = login.Text
			mn.Login(ui.User, ui.UdpAddr)
			ui.CreateUserlistWindow()
			ui.window.SetContent(ui.mainBox)
		}),
	)
}

func (ui *UI) ShowWindow(c chan string, r chan string) {
	entry := widget.NewEntry()
	ui.answers = widget.NewMultiLineEntry()
	ui.answers.SetReadOnly(true)
	ui.mainBox = widget.NewVBox(
		entry,
		widget.NewButton("Send", func() {
			c <- entry.Text
			newText := ui.NewText(entry.Text, "me")
			ui.answers.SetText(newText + ui.answers.Text)
			entry.SetText("")
		}),
		ui.answers,
	)
	ui.mainBox.Hide()

	ui.window.SetContent(ui.loginForm)
	go ui.Reciever(r)

	ui.window.Show()
}

func (ui *UI) NewText(text string, from string) string {
	if from == "me" {
		text = ui.User + ": " + text
	} else if from == "you" {
		text = "You: " + text
	}

	return text + "\n"
}

func (ui *UI) Reciever(r chan string) {
	for {
		rec := <-r
		rec = ui.NewText(rec, "you")
		ui.answers.SetText(rec + ui.answers.Text)
		//fmt.Println("Received from net: " + rec)
		//fmt.Println("History: " + answers.Text)
	}
}

func (ui *UI) CreateUserlistWindow() {
	userlist := mn.GetUserlist(ui.User)

	ui.userlist = ui.App.NewWindow("Userlist")
	mainView := widget.NewVBox()
	for _, user := range userlist {
		mainView.Append(widget.NewButton(user, func() {}))
		fmt.Println(user)
	}
	ui.userlist.SetContent(mainView)
	ui.userlist.Show()
}
