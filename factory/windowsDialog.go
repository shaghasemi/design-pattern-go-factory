package factory

import "design-pattern-go-factory/button"

type WindowsDialog struct{}

func (d *WindowsDialog) createButton() button.Button {
	return &button.WindowsButton{}
}

func (d *WindowsDialog) RenderWindow() {
	ok := d.createButton()
	ok.Render()
}
