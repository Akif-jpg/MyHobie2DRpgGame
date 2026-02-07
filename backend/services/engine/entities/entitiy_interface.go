package entities

import "context"

type Entity interface {
	GetID() string
	GetName() string
	GetPosition() (x, y float64)
	Move(dx, dy float64)

	PreUpdate(ctx context.Context, deltaTime float64)
	Update(ctx context.Context, deltaTime float64)
	LastUpdate(ctx context.Context, deltaTime float64)

	Interact(other Entity)
	TakeDamage(amount int)
	IsAlive() bool
}
