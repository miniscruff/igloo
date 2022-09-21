package igloo

type FSMTransition[T ~string] struct {
	From T
	To []T
}

func NewFSMTransition[T ~string](from T, to ...T) FSMTransition[T] {
	return FSMTransition[T]{
		From: from,
		To: to,
	}
}

// FSM is a mini finite state machine
type FSM[T ~string] struct {
	current T
	transitions map[T]map[T]struct{}
	handlers map[T]*EventStoreZero
}

// NewWatchable will create a watchable with a starting value
func NewFSM[T ~string](startingValue T, transitions ...FSMTransition[T]) *FSM[T] {
	fsm := &FSM[T]{
		current: startingValue,
		transitions: make(map[T]map[T]struct{}),
		handlers: make(map[T]*EventStoreZero),
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

// Watch for changes to our value
func (fsm *FSM[T]) Transition(value T) {
	if fsm.current == value {
		return
	}

	possibleTransitions := fsm.transitions[fsm.current]
	_, canTransition := possibleTransitions[value]
	if !canTransition {
		return
	}

	fsm.current = value
	if fsm.handlers[value] != nil {
		fsm.handlers[value].Publish()
	}
}

func (fsm *FSM[T]) OnTransition(state T, handler EventHandlerZero) {
	if fsm.handlers[state] == nil {
		fsm.handlers[state] = &EventStoreZero{}
	}

	fsm.handlers[state].Subscribe(handler)
}
