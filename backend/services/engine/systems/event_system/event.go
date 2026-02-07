package systems

// Event struct represents an event that can be emitted and listened to within the game engine.
type Event[T any] struct {
	// Type, event routing (hangi abonelere gideceği) için kullanılır.
	Type string

	// Payload, event ile taşınan gerçek veridir.
	// T tipi compile-time'da garanti edilir.
	Payload T
}
