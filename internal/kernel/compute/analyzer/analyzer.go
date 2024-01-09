package analyzer

import "context"

type Analyser struct {
}

func NewAnalyser() *Analyser {
	return &Analyser{}
}

func (a *Analyser) AnalyzeQuery(ctx context.Context, tokens []string) error {
	return nil
}
