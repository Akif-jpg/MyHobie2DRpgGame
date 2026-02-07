package systems

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test payload tipi.
// Generic type safety'nin gerçekten çalıştığını doğrulamak için kullanılır.
type TestPayload struct {
	Message string
	Value   int
}

func TestEventManager_SubscribeAndPublish(t *testing.T) {
	// Arrange
	em := NewEventManager[TestPayload]()

	called := false

	id := em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		called = true
		assert.Equal(t, "hello", e.Payload.Message)
		assert.Equal(t, 42, e.Payload.Value)
	})

	assert.NotZero(t, id)

	// Act
	em.Publish(Event[TestPayload]{
		Type: "TEST_EVENT",
		Payload: TestPayload{
			Message: "hello",
			Value:   42,
		},
	})

	// Assert
	assert.True(t, called, "handler should be executed")
}

func TestEventManager_UnsubscribeByID(t *testing.T) {
	// Arrange
	em := NewEventManager[TestPayload]()

	called := false

	id := em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		called = true
	})

	// Unsubscribe before publish
	em.Unsubscribe("TEST_EVENT", id)

	// Act
	em.Publish(Event[TestPayload]{
		Type: "TEST_EVENT",
		Payload: TestPayload{
			Message: "should not be received",
		},
	})

	// Assert
	assert.False(t, called, "handler should NOT be executed after unsubscribe")
}

func TestEventManager_MultipleSubscribers(t *testing.T) {
	// Arrange
	em := NewEventManager[TestPayload]()

	count := 0

	em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		count++
	})

	em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		count++
	})

	// Act
	em.Publish(Event[TestPayload]{
		Type: "TEST_EVENT",
		Payload: TestPayload{
			Message: "fan-out",
		},
	})

	// Assert
	assert.Equal(t, 2, count, "all subscribers should be executed")
}

func TestEventManager_UnsubscribeOneDoesNotAffectOthers(t *testing.T) {
	// Arrange
	em := NewEventManager[TestPayload]()

	called1 := false
	called2 := false

	id1 := em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		called1 = true
	})

	em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		called2 = true
	})

	// Act
	em.Unsubscribe("TEST_EVENT", id1)

	em.Publish(Event[TestPayload]{
		Type: "TEST_EVENT",
		Payload: TestPayload{
			Message: "partial unsubscribe",
		},
	})

	// Assert
	assert.False(t, called1, "unsubscribed handler should not be called")
	assert.True(t, called2, "remaining handler should still be called")
}

func TestEventManager_HandlerExecutionIsSynchronous(t *testing.T) {
	// Arrange
	em := NewEventManager[TestPayload]()

	start := time.Now()

	em.Subscribe("TEST_EVENT", func(e Event[TestPayload]) {
		time.Sleep(50 * time.Millisecond)
	})

	// Act
	em.Publish(Event[TestPayload]{
		Type: "TEST_EVENT",
		Payload: TestPayload{
			Message: "sync test",
		},
	})

	// Assert
	elapsed := time.Since(start)
	assert.GreaterOrEqual(
		t,
		elapsed,
		50*time.Millisecond,
		"publish should block until handlers finish",
	)
}
