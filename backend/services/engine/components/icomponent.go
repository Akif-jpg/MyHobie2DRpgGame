package components

type Component interface {
	// Define the methods that all components must implement
	NewComponent() Component
	ComponentID() string
	Name() string
	IsActive() bool
	SetActive(active bool) bool
	Render() bool
	Serialize() []byte
	Deserialize(data []byte) error
	Clone() Component
	Reset() error
	Start() error
	Update(deltaTime float64)
	OnCreate()
	OnDestroy()
}
