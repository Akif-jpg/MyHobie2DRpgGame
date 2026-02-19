package components

import (
	"github.com/Akif-jpg/MyHobieMMORPGGame/services/engine/physics/collider"
	"github.com/Akif-jpg/MyHobieMMORPGGame/services/engine/physics/geometry"
	"github.com/google/uuid"
)

type physicType uint16

const (
	staticBody    physicType = 0
	kinematicBody physicType = 1
	rigidBody     physicType = 2
)

type PhysicComponent struct {
	collider         *collider.Collider
	PhysicType       *physicType
	Velocity         *geometry.Vector2
	RotationVelocity float64
	IID              string
}

func NewPhysicComponent(pt physicType, c *collider.Collider) *PhysicComponent {
	return &PhysicComponent{
		collider:         c,
		PhysicType:       &pt,
		Velocity:         geometry.NewVector2(0, 0),
		RotationVelocity: 0,
		IID:              uuid.New().String(),
	}
}

func (c *PhysicComponent) OnCreate() {
}

func (c *PhysicComponent) OnDestroy() {
}

func (c *PhysicComponent) Update(deltaTime float64) {
}

func (c *PhysicComponent) ComponentID() string {
	return c.IID
}
