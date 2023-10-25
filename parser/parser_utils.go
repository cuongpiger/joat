package parser

import "github.com/cuongpiger/joat/parser/entity"

func NewParser() (IParser, error) {
	parserOnce.Do(func() {
		parserIns = new(entity.Parser)
	})

	// currently, there is no error to return
	return parserIns, nil
}
