package factory

import "design-pattern-go-factory/button"

type LinuxDialog struct{}

func (d *LinuxDialog) createButton() button.Button {
	return &button.LinuxButton{}
}

func (d *LinuxDialog) RenderWindow() {
	ok := d.createButton()
	ok.Render()
}
