package entities

type Entity interface {
	// GetID returns the unique identifier of the entity
	GetID() string

	// GetName returns the name of the entity
	GetName() string

	// GetPosition returns the current position of the entity in the game world
	GetPosition() (x, y float64)

	// Move changes the position of the entity by a specified amount
	Move(dx, dy float64)

	PreUpdate(deltaTime float64) // PreUpdate allows the entity to perform any necessary actions before the main update, such as resetting temporary states or preparing for interactions

	Update(deltaTime float64) // Update allows the entity to update its state, such as processing AI or animations

	LastUpdate(deltaTime float64) // LastUpdate allows the entity to perform any necessary actions after the main update, such as cleaning up temporary states or finalizing interactions
	// Interact allows the entity to interact with another entity
	Interact(other Entity)

	// TakeDamage reduces the health of the entity by a specified amount
	TakeDamage(amount int)

	// IsAlive checks if the entity is still alive (health > 0)
	IsAlive() bool
}
