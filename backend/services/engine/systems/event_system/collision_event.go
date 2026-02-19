package eventsystem

import "github.com/Akif-jpg/MyHobieMMORPGGame/services/engine/physics/collider"

type CollisionData struct {
	collider1 collider.Collider
	collider2 collider.Collider
}

type CollisionEvent struct {
	*EventManager[CollisionData]
}

func (c *CollisionEvent) OnCollide(collisionData CollisionData) {
}
