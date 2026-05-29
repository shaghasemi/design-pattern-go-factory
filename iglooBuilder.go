package main

type IglooBuilder struct {
	windowsType string
	doorType    string
	floor       int
}

// setHouse implements [IBuilder].
func (b *IglooBuilder) setHouse() House {
	panic("unimplemented")
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowsType = "Wooden Windows"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 2
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowsType,
		floor:      b.floor,
	}
}
