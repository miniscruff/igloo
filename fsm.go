package igloo

type FSMTransition[T ~string] struct {
	From T
	To   []T
}

func NewFSMTransition[T ~string](from T, to ...T) FSMTransition[T] {
	return FSMTransition[T]{
		From: from,
		To:   to,
	}
}

// FSM is a mini finite state machine
type FSM[T ~string] struct {
	current      T
	last         T
	transitions  map[T]map[T]struct{}
	fromHandlers map[T]*EventStoreZero
	toHandlers   map[T]*EventStoreZero
}

// NewWatchable will create a watchable with a starting value
func NewFSM[T ~string](startingValue T, transitions ...FSMTransition[T]) *FSM[T] {
	fsm := &FSM[T]{
		current:      startingValue,
		transitions:  make(map[T]map[T]struct{}),
		fromHandlers: make(map[T]*EventStoreZero),
		toHandlers:   make(map[T]*EventStoreZero),
	}
	for _, t := range transitions {
		fsm.transitions[t.From] = map[T]struct{}{}
		for _, to := range t.To {
			fsm.transitions[t.From][to] = struct{}{}
		}
	}

	return fsm
}

func (fsm *FSM[T]) Current() T {
	return fsm.current
}

func (fsm *FSM[T]) Last() T {
	return fsm.last
}

func (fsm *FSM[T]) CanTransition(value T) bool {
	if fsm.current == value {
		return false
	}

	possibleTransitions := fsm.transitions[fsm.current]
	_, canTransition := possibleTransitions[value]

	return canTransition
}

// Transition to a new state
// will return true if we were able to transition, otherwise false
func (fsm *FSM[T]) Transition(value T) bool {
	if !fsm.CanTransition(value) {
		return false
	}

	fsm.last = fsm.current
	fsm.current = value

	if fsm.fromHandlers[fsm.last] != nil {
		fsm.fromHandlers[fsm.last].Publish()
	}

	if fsm.toHandlers[value] != nil {
		fsm.toHandlers[value].Publish()
	}

	return true
}

func (fsm *FSM[T]) OnTransitionTo(state T, handler EventHandlerZero) {
	if fsm.toHandlers[state] == nil {
		fsm.toHandlers[state] = &EventStoreZero{}
	}

	fsm.toHandlers[state].Subscribe(handler)
}

func (fsm *FSM[T]) OnTransitionFrom(state T, handler EventHandlerZero) {
	if fsm.fromHandlers[state] == nil {
		fsm.fromHandlers[state] = &EventStoreZero{}
	}

	fsm.fromHandlers[state].Subscribe(handler)
}
