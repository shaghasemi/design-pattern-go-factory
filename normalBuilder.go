package main

type NormalBuilder struct {
	windowsType string
	doorType    string
	floor       int
}

// setHouse implements [IBuilder].
func (b *NormalBuilder) setHouse() House {
	panic("unimplemented")
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowsType = "Wooden Windows"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowsType,
		floor:      b.floor,
	}
}
