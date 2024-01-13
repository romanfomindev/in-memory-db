package compute

import (
	"context"

	"github.com/romanfomindev/in-memory-db/internal/kernel"
)

type Compute struct {
	parser   kernel.Parser
	analyzer kernel.Analyzer
}

func NewCompute(parser kernel.Parser, analyzer kernel.Analyzer) kernel.Compute {
	return &Compute{
		parser:   parser,
		analyzer: analyzer,
	}
}

func (c *Compute) HandleQuery(ctx context.Context, queryString string) error {
	_, err := c.parser.ParseQuery(ctx, queryString)
	if err != nil {
		return err
	}

	return nil
}
