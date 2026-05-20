package main

import "design-pattern-go-factory/factory"

func main() {
	dialog := factory.CreateDialog()
	dialog.RenderWindow()

}
