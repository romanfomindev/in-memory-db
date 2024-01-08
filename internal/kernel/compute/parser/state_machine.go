package parser

import (
	"errors"
	"strings"
)

const (
	FIND_LETTER_EVENT = iota
	FIND_WHITESPACE_EVENT
	EVENTS_COUNT
)

const (
	INITIAL_STATE = iota
	WORD_STATE
	WHITESPACE_STATE
	INVALID_STATE
	STATES_COUNT
)

type transition struct {
	jump   func(byte) int
	action func()
}

type stateMachine struct {
	transitions [STATES_COUNT][EVENTS_COUNT]transition
	state       int
	tokens      []string
	sb          strings.Builder
}

func newStateMachine() *stateMachine {
	sm := &stateMachine{
		state: INITIAL_STATE,
	}

	sm.transitions = [STATES_COUNT][EVENTS_COUNT]transition{
		INITIAL_STATE: {
			FIND_LETTER_EVENT:     transition{jump: sm.addSymbol},
			FIND_WHITESPACE_EVENT: transition{jump: sm.skipWhiteSpace},
		},
		WORD_STATE: {
			FIND_LETTER_EVENT:     transition{jump: sm.addSymbol},
			FIND_WHITESPACE_EVENT: transition{jump: sm.skipWhiteSpace, action: sm.appendToken},
		},
		WHITESPACE_STATE: {
			FIND_LETTER_EVENT:     transition{jump: sm.addSymbol},
			FIND_WHITESPACE_EVENT: transition{jump: sm.skipWhiteSpace},
		},
		INVALID_STATE: {},
	}

	return sm
}

func (sm *stateMachine) parse(query string) ([]string, error) {

	for i := 0; i < len(query); i++ {
		if isLetter(query[i]) {
			sm.processEvent(FIND_LETTER_EVENT, query[i])
		} else if isWhiteSpace(query[i]) {
			sm.processEvent(FIND_WHITESPACE_EVENT, query[i])
		} else {
			return nil, errors.New("INVALID SYMBOL IN QUERY")
		}
	}

	return sm.tokens, nil
}

func (sm *stateMachine) processEvent(event int, symbol byte) {
	transition := sm.transitions[sm.state][event]
	sm.state = transition.jump(symbol)
	if transition.action != nil {
		transition.action()
	}
}

func (sm *stateMachine) addSymbol(b byte) int {
	sm.sb.WriteByte(b)

	return WORD_STATE
}

func (sm *stateMachine) skipWhiteSpace(b byte) int {
	return WHITESPACE_STATE
}

func (sm *stateMachine) appendToken() {
	sm.tokens = append(sm.tokens, sm.sb.String())
	sm.sb.Reset()
}
