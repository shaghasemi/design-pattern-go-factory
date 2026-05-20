package button

type WindowsButton struct{}

func (b *WindowsButton) Render() {
	println("Rendering Windows Button")
	// b.Render()
}

func (b *WindowsButton) OnClick() {
	println("Windows Button clicked")
	// b.OnClick()
}
