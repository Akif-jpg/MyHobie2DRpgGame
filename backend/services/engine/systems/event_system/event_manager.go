package systems

import "sync"

// EventManager struct is responsible for managing event subscriptions
// and executing subscriber functions when an event is published.
type EventManager[T any] struct {
	// subscribers map'i:
	// key   -> eventType (string)
	// value -> eventType'a ait handler kayıtları
	subscribers map[string]map[uint64]func(Event[T])

	// nextID, her yeni subscription için
	// benzersiz bir kimlik üretmekte kullanılır.
	nextID uint64

	// mu, concurrent access durumlarında
	// subscribers map'ini korumak için kullanılır.
	mu sync.RWMutex
}

// NewEventManager function is a constructor for EventManager.
func NewEventManager[T any]() *EventManager[T] {
	return &EventManager[T]{
		subscribers: make(map[string]map[uint64]func(Event[T])),
	}
}

// Subscribe function registers a handler function for a specific event type.
//
// Fonksiyon bir subscription ID döner.
// Bu ID, Unsubscribe işlemi için kullanılır.
func (em *EventManager[T]) Subscribe(
	eventType string,
	handler func(Event[T]),
) uint64 {
	em.mu.Lock()
	defer em.mu.Unlock()

	em.nextID++

	if _, ok := em.subscribers[eventType]; !ok {
		em.subscribers[eventType] = make(map[uint64]func(Event[T]))
	}

	em.subscribers[eventType][em.nextID] = handler

	return em.nextID
}

// Publish function executes all subscribed handler functions
// for the given event's type.
func (em *EventManager[T]) Publish(event Event[T]) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	if handlers, ok := em.subscribers[event.Type]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

// Unsubscribe function removes a previously registered handler
// using its subscription ID.
func (em *EventManager[T]) Unsubscribe(
	eventType string,
	subscriptionID uint64,
) {
	em.mu.Lock()
	defer em.mu.Unlock()

	if handlers, ok := em.subscribers[eventType]; ok {
		delete(handlers, subscriptionID)

		// Handler kalmadıysa map temizlenir
		if len(handlers) == 0 {
			delete(em.subscribers, eventType)
		}
	}
}
