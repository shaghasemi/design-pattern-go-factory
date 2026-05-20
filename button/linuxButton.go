package button

type LinuxButton struct{}

func (b *LinuxButton) Render() {
	println("Rendering Linux Button")
	// b.Render()
}

func (b *LinuxButton) OnClick() {
	println("Linux Button clicked")
	// b.OnCLick()
}
