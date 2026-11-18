package parser

import "Nginx/internal"

type LineParser interface {
	ParseLine(line string, lineNumber int) (internal.Event, error)
}
