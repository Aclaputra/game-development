package npc

import (
	"fmt"

	"github.com/Aclaputra/game-development/constant"
	"github.com/Aclaputra/game-development/helper"
	"github.com/Aclaputra/game-development/model"
)

type (
	Skeleton interface {
		Render() error
	}
	skeleton struct {
		skeletonDirection string
	}
)

func NewSkeleton(direction string) Skeleton {
	return &skeleton{
		skeletonDirection: direction,
	}
}

func (sk *skeleton) Render() error {
	model.SkeletonFramePixel = model.SkeletonStepFrames[model.SkeletonFrameIndex]
	reqLoadAndCropImage := &model.LoadAndCropImageRequest{
		Path:   constant.SKELETON_SPRITE_PATH,
		X:      model.SkeletonFramePixel,
		Y:      model.SkeletonDirectionFrames[sk.skeletonDirection],
		Width:  30,
		Height: 60,
	}
	skeletonImg, err := helper.LoadAndCropImage(reqLoadAndCropImage)
	if err != nil {
		return fmt.Errorf("cannot get %v", constant.SKELETON_SPRITE_PATH)
	}
	model.SkeletonSprite = skeletonImg

	if reachedSomeTick := model.TimeCounter >= 5; reachedSomeTick {
		model.SkeletonFrameIndex++
		model.TimeCounter = constant.RESET_FROM_START
	}

	if skeletonReachedTheLastFrame := model.SkeletonFrameIndex >= len(model.SkeletonStepFrames); skeletonReachedTheLastFrame {
		model.SkeletonFrameIndex = constant.RESET_FROM_START + 1
	}

	if skeletonReachSomeDistance := model.CountMovement >= 900; skeletonReachSomeDistance {
		model.CountMovement = constant.RESET_FROM_START
	}

	return nil
}
