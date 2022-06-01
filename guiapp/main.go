package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type randomFact struct {
	Text      string `json:"text"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
	Language  string `json:"language"`
	Permalink string `json:"permalink"`
}

func getRandomFact() (randomFact, error) {
	var fact randomFact
	resp, err := client.Get("https://uselessfacts.jsph.pl//random.json?language=en")
	if err != nil {
		return randomFact{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return randomFact{}, err
	}
	return fact, nil
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("Get useless fact")
	win.Resize(fyne.NewSize(800, 500))

	title := canvas.NewText("Get Your useless Facts", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Fact", func() {
		fact, err := getRandomFact()
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			factText.SetText(fmt.Sprintf("Text : %s\nSource : %s\nSource Url : %s\nLanguage : %s\nPermalink : %s", fact.Text, fact.Source, fact.SourceUrl, fact.Language, fact.Permalink))
		}
	})
	button.Resize(fyne.NewSize(200, 80))

	hbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vbox := container.New(layout.NewVBoxLayout(), title, hbox, widget.NewSeparator(), factText)

	win.SetContent(vbox)

	// hello := widget.NewLabel("Hello Fyne!")
	// win.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))

	win.ShowAndRun()
}
