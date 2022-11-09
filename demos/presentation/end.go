package main

import (
	"fmt"

	"github.com/tinywolf3/tcell/v2"
	"github.com/tinywolf3/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/tinywolf3/tview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
