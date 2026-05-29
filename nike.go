package main

type Nike struct{}

// makeShoe implements [ISportsFactory].
func (n *Nike) makeShoe() IShoe {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

// makeShirt implements [ISportsFactory].
func (n *Nike) makeShirt() IShirt {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}
