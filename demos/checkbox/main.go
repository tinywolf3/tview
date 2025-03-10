// Demo code for the Checkbox primitive.
package main

import "github.com/tinywolf3/tview"

func main() {
	app := tview.NewApplication()
	checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	if err := app.SetRoot(checkbox, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
