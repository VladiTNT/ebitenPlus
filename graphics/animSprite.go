package graphics

import (
	"github.com/VladiTNT/ebitenPlus"
	"github.com/hajimehoshi/ebiten/v2"
)

// An animated sprite is just an animation system with a position vector.
// You will have to register each animation manually.
type AnimatedSprite struct {
	Anim *AnimationSystem
	Pos  ebitenPlus.Point
}

// Makes a basic animated sprite.
func NewAnimatedSprite(pos ebitenPlus.Point) *AnimatedSprite {
	return &AnimatedSprite{
		Anim: NewAnimationSystem(),
		Pos:  pos,
	}
}

// Draws the sprite to the screen.
func (s *AnimatedSprite) Draw(screen *ebiten.Image, p ebitenPlus.Point) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(s.Anim.Sprite(), op)
}
