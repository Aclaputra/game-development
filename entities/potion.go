package entities

// embedded struct to Sprite
type Potion struct {
	*Sprite
	AmtHeal uint
}
