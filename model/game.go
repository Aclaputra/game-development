package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	ArcadeFontText     *text.GoTextFaceSource
	SkeletonSprite     *ebiten.Image
	GameTileSprite     = make([]*ebiten.Image, 0)
	SkeletonFrameIndex int
	SkeletonFramePixel int
	SkeletonStepFrames = []int{
		16,  // ok
		80,  // ok
		144, // ok
		208, // ok
		272, // ok
		336, // ok
		400, // ok
		464, // ok
		528, // ok
	}
	SkeletonDirectionFrames = map[string]int{
		"north": 16,
		"west":  80,
		"south": 144,
		"east":  208,
	}
	BasePosition      = 300
	CountMovementX    = BasePosition
	CountMovementY    = BasePosition
	TimeCounter       int
	SkeletonDirection string
)
