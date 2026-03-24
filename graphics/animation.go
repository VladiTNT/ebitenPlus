package graphics

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// An Animation is just the basic data for one animation loop.
type Animation struct {
	Frames       []*ebiten.Image
	CurrentFrame int
	Delay        time.Duration
}

// Returns a new animation with the frame loop and delay specified.
func NewAnimation(frames []*ebiten.Image, delay time.Duration) *Animation {
	return &Animation{
		Frames:       frames,
		CurrentFrame: 0,
		Delay:        delay,
	}
}

// An animation system keeps track of several animations.
type AnimationSystem struct {
	animations       map[string]*Animation
	CurrentAnimation string
	lastFrame        time.Time
}

// Makes an animation system.
func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{
		animations: make(map[string]*Animation),
		lastFrame:  time.Now(),
	}
}

// Register's an animation on that key, it also sets the CurrentAnimation field to m
// to prevent you from using invalid animations.
func (a *AnimationSystem) Register(key string, m *Animation) {
	a.animations[key] = m
	a.CurrentAnimation = key
}

func (a *AnimationSystem) Sprite() *ebiten.Image {
	animation, ok := a.animations[a.CurrentAnimation]
	// Check if the animation is valid.
	if !ok {
		panic(fmt.Errorf("Error: Animation '%s' is invalid", a.CurrentAnimation))
	}

	// If the next frame has to play.
	if time.Since(a.lastFrame) > animation.Delay {
		a.lastFrame = time.Now()
		// We check if the next frame is the last in the loop.
		if animation.CurrentFrame+1 < len(animation.Frames) {
			// If it isn't, we return the next frame.
			animation.CurrentFrame++
			return animation.Frames[animation.CurrentFrame]
		} else {
			// If it is, we restart the loop.
			animation.CurrentFrame = 0
			return animation.Frames[0]
		}
	} else {
		return animation.Frames[animation.CurrentFrame]
	}
}

// Play changes which animation is currently playing.
func (a *AnimationSystem) Play(key string) {
	if a.CurrentAnimation != key {
		a.animations[a.CurrentAnimation].CurrentFrame = 0
	}
	a.CurrentAnimation = key
}
