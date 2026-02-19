package eventsystem

type Event[T any] struct {
	EventType string
	Data      T
}

// EventHandler function usage example:
// func OnPlayerMove(event Event[PlayerMoveData]) {
//     fmt.Printf("Player moved to %v", event.Data)
// }

type EventHandler[T any] func(event Event[T])

func NewEvent[T any](data T, eventType string) *Event[T] {
	return &Event[T]{
		EventType: eventType,
		Data:      data,
	}
}
