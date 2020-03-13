package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rivo/tview"
)

func exitOnError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func makeGorilla() tview.Primitive {
	cmd := exec.Command("pixterm", "gorilla.png")

	b, err := cmd.Output()
	if err != nil {
		exitOnError(err)
	}

	view := tview.NewTextView()
	view.SetBorder(false)
	view.SetTextAlign(tview.AlignCenter)
	view.SetDynamicColors(true)

	view.SetText(tview.TranslateANSI(string(b)))
	return view
}

func makeGorillaStatus() tview.Primitive {
	view := tview.NewTextView()
	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTextAlign(tview.AlignCenter)
	view.SetText(`六道ゴリラ
----------
HP	999
MP	999
ATK	999
DEF	999`)
	return view
}

func makeMenu() tview.Primitive {
	form := tview.NewForm().
		AddDropDown("攻撃", []string{
			"バナナを投げる",
			"バナナを奪う",
			"バナナを食べる",
			"バナナを捨てる",
		}, 0, nil)

	box := tview.NewBox()
	box.SetBorder(true)

	grid := tview.NewGrid()
	grid.AddItem(form, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(box, 0, 0, 1, 1, 0, 0, true)

	return grid
}

func main() {
	status := makeGorillaStatus()
	gorilla := makeGorilla()
	menu := makeMenu()

	grid := tview.NewGrid().SetRows(8, 0, 5)
	grid.AddItem(status, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(gorilla, 1, 0, 1, 1, 0, 0, true)
	grid.AddItem(menu, 2, 0, 1, 1, 0, 0, true)

	pages := tview.NewPages().AddAndSwitchToPage("main", grid, true)

	app := tview.NewApplication().SetRoot(pages, true)
	app.SetFocus(menu)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
