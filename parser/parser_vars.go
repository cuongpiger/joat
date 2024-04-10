package parser

import (
	"sync"
)

var (
	parserIns  IParser
	parserOnce sync.Once
)
