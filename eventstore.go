package igloo

// EventHandlerZero handles events with no arguments
type EventHandlerZero func()

// EventStoreZero contains all the events we are subscribed to with no arguments
type EventStoreZero struct {
	subs []EventHandlerZero
}

// Subscribe to our event
func (e *EventStoreZero) Subscribe(fn EventHandlerZero) {
	e.subs = append(e.subs, fn)
}

// Publish an event to all subscribers
func (e *EventStoreZero) Publish() {
	for _, ev := range e.subs {
		ev()
	}
}

// EventHandlerOne handles events with one argument
type EventHandlerOne[T any] func(T)

// EventStoreOne contains all the events we are subscribed to with one argument
type EventStoreOne[T any] struct {
	subs []EventHandlerOne[T]
}

// Subscribe to our event
func (e *EventStoreOne[T]) Subscribe(fn EventHandlerOne[T]) {
	e.subs = append(e.subs, fn)
}

// Publish an event to all subscribers
func (e *EventStoreOne[T]) Publish(value T) {
	for _, ev := range e.subs {
		ev(value)
	}
}

// EventHandlerTwo handles events with two arguments
type EventHandlerTwo[T0, T1 any] func(T0, T1)

// EventStoreTwo contains all the events we are subscribed to with two arguments
type EventStoreTwo[T0, T1 any] struct {
	subs []EventHandlerTwo[T0, T1]
}

// Subscribe to our event
func (e *EventStoreTwo[T0, T1]) Subscribe(fn EventHandlerTwo[T0, T1]) {
	e.subs = append(e.subs, fn)
}

// Publish an event to all subscribers
func (e *EventStoreTwo[T0, T1]) Publish(t0 T0, t1 T1) {
	for _, ev := range e.subs {
		ev(t0, t1)
	}
}
