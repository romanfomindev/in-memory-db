package parser

import (
	"context"

	"github.com/romanfomindev/in-memory-db/internal/kernel"
)

type Parser struct {
	sm *stateMachine
}

func NewParser() kernel.Parser {
	sm := newStateMachine()
	return &Parser{
		sm: sm,
	}
}

func (p *Parser) ParseQuery(ctx context.Context, queryStr string) ([]string, error) {
	tokens, err := p.sm.parse(queryStr)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func isWhiteSpace(symbol byte) bool {
	return symbol == '\t' || symbol == '\n' || symbol == ' '
}

func isLetter(symbol byte) bool {
	return (symbol >= 'a' && symbol <= 'z') ||
		(symbol >= 'A' && symbol <= 'Z') ||
		(symbol >= '0' && symbol <= '9') ||
		(symbol == '_')
}
