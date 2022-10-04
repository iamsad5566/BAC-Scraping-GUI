package gui

import (
	"bac-scraper-gui/config"
	"bac-scraper-gui/model"
	"image/color"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var width float32 = 600
var height float32 = 800
var year string = ""
var month string = ""
var mp map[string]*widget.Entry = make(map[string]*widget.Entry)

func OpenWindow() {
	guiApp := app.New()
	guiApp.Settings().SetTheme(&myTheme{})
	window := guiApp.NewWindow("Web scraper for BAC")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(true)
	window.SetContent(sortLayOut(window))
	window.ShowAndRun()
}

func sortLayOut(window fyne.Window) fyne.CanvasObject {
	title := title()
	firstRow := yearAndMonthEntry()
	secondRow := membersEntry()

	space := canvas.NewText(" ", color.White)
	space.TextSize = 40.0
	content := container.NewBorder(space, nil, nil, nil, container.NewVBox(title, firstRow, secondRow, handleStart(window)))
	return content
}

func title() fyne.CanvasObject {
	title := canvas.NewText("BAC lab 報帳自動結算軟體", color.Black)
	title.TextSize = 40.0
	verticle := canvas.NewText(" ", color.White)
	horizontal := canvas.NewText(" ", color.White)
	horizontal.TextSize = 70.0
	content := container.NewBorder(verticle, verticle, horizontal, horizontal, container.NewCenter(title))
	return content
}

func yearAndMonthEntry() fyne.CanvasObject {
	labelYear := widget.NewLabel("報帳年度")
	labelMonth := widget.NewLabel("報帳月份")
	entryY := widget.NewEntry()
	entryM := widget.NewEntry()
	entryY.SetPlaceHolder("請輸入民國")
	entryM.SetPlaceHolder("請輸入月份")
	mp["year"] = entryY
	mp["month"] = entryM
	content := container.NewVBox(labelYear, entryY, labelMonth, entryM)

	return content
}

func membersEntry() fyne.CanvasObject {
	label := widget.NewLabel("Members")
	input := widget.NewEntry()
	nameList := func() string {
		yml := config.LoadYml()
		t := ""
		for _, name := range yml.Members {
			t += (name + "\n")
		}
		return t
	}()
	input.SetText(nameList)
	mp["members"] = input
	size := fyne.Size{Width: width, Height: 200}

	content := container.NewVBox(label, container.NewGridWrap(size, input))
	return content
}

func handleStart(window fyne.Window) fyne.CanvasObject {
	button := widget.NewButton("Go!", func() {
		envar := config.LoadYml()
		model.OverWriteYml(mp, envar)
		cmd := exec.Command("java", "-jar", "Scraping.jar", mp["year"].Text, mp["month"].Text)
		cmd.Run()
		// window.Close()
	})

	space := canvas.NewText(" ", color.White)

	content := container.NewBorder(space, space, space, space, button)
	return content
}
