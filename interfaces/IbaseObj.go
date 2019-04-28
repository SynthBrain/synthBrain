package interfaces

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/math32"
)

type IBaseObj interface {
	Node() *core.Node
	GetLocation() math32.Vector3
	SetLocation(math32.Vector3)
}
