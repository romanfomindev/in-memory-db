package kernel

import "context"

type Compute interface {
	HandleQuery(context.Context, string) error
}

type Parser interface {
	ParseQuery(context.Context, string) ([]string, error)
}

type Analyzer interface {
	AnalyzeQuery(context.Context, []string) error
}
